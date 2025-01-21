package dip

import (
	"fmt"

	"github.com/akshay237/solidprinciples_in_go/dip/after"
)

type Factory struct {
	MessageSocmed after.IMessageSocmed
	MessageEmail  after.IMessageEmail
	User          after.IUser
}

func Run() {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Run DIP (Dependency Inversion Principle)")

	messageTempWebinar := &after.MessageTemplateWebinar{}
	messageSocemed := &after.MessageSocmed{
		MessageTemplate: messageTempWebinar,
	}

	messageEmail := &after.MessageEmail{
		MessageTemplate: messageTempWebinar,
	}
	user := &after.User{}

	factory := Factory{
		MessageSocmed: messageSocemed,
		MessageEmail:  messageEmail,
		User:          user,
	}
	Execute(&factory)
}

func Execute(f *Factory) {
	messageSocmedPayload := f.MessageSocmed.Create()
	messageEmailPayload := f.MessageEmail.Create()
	fmt.Println()

	senderSocmed := &after.SenderSocmed{
		Sender:   f.User.GetSender(),
		Receiver: f.User.GetReceiver(),
		Message:  messageSocmedPayload.Body,
	}
	senderEmail := &after.SenderEmail{
		Sender:   f.User.GetSender(),
		Receiver: f.User.GetReceiver(),
		Subject:  messageEmailPayload.Subject,
		Body:     messageEmailPayload.Body,
	}

	senderSocmed.SendWhatsapp()
	fmt.Println()
	senderSocmed.SendTelegram()
	fmt.Println()
	senderEmail.SendEmail()
}
