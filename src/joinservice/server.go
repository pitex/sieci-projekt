package joinservice

import "net"
import "strings"

type ServerFullError struct {
	Address	string
}

func (err *ServerFullError) Error() string {
	return "Server " + err.Address + " is already full!"
}

type Server struct {
	Address		string
	Parent		net.Conn
	Children	[]net.Conn
	ChildNumber	int
	Root		bool
}

/* only type of message */
func ExtractType(msg string) (string) {
	return msg[:3]
}

/*	returns feedback message */
func InfoMsg(msg string) (string) {
	var result string
	result = "INF" + msg[3:]
	return result
}

func BuildChart() {
	
}

func HandleNewMachine() {
	
}

func RootReaction(msg string) {
	switch ExtractType(msg) {
		case "BLD": BuildChart()
		case "REQ": HandleNewMachine()
	}
}

func ReceiveChart() {
	
}

//TODO na razie info zwrotne idzie do wszystkich dzieci
func (s *Server) SIPMessageReaction(msg string) {
	switch ExtractType(msg) {
		//case "INF" : 
		case "BLD", "REQ" :
			AskChildren(InfoMsg(msg))
			if Root {
				RootReaction(msg)
			} else {
				TellParent(msg)
			}
		case "TRA" :
			TellParent(InfoMsg(msg))
			ReceiveChart()
		case "FND" :
			TellParent(InfoMsg(msg))
			// if { 
			// 	AddChild(...)
			// }
			AskChildren(msg)
	}
}

func (s *Server) AskChildren(msg string) {
	for child := range s.Children {
		child.Write(strings.Bytes(msg))
	}
}

func (s *Server) TellParent(msg string) {
	s.Parent.Write(strings.Bytes(msg))
}

func (s *Server) AddChild(address string) error {
	if len(children) == childNumber {
		return ServerFullError(s.Address)
	}

	conn, err := net.Dial("tcp", address + ":666")

	if err != nil {
		return err
	}

	children[childNumber] = conn
	childNumber++

	return nil
}
