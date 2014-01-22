package stp

import (
	"fmt"
	"../"
	"net"
)

type Message struct {
	Data	string
	Error	string
}

func (msg *Message) ToString() string {
	return fmt.Sprintf("%s%s%s", msg.Data, protocols.GetSep(), msg.Error)
}

//	Sends an stp request through given socket.
func Request(socket net.Conn, msg Message) error {
	_, err := socket.Write([]byte(msg.ToString()))
	if err != nil {
		return err
	}

	resp := make([]byte, 4096)

	_, err = socket.Read(resp)

	if err != nil {
		return err
	}

	return nil
}