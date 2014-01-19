package joinservice

import (
	"net"
	"fmt"
	"strings"
	"../protocols/sip"
	"../protocols"
)

type Client struct {
	Address		string
	KnownIp		string
	Capacity	int
	IsRoot		bool
}

func NewClient(myip string, capacity int, ip string, root bool) *Client {
	return &Client{myip,ip,capacity,root}
}

//	Connects client to networkz
func (c *Client) Connect() (string, error) {
	//	Setup connection
	conn, err := net.Dial("tcp", c.KnownIp) 
	if err != nil {
		return "", err
	}



	//	Create request
	request := new(sip.Message)
	request.Type = "FIND"
	request.Data = fmt.Sprintf("IP=%s,CAP=%d", c.Address, c.Capacity)

	byteRequest := []byte(request.ToString())

	//	Send request
	_, err = conn.Write(byteRequest)

	//	Check if sending was successful
	if err != nil {
		return "", err
	}



	//	Get response
	var respBytes = make([]byte, 1024)
	n, err := conn.Read(respBytes)

	//	Check if there was error with read
	if err != nil {
		return "", err
	}

	//	Save response
	response := strings.Split(string(respBytes[:n]), protocols.GetSep())



	//	Get received data
	data := strings.Split(response[1], ",")

	//	Find server address
	for _, s := range data {
		kv := strings.Split(s, "=")
		if (kv[0] == "IP") {
			return kv[1], nil
		}
	}

	return "", nil
}
