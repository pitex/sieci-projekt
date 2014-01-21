package sip

import (
	"fmt"
	"../"
	"net"
	"strings"
)

//	Message type used in Simple Information Protocol.
type Message struct {
	//	Type of the message, must be one of:
	//	INF - 'yes, I received your msg'
	// 
	//	TRA - 'after you receive this message I will begin sending you parted file', sent from parent to child only
	// 
	//	REQ - 'I received a request of a new machine' sent to parent only, data fields: ip, capacity
	//	FND - 'I am sending a message about who should be new machine's parent' sent do child only, data fields: parent, child 
	Type	string

	//	Additional data for message, string form is given as pairs KEY=VALUE separated with commas.
	Data	map[string]string
	Error	string
}

func ParseDataToString(data map[string]string) string {
	result := ""

	for k, v := range data {
		result = result + k
		result = result + "="
		result = result + v
	}

	return result
}

//	Returns string representation of the message.
func (msg *Message) ToString() string {
	return fmt.Sprintf("%s%s%s%s%s", msg.Type, protocols.GetSep(), ParseDataToString(msg.Data), protocols.GetSep(), msg.Error)
}

func GetDataSep() string {
	return ","
}

// Returns type of given string representing message. 
func ExtractType(msg string) (string) {
	return msg[:3]
}

// Returns string representing feedback message.
func InfoMsg(msg string) (string) {
	var result string
	result = "INF" + msg[3:]
	return result
}

/* do not delete, may be useful
// Extracts parent's and new child's addresses
func FNDInterpretation(msg string) (string, string) {
	var parent string
	var child string
	//TODO depending on message format
	splited_msg = Split(msg, protocols.GetSep())
	data = Split(splited_msg[1], GetDataSep())
	
	if Contains(data[0], 'parent') {
		parent = data[0][strings.Index(data[0], '='):]
		child = data[1][strings.Index(data[1], '='):]
	} else {
		child = data[0][strings.Index(data[0], '='):]
		parent = data[1][strings.Index(data[1], '='):]
	}

	return parent, child
}*/

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

// Returns struct Message while given a string describing it.
func GetMessage(msg string) (Message){
	splited_msg := strings.Split(msg, protocols.GetSep())
	
	return Message{splited_msg[0], InterpreteData(splited_msg[1]), splited_msg[2]}
}

//	Performs request msg through socket and returns response.
func Request(socket net.Conn, msg Message) (Message, error) {
	//	NOT YET IMPLEMENTED
	return Message{}, nil
}