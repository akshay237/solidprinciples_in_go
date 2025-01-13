package after

import "fmt"

type Sender struct {
	Sender   *User
	Receiver *User
	Message  string
}

func (s *Sender) SendViaWhatsapp() bool {
	fmt.Println("Send via whatsapp")
	fmt.Println("Sender:", s.Sender.phone)
	fmt.Println("Receiver:", s.Receiver.phone)
	fmt.Println("Body:", s.Message)
	return true
}

func (s *Sender) SendViaTelegram() bool {
	fmt.Println("Send via telegram")
	fmt.Println("Sender:", s.Sender.phone)
	fmt.Println("Receiver:", s.Receiver.phone)
	fmt.Println("Body:", s.Message)
	return true
}
