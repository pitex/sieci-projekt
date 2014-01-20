package sip

import "fmt"
import "../"

//	Message type used in Simple Information Protocol.
type Message struct {
	//	Type of the message, must be one of:
	//	INF - 'yes, I received your msg'
	// 
	//	BLD - 'begin building tree chart' - msg sent to parent only
	//	TRA - 'after you receive this message I will begin sending you parted file', sent from parent to child only
	// 
	//	REQ - 'I received a request of a new machine' sent to parent only
	//	FND - 'I am sending a message about who should be new machine's parent' sent do child only, data fields: parent, child 
	Type	string
	
	//	Additional data for message given as pairs KEY=VALUE separated with commas.
	Data	string
	Error	string
}

//	Returns string representation of the message.
func (msg *Message) ToString() string {
	return fmt.Sprintf("%s%s%s%s%s", msg.Type, protocols.GetSep(), msg.Data, protocols.GetSep(), msg.Error)
}

func GetSep() string {
	return "|"
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

// Extracts parent's and new child's addresses
func FNDInterpretation(msg string) (string, string) {
	var parent string
	var child string
	//TODO depending on message format
	splited_msg = Split(msg, protocols.GetSep())
	data = Split(splited_msg[1], GetDataSep())
	
	if Contains(data[0], 'parent') {
		parent = data[0][Index(data[0], '='):]
		child = data[1][Index(data[1], '='):]
	} else {
		child = data[0][Index(data[0], '='):]
		parent = data[1][Index(data[1], '='):]
	}

	return parent, child
}