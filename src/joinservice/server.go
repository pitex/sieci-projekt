package joinservice

import "net"
import "log"
import "tree"

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

// We receive the chart script from our parent, we have to handle the data and send it to children.
func (s *Server) HandleChartTransfer() {
	
}

// Data about node in format required by Google Charts.
func NodeFormatted(node *tree.Node, string parent, string ToolTip) string {
	return fmt.Sprintf("['%s','%s','%s'],", *node.IP, parent, ToolTip)
}

// ROOT ONLY - travelling tree and adding nodes' description into our script
func (s *Server) BuildChart() {

}

// Rewrites input file to output file in APPEND MODE
func RewriteFile(input string, output string) {
	
}

// ROOT ONLY - We create chart script and send it to children so they can update their charts
func (s *Server) CreateChart() {
	os.Create("../resources/chart.html")
	RewriteFile("../resources/chart_beg", "../resources/chart.html")
	BuildChart()
	RewriteFile("../resources/chart_end", "../resources/chart.html")
}

// ROOT ONLY - When we know that there is a new machine pending. 
// We need to find it place in out net and send the information about it to our children.
// We also need to create updated chart and send it to children, too.
func (s *Server) HandleNewMachine(msg string) {
	
}

// Determines how to react for a SIM message depending on its type.
//TODO na razie info zwrotne idzie do wszystkich dzieci
func (s *Server) SIPMessageReaction(msg string) {
	switch sip.ExtractType(msg) {
		//case "INF" : 
		case "BLD", "REQ" :
			s.AskChildren(sip.InfoMsg(msg))
			if s.Root {
				s.HandleNewMachine(msg)
			} else {
				s.TellParent(msg)
			}
		case "TRA" :
			s.TellParent(sip.InfoMsg(msg))
			s.HandleChart()
		case "FND" :
			s.TellParent(sip.InfoMsg(msg))
			pa, ca := sip.FNDInterpretation(msg)
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
