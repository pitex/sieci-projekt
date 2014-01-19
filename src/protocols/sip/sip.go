package sip

import "fmt"

type Message struct {
	Type	string
	Data	string
	Error	string
}

func (msg *Message) ToString() string {
	return fmt.Sprintf("%s;%s;%s", msg.Type, msg.Data, msg.Error)
}