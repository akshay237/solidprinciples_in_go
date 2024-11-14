package before

import "fmt"

type MessageHandler interface {
	Create() *Message
	Send(msg *Message) bool
}

type Message struct {
	body     string
	sender   int64
	receiver int64
}

func (m *Message) Create() *Message {
	m = &Message{
		body:     "Hi, Akshay !",
		sender:   9876543210,
		receiver: 9753186420,
	}

	fmt.Println("Message Created")
	fmt.Printf("%v\n", m)
	return m
}

func (m *Message) Send(msg *Message) bool {
	fmt.Println("Send Message")
	fmt.Println("Sender:", m.sender)
	fmt.Println("Receiver:", m.receiver)
	return true
}
