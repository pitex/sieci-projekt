package stp

import (
	"fmt"
	"../"
)

type Message struct {
	Data	string
	Error	string
}

func (msg *Message) ToString() string {
	return fmt.Sprintf("%s%s%s", msg.Data, protocols.GetSep(), msg.Error)
}