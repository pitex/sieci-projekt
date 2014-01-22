package joinservice

import (
	"net"
	"log"
	"./tree"
	"os"
	"../protocols/sip"
	"strconv"
	"../protocols/stp"
	"time"
)

type ServerFullError struct {
	Address	string
}

func (err ServerFullError) Error() string {
	return "Server " + err.Address + " is already full!"
}

type Server struct {
	//	Server IP.
	Address		string

	//	Network parent IP.
	Parent		string

	//	Children IPs.
	Children	[]string

	//	How many children already exist.
	ChildNumber	int

	// If the server is root, it is pointer to node representing it, otherwise it is nil.
	Root		*tree.Node
}

//	Creates a new server with giben ip and capacity:
//	-	if root is false, it gets connected to parent,
//	-	else initializes tree node representing network structure.
func NewServer(ip string, parent string, capacity int, root bool) *Server{
	if !root {
		return &Server{ip, parent, make([]string, capacity-1), 0, nil}
	}
	return  &Server{ip, parent, make([]string, capacity-1), 0, tree.NewNode(ip, capacity)}
}

//	Sends chart through given socket
func SendChart(ip string) {
	socket, err := net.Dial("tcp", ip+":666")

	log.Println(err)

	sipMsg := sip.Message{}
	sipMsg.Type = "TRA"
	err = sip.Request(socket, sipMsg)

	if err != nil {
		log.Fatal(err)
	}

	chart, _ := os.Open("./resources/chart.html")
	fileBytes := make([]byte, 2048)

	log.Println("Begginning transfer")

	for {
		n, _ := chart.Read(fileBytes)

		stpMsg := stp.Message{}
		stpMsg.Data = string(fileBytes[:n])

		err = stp.Request(socket, stpMsg)

		if err != nil {
			log.Fatal(err)
		}

		if n < 2048 {
			break
		}
	}

	log.Println("Transfer finished")

	socket.Write([]byte("END||"))
}

//	We receive the chart script from our parent, we have to handle the data and send it to children.
func (s *Server) HandleChartTransfer(socket net.Conn) {
	file, _ := os.Create("./resources/chart.html")

	log.Println("Begginning receiving chart")

	for {
		read := make([]byte, 4096)

		n, _ := socket.Read(read)

		if sip.ExtractType(string(read[:n])) == "END" {
			break
		}

		file.Write(read[:n])

		socket.Write([]byte("INF||"))
	}

	log.Println("Done receiving chart")

	file.Close()

	for i := 0; i<s.ChildNumber; i++ {
		SendChart(s.Children[i])
	}
}

//	ROOT ONLY - travelling tree and adding nodes' description into our script
func (s *Server) BuildChart() {
	f, _ := os.OpenFile("./resources/chart.html", os.O_RDWR|os.O_APPEND, 0660)
	tree.DFS(s.Root, "", f)
	f.Close()
}

//	Rewrites input file to output file in APPEND MODE
func RewriteFile(input string, output string) {
	log.Printf("Opening %s\n", input)
	infile, _ := os.Open(input)
	log.Printf("Opening %s\n", output)
	outfile, _ := os.OpenFile(output, os.O_RDWR|os.O_APPEND, 0660)

	defer infile.Close()
	defer outfile.Close()

	log.Println("Rewriting")
	for {
		b := make([]byte, 1024)
		read, _ := infile.Read(b)
		
		outfile.Write(b[:read])

		if read < 1024 {
			break
		}
	}
}

//	ROOT ONLY - We create chart script and send it to children so they can update their charts
func (s *Server) CreateChart() {
	log.Println("Opening chart.html")
	os.Create("./resources/chart.html")

	log.Println("Creating begginning of chart")
	RewriteFile("./resources/chart_beg", "./resources/chart.html")

	log.Println("Building chart")
	s.BuildChart()

	log.Println("Creating ending of chart")
	RewriteFile("./resources/chart_end", "./resources/chart.html")
}

//	ROOT ONLY - When we know that there is a new machine pending. 
//	We need to find it place in out net and send the information about it to our children.
//	We also need to create updated chart and send it to children, too.
func (s *Server) HandleNewMachine(socket net.Conn, msg string) {
	log.Println("Starting parent search")
	DataMap := sip.InterpreteData(sip.ExtractData(msg))
	fatherNode, _ := tree.FindSolution(s.Root, -1)

	log.Printf("Parent = %s\n",fatherNode.IP)

	capacity, _ := strconv.Atoi(DataMap["capacity"])
	tree.AddNewChild(fatherNode, tree.NewNode(DataMap["ip"], capacity))
	newMes := sip.FND(DataMap["ip"], fatherNode.IP)
	
	if fatherNode.IP == s.Address {
		log.Printf("Adding child %s %s\n",s.Address, DataMap["ip"])
		s.AddChild(DataMap["ip"])
		_, err := socket.Write([]byte(newMes.ToString()))
		log.Println(err)
	} else {
		s.AskChildren(newMes.ToString())
	}

	s.CreateChart()

	time.Sleep(1 * time.Second)

	for i := 0; i<s.ChildNumber; i++ {
		log.Printf("Sending chart to %s\n", s.Children[i])
		SendChart(s.Children[i])
	}
}

//	Determines how to react for a SIM message depending on its type.
func (s *Server) SIPMessageReaction(socket net.Conn, msg string) {
	log.Print("Message type: %s\n", sip.ExtractType(msg))
	switch sip.ExtractType(msg) {
		case "BLD", "REQ" :
			sip.SendInfo(socket, msg)
			if s.Root != nil {
				s.HandleNewMachine(socket, msg)
			} else {
				s.TellParent(msg)
			}
		case "TRA" :
			sip.SendInfo(socket, msg)
			s.HandleChartTransfer(socket)
		case "FND" :
			sip.SendInfo(socket, msg)
			DataMap := sip.InterpreteData(sip.ExtractData(msg))

			if s.Address == DataMap["parent"] { 
				s.AddChild(DataMap["child"])
				break
			}
			s.AskChildren(msg)
	}
}

//	Sending msg to all children.
func (s *Server) AskChildren(msg string) {
	message := sip.GetMessage(msg)

	for i := 0; i < s.ChildNumber; i++ {
		child := s.Children[i]
		socket, _ := net.Dial("tcp", child+":666")
		sip.Request(socket, message)
	}
}

//	Sending msg to parent.
func (s *Server) TellParent(msg string) {
	message := sip.GetMessage(msg)

	socket, _ := net.Dial("tcp", s.Parent+":666")
	sip.Request(socket, message)
}

//	When we receive information that we are to create a connection with new machine, it becomes our child.
func (s *Server) AddChild(address string) error {
	if len(s.Children) == s.ChildNumber {
		return ServerFullError{s.Address}
	}

	s.Children[s.ChildNumber] = address
	s.ChildNumber++

	return nil
}

//	Reads from socket and then runs function which checks received message.
func (s *Server) handleConnection(sock net.Conn) {
	msg := make([]byte, 4096)

	n, _ := sock.Read(msg)

	log.Printf("Incoming msg: %s\n", string(msg[:n]))

	s.SIPMessageReaction(sock, string(msg[:n]))
}

//	Starts server - it listens on port 666.
func (s *Server) Start() error {
	log.Printf("Starting server with parent: %s\n", s.Parent)
	ln, err := net.Listen("tcp", s.Address + ":666")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, _ := ln.Accept()

		log.Println("Got something!")

		go s.handleConnection(conn)
	}

	return nil
}
