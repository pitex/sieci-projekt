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

func Request(socket net.Conn, msg Message) error {
	_, err := socket.Write([]byte(msg.ToString()))
	if err != nil {
		return err
	}

	resp := make([]byte, 2048)

	_, err = socket.Read(resp)

	if err != nil {
		return err
	}

	return nil
}