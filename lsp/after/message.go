package after

import "fmt"

type MessageHandler interface {
	Create() string
}

type Message struct {
	Body            string
	MessageTemplate IMessageTemplate
}

func (m *Message) Create() string {
	template := m.MessageTemplate.Create()
	m = &Message{
		Body: "Hai malik !" + template,
	}

	res := m.Body

	fmt.Println("Create Message")
	fmt.Printf("%+v\n", m)
	return res
}
