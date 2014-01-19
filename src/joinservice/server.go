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
