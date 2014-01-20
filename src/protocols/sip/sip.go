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
	//	FND - 'I am sending a message about who should be new machine's parent' sent do child only
	Type	string
	
	//	Additional data for message given as pairs KEY=VALUE separated with commas.
	Data	string
	Error	string
}

//	Returns string representation of the message.
func (msg *Message) ToString() string {
	return fmt.Sprintf("%s%s%s%s%s", msg.Type, protocols.GetSep(), msg.Data, protocols.GetSep(), msg.Error)
}