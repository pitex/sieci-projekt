package joinservice

import "net"

type Client struct {
	address		string
	knownIp		string
	capacity	int
}

// Connects client to network containing knownIp
func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.knownIp) 
	if err != nil {
		return err
	}

	var request = string.Bytes("") //TODO define request message
	n, err := conn.Write(request)

	var response = make([]byte, 1024)
	n, err := conn.Read(response)
}