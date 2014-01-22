package sip

import (
	"fmt"
	"../"
	"net"
	"strings"
	"log"
)

//	Message type used in Simple Information Protocol.
type Message struct {
	//	Type of the message, must be one of:
	//	INF - 'yes, I received your msg'
	// 
	//	TRA - 'after you receive this message I will begin sending you parted file', sent from parent to child only
	//	END - 'sending is complete'
	// 
	//	REQ - 'I received a request of a new machine' sent to parent only, data fields: ip, capacity
	//	FND - 'I am sending a message about who should be new machine's parent' sent do child only, data fields: parent, child 
	Type	string

	//	Additional data for message, string form is given as pairs KEY=VALUE separated with commas.
	Data	map[string]string
	Error	string
}

//	Adds a k=v pair to data.
func (msg *Message) AddData(k, v string) {
	msg.Data[k] = v
}

//	Removes value associated with key k.
func (msg *Message) RemoveData(k string) {
	delete(msg.Data, k)
}

//	Parses key-value map to a string.
func ParseDataToString(data map[string]string) string {
	result := ""

	for k, v := range data {
		result = result + k + "=" + v + GetDataSep()
	}

	if len(result) > 0 {
		result = result[:len(result)-1]
	}

	return result
}

//	Returns string representation of the message.
func (msg *Message) ToString() string {
	return fmt.Sprintf("%s%s%s%s%s", msg.Type, protocols.GetSep(), ParseDataToString(msg.Data), protocols.GetSep(), msg.Error)
}

//	Returns string separating key-value pairs in data.
func GetDataSep() string {
	return ","
}

//	Returns type of given string representing message. 
func ExtractType(msg string) (string) {
	return strings.Split(msg, protocols.GetSep())[0]
}

//	Returns data as string from string.
func ExtractData(msg string) (string) {
	splited_msg := strings.Split(msg, protocols.GetSep())
	return splited_msg[1]
}

//	Returns string representing feedback message.
func InfoMsg(msg string) (string) {
	var result string
	result = "INF" + msg[3:]
	return result
}

func FND(child string, parent string) (Message) {
	var data = make(map[string]string)
	data["child"] = child
	data["parent"] = parent

	return Message{"FND", data, ""}
}

//	Parses data string to key-value map.
func InterpreteData(data string) map[string]string {
	sp := strings.Split(data, GetDataSep())
	var result = make(map[string]string)

	for i := 0; i < len(sp); i++ {
		ind := strings.Index(sp[i], "=")
		val := sp[i][ind + 1:]
		key := sp[i][:ind]
		result[key] = val
	}

	return result
}

//	Returns struct Message while given a string describing it.
func GetMessage(msg string) (Message){
	splited_msg := strings.Split(msg, protocols.GetSep())
	
	return Message{splited_msg[0], InterpreteData(splited_msg[1]), splited_msg[2]}
}

//	Performs request msg through socket and returns response.
func Request(socket net.Conn, msg Message) (error) {
	log.Printf("Sending message: %s\n",msg.ToString())

	_, err := socket.Write([]byte(msg.ToString()))

	if err != nil {
		return err
	}

	log.Printf("Waiting for response\n")

	resp := make([]byte, 4096)
	var n int

	n, err = socket.Read(resp)

	if err != nil {
		return err
	}

	return nil
}