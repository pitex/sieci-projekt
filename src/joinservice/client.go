package joinservice

import (
	"net"
	"fmt"
	"strings"
	"../protocols/sip"
	"../protocols"
	"log"
)

type Client struct {
	Address		string
	KnownIp		string
	Capacity	int
	IsRoot		bool
}

func NewClient(myip string, capacity int, ip string, root bool) *Client {
	log.Println("Creating new client")

	return &Client{myip,ip,capacity,root}
}

//	Connects client to network
func (c *Client) Connect() (string, error) {
	log.Printf("Connecting client %s using %s\n", c.Address, c.KnownIp)

	log.Printf("Creating socket to %s\n", c.KnownIp)
	//	Setup connection
	conn, err := net.Dial("tcp", c.KnownIp) 
	if err != nil {
		return "", err
	}

	log.Println("Creating message")
	//	Create request
	request := new(sip.Message)
	request.Type = "FIND"
	request.Data = fmt.Sprintf("IP=%s,CAP=%d", c.Address, c.Capacity)

	byteRequest := []byte(request.ToString())

	log.Printf("Sending message: %s\n", request.ToString())
	//	Send request
	_, err = conn.Write(byteRequest)

	//	Check if sending was successful
	if err != nil {
		return "", err
	}

	log.Println("Waiting for response")
	//	Get response
	var respBytes = make([]byte, 1024)
	n, err := conn.Read(respBytes)

	//	Check if there was error with read
	if err != nil {
		return "", err
	}

	log.Println("Saving response")
	//	Save response
	response := strings.Split(string(respBytes[:n]), protocols.GetSep())

	log.Printf("Getting data from response: %s\n", response)
	//	Get received data
	data := strings.Split(response[1], ",")

	log.Printf("Looking for IP in data: %s\n", data)
	//	Find server address
	for _, s := range data {
		kv := strings.Split(s, "=")
		if (kv[0] == "IP") {
			return kv[1], nil
		}
	}

	log.Println("No IP received")
	return "", nil
}
