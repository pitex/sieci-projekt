package joinservice

import "net"
import "log"

type ServerFullError struct {
	Address	string
}

func (err ServerFullError) Error() string {
	return "Server " + err.Address + " is already full!"
}

type Server struct {
	Address		string
	Parent		net.Conn
	Children	[]net.Conn
	ChildNumber	int
	Root		bool
}

func NewServer(ip string, parent string, capacity int, root bool) *Server{
	var socket net.Conn
	var err error

	if !root {
		socket, err = net.Dial("tcp",parent)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &Server{ip,socket,make([]net.Conn, capacity-1),0,root}
}

// Building chart script.
func (s *Server) BuildChart() {
	
}

// We receive 'data fragment' of chart script from our parent, we have to handle the data, build our own chart and send it to children.
func (s *Server) HandleChart() {
	
}

// ROOT ONLY - We create chart script and send it to children so they can update their charts
func (s *Server) CreateChart() {
	
}

// ROOT ONLY - When we know that there is a new machine pending. 
// We need to find it place in out net and send the information about it to our children.
// We also need to create updated chart and send it to children, too.
func (s *Server) HandleNewMachine(msg string) {
	
}

// Determines how to react for a SIM message depending on its type.
//TODO na razie info zwrotne idzie do wszystkich dzieci
func (s *Server) SIPMessageReaction(msg string) {
	switch ExtractType(msg) {
		//case "INF" : 
		case "BLD", "REQ" :
			s.AskChildren(InfoMsg(msg))
			if s.Root {
				s.HandleNewMachine(msg)
			} else {
				s.TellParent(msg)
			}
		case "TRA" :
			s.TellParent(InfoMsg(msg))
			s.HandleChart()
		case "FND" :
			s.TellParent(InfoMsg(msg))
			pa, ca := FNDInterpretation(msg)
			if Address == pa { 
				AddChild(ca)
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
