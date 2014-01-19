package joinservice

import "net"
import "fmt"
import "strings"
import "../protocols/sip"

type Client struct {
	Address		string
	KnownIp		string
	Capacity	int
}

//	Connects client to network containing knownIp
func (c *Client) Connect() error {
	//	Setup connection
	conn, err := net.Dial("tcp", c.KnownIp) 
	if err != nil {
		return err
	}


	//	Create request
	request := new(sip.Message)
	request.Type = "FIND"
	request.Data = fmt.Sprintf("IP=%s,CAP=%d", c.Address, c.Capacity)

	//	Send request
	n, err := conn.Write(strings.Bytes(request.ToString()))

	//	Check if sending was successful
	if err != nil {
		return err
	}


	//	Get response
	var respBytes = make([]byte, 1024)
	n, err := conn.Read(respBytes)

	//	Check if there was error with read
	if err != nil {
		return err
	}

	//	Save response as string
	response := string(respBytes[:n])

	return nil
}
