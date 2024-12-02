package ocp

import (
	"fmt"
	"time"
)

type MessageTemplate struct {
	StartedAt time.Time
	EndedAt   time.Time
}

type MessageTemplateWebinar struct {
	MessageTemplate
	Title  string
	Author string
}

type MessageTemplateCompetition struct {
	MessageTemplate
	Name  string
	Level string
}

func (m *MessageTemplateWebinar) Create() string {
	m = &MessageTemplateWebinar{
		MessageTemplate: MessageTemplate{
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(1 * time.Hour),
		},
		Title:  "Generative AI for Developers",
		Author: "Akshay Beniwal",
	}
	res := fmt.Sprintf("Webinar %s by %s will be starting at %s and till %s", m.Title, m.Author, m.StartedAt, m.EndedAt)
	return res
}

func (m *MessageTemplateCompetition) Create() string {
	m = &MessageTemplateCompetition{
		MessageTemplate: MessageTemplate{
			StartedAt: time.Now(),
			EndedAt:   time.Now().Add(2 * time.Hour),
		},
		Name:  "Competitive Programming",
		Level: "Medium",
	}
	res := fmt.Sprintf("Competition %s, level %s akan dimulai pada %s sampai %s", m.Name, m.Level, m.StartedAt, m.EndedAt)
	return res
}
