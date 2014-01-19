package stp

type Message struct {
	Data	string
	Error	string
}

func (msg *Message) ToString() string {
	return fmt.Sprintf("%s|%s", msg.Data, msg.Error)
}