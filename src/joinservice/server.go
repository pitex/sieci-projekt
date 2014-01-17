package joinservice

import "net"
import "string"

type ServerFullError struct {
	address	string
}

func (err *ServerFullError) Error() string {
	return "Server " + err.address + " is already full!"
}

type Server struct {
	address		string
	parent		net.Conn
	children	[]net.Conn
	childNumber	int
}

func (s *Server) AskChildren(msg string) {
	for child := range s.children {
		child.Write(string.Bytes(msg))
	}
}

func (s *Server) TellParent(msg string) {
	s.parent.Write(string.Bytes(msg))
}

func (s *Server) AddChild(address string) error {
	if len(children) == childNumber {
		return ServerFullError(s.address)
	}

	conn, err := net.Dial("tcp", address + ":666")

	if err != nil {
		return err
	}

	children[childNumber] = conn
	childNumber++

	return nil
}