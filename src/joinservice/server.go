package joinservice

import (
	"net"
	"log"
	"./tree"
	"os"
	"../protocols/sip"
	"strconv"
	// "../protocols/stp"
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

	//	Connection to parent in network.
	Parent		net.Conn

	//	Connections to children.
	Children	[]net.Conn

	//	How many children already exist.
	ChildNumber	int

	// If the server is root, it is pointer to node representing it, otherwise it is nil.
	Root		*tree.Node
}

//	Creates a new server with giben ip and capacity:
//	-	if root is false, it gets connected to parent,
//	-	else initializes tree node representing network structure.
func NewServer(ip string, parent string, capacity int, root bool) *Server{
	var socket net.Conn
	var err error

	if !root {
		socket, err = net.Dial("tcp",parent)
		if err != nil {
			log.Fatal(err)
		}
	}

	if !root {
		return &Server{ip, socket, make([]net.Conn, capacity-1), 0, nil}
	}
	return  &Server{ip, socket, make([]net.Conn, capacity-1), 0, tree.NewNode(ip, capacity)}
}

// We receive the chart script from our parent, we have to handle the data and send it to children.
func (s *Server) HandleChartTransfer() {
}

// ROOT ONLY - travelling tree and adding nodes' description into our script
func (s *Server) BuildChart() {
	f, _ := os.OpenFile("../resources/chart.html", os.O_APPEND, os.ModeAppend)
	tree.DFS(s.Root, "", f)
	f.Close()
}

// Rewrites input file to output file in APPEND MODE
func RewriteFile(input string, output string) {
	infile, _ := os.Open(input)
	outfile, _ := os.OpenFile(output, os.O_RDWR|os.O_APPEND, 0660)

	defer infile.Close()
	defer outfile.Close()

	for {
		b := make([]byte, 1024)
		read, _ := infile.Read(b)
		
		outfile.Write(b[:read])

		if read < 1024 {
			break
		}
	}
}

// ROOT ONLY - We create chart script and send it to children so they can update their charts
func (s *Server) CreateChart() {
	os.Create("../resources/chart.html")
	RewriteFile("../resources/chart_beg", "../resources/chart.html")
	s.BuildChart()
	RewriteFile("../resources/chart_end", "../resources/chart.html")
}

// ROOT ONLY - When we know that there is a new machine pending. 
// We need to find it place in out net and send the information about it to our children.
// We also need to create updated chart and send it to children, too.
func (s *Server) HandleNewMachine(msg string) {
	DataMap := sip.InterpreteData(sip.ExtractData(msg))
	fatherNode, _ := tree.FindSolution(s.Root, -1)
	cap, _ := strconv.Atoi(DataMap["capacity"])
	tree.AddNewChild(fatherNode, tree.NewNode(DataMap["ip"], cap))
	newMes := sip.FND(fatherNode.IP, DataMap["ip"])
	s.AskChildren(newMes.ToString())
}

// Determines how to react for a SIM message depending on its type.
//TODO na razie info zwrotne idzie do wszystkich dzieci
func (s *Server) SIPMessageReaction(msg string) {
	switch sip.ExtractType(msg) {
		case "BLD", "REQ" :
			s.AskChildren(sip.InfoMsg(msg))
			if s.Root != nil {
				s.HandleNewMachine(msg)
			} else {
				s.TellParent(msg)
			}
		case "TRA" :
			s.TellParent(sip.InfoMsg(msg))
			s.HandleChartTransfer()
		case "FND" :
			s.TellParent(sip.InfoMsg(msg))
			DataMap := sip.InterpreteData(sip.ExtractData(msg))

			if s.Address == DataMap["parent"] { 
				s.AddChild(DataMap["child"])
				break
			}
			s.AskChildren(msg)
	}
}

// Sending msg to all children.
func (s *Server) AskChildren(msg string) {
	byteMsg := []byte(msg)

	for _, child := range s.Children {
		child.Write(byteMsg)
	}
}

// Sending msg to parent.
func (s *Server) TellParent(msg string) {
	byteMsg := []byte(msg)

	s.Parent.Write(byteMsg)
}

// When we receive information that we are to create a connection with new machine, it becomes our child.
func (s *Server) AddChild(address string) error {
	if len(s.Children) == s.ChildNumber {
		return ServerFullError{s.Address}
	}

	conn, err := net.Dial("tcp", address + ":666")

	if err != nil {
		return err
	}

	s.Children[s.ChildNumber] = conn
	s.ChildNumber++

	return nil
}

func (s *Server) Start() error {
	return nil
}
