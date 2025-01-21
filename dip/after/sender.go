package after

import "fmt"

type ISenderSocmed interface {
	SendWhatsapp()
	SendTelegram()
}

type SenderSocmed struct {
	Sender   *User
	Receiver *User
	Message  string
}

type ISenderEmail interface {
	SendEmail()
}

type SenderEmail struct {
	Sender   *User
	Receiver *User
	Subject  string
	Body     string
}

func (s *SenderSocmed) SendWhatsapp() bool {
	fmt.Println("Send via Whatsapp")
	fmt.Println("Sender phone no:", s.Sender.phone)
	fmt.Println("Reciever phone no:", s.Receiver.phone)
	fmt.Println("Body", s.Message)
	return true
}

func (s *SenderSocmed) SendTelegram() bool {
	fmt.Println("Send via Telegram")
	fmt.Println("Sender Phone Number:", s.Sender.phone)
	fmt.Println("Reciever Phone Number:", s.Receiver.phone)
	fmt.Println("Body:", s.Message)
	return true
}

func (s *SenderEmail) SendEmail() bool {
	fmt.Println("Send via email")
	fmt.Println("Sender:", s.Sender.email)
	fmt.Println("Reciever:", s.Receiver.email)
	fmt.Println("Subject:", s.Subject)
	fmt.Println("Body:", s.Body)
	return true
}
