package srp

import (
	"fmt"

	"github.com/akshay237/solidprinciples_in_go/srp/after"
	"github.com/akshay237/solidprinciples_in_go/srp/before"
)

func Run() {
	fmt.Println("Run SRP (Single Responsibility Principle)")
	// Before applying Single Responsibility Princip;e
	fmt.Println("Before Applying SRP:")
	message := &before.Message{}
	messagePayload := message.Create()
	fmt.Println()
	isSent := message.Send(messagePayload)
	if isSent {
		fmt.Println("Message have been Sent")
	}

	// after applying Single Responsibility Principle
	fmt.Println("After applying SRP:")
	msg := &after.Message{}
	msgPayload := msg.Create()

	user := &after.User{}
	sender := &after.Sender{
		Sender:   user.GetSender(),
		Receiver: user.GetReceiver(),
		Message:  msgPayload,
	}

	// send msg via diff channels
	isok := sender.SendViaWhatsapp()
	if !isok {
		sender.SendViaTelegram()
	}

}
