package sip

import "fmt"

/*
	possible message types:
	INF - 'yes, I received your msg'

	BLD - 'begin building tree chart' - msg sent to parent only
	TRA - 'after you receive this message I will begin sending you parted file', sent from parent to child only

	REQ - 'I received a request of a new machine' sent to parent only
	FND - 'I am sending a message about who should be new machine's parent' sent do child only
*/
type Message struct {
	Type	string
	Data	string
	Error	string
}

func (msg *Message) ToString() string {
	return fmt.Sprintf("%s|%s|%s", msg.Type, msg.Data, msg.Error)
}