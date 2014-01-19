package joinservice

import "net"

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

//TODO na razie info zwrotne idzie do wszystkich dzieci
func (s *Server) SIPMessageReaction(msg string) {
	switch ExtractType(msg) {
		//case "INF" : 
		case "BLD" :
			s.AskChildren(InfoMsg(msg))
		case "TRA" :
			s.TellParent(InfoMsg(msg))
		case "REQ" :
			s.AskChildren(InfoMsg(msg))
			s.TellParent(msg)
		case "FND" :
			s.TellParent(InfoMsg(msg))
			// if { 
			// 	AddChild(...)
			// }
			s.AskChildren(msg)
	}
}

func (s *Server) AskChildren(msg string) {
	byteMsg := []byte(msg)

	for _, child := range s.Children {
		child.Write(byteMsg)
	}
}

func (s *Server) TellParent(msg string) {
	byteMsg := []byte(msg)

	s.Parent.Write(byteMsg)
}

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
