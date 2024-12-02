package ocp

import "fmt"

type MessageHandler interface {
	Create() *Message
}

type Message struct {
	Body                       string
	MessageTemplateCompetition *MessageTemplateCompetition
	MessageTemplateWebinar     *MessageTemplateWebinar
}

func (m *Message) Create() string {
	var template string
	if m.MessageTemplateCompetition != nil {
		template = m.MessageTemplateCompetition.Create()
	}

	if m.MessageTemplateWebinar != nil {
		template = m.MessageTemplateWebinar.Create()
	}

	m = &Message{
		Body: "Hey Ram !" + template,
	}
	res := m.Body
	fmt.Printf("Created Message: %+v\n", m)
	return res
}
