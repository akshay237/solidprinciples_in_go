package ocp

import "fmt"

func Run() {
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Run OCP (Open Closed Principle)")
	messageTemplateCompetition := MessageTemplateCompetition{}
	message := &Message{
		MessageTemplateCompetition: &messageTemplateCompetition,
	}

	messagePayload := message.Create()
	fmt.Println()

	user := &User{}
	sender := &Sender{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Message:  messagePayload,
	}

	sender.SendViaWhatsapp()
	fmt.Println()
	sender.SendViaTelegram()
}
