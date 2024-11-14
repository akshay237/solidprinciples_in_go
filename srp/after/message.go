package after

import "fmt"

type MessageHandler interface {
	Create() *Message
}

type Message struct {
	Body string
}

func (m *Message) Create() string {
	m = &Message{
		Body: "Hi, Akshay !",
	}

	res := m.Body
	fmt.Println("Message Created")
	fmt.Printf("%v\n", m)
	return res
}
