package joinservice

import (
	"net"
	"strings"
	"../protocols/sip"
	"../protocols"
	"log"
)

type Client struct {
	//	Address of the client computer.
	Address		string

	//	Address of a computer in the existing network.
	KnownIp		string

	//	How many connections can be handled.
	Capacity	int

	//	Is this computer first in the network.
	IsRoot		bool
}

//	Creates a new Client object using given parameters and returns it.
func NewClient(address string, capacity int, knownIp string, root bool) *Client {
	log.Println("Creating new client")

	return &Client{address,knownIp,capacity,root}
}

//	Asks computer from network where to connect to.
//	Returns address of the computer to which client should connect.
func (c *Client) Connect() (string, error) {
	log.Printf("Connecting client %s using %s\n", c.Address, c.KnownIp)

	//	Setup connection
	log.Printf("Creating socket to %s\n", c.KnownIp)
	conn, err := net.Dial("tcp", c.KnownIp+":666") 
	if err != nil {
		return "", err
	}

	//	Create request
	log.Println("Creating message")
	request := sip.Message{}
	request.Type = "REQ"
	request.AddData("ip", c.Address)
	request.AddData("capacity", string(c.Capacity))

	//	Send request
	err = sip.Request(conn, request)

	//	Check if sending was successful
	if err != nil {
		return "", err
	}

	//	Get response
	log.Println("Waiting for response")
	var respBytes = make([]byte, 1024)
	n, err := conn.Read(respBytes)

	//	Check if there was error with read
	if err != nil {
		return "", err
	}

	//	Save response
	log.Println("Saving response")
	response := strings.Split(string(respBytes[:n]), protocols.GetSep())

	//	Get received data
	log.Printf("Getting data from response: %s\n", response)
	data := strings.Split(response[1], sip.GetDataSep())

	//	Find server address
	log.Printf("Looking for parent IP in data: %s\n", data)
	for _, s := range data {
		kv := strings.Split(s, "=")
		if (kv[0] == "parent") {
			return kv[1], nil
		}
	}

	log.Println("No IP received")
	return "", nil
}
