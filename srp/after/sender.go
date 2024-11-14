package after

import "fmt"

type Sender struct {
	Sender   *User
	Receiver *User
	Message  string
}

func (s *Sender) SendViaWhatsapp() bool {
	fmt.Println("Send Message Via Whatsapp")
	fmt.Println("Sender: ", s.Sender.phone)
	fmt.Println("Receiver: ", s.Receiver.phone)
	fmt.Println("Message: ", s.Message)
	return true
}

func (s *Sender) SendViaTelegram() bool {
	fmt.Println("Send Message Via Telegram")
	fmt.Println("Sender: ", s.Sender.phone)
	fmt.Println("Receiver: ", s.Receiver.phone)
	fmt.Println("Message: ", s.Message)
	return true
}
