## Dependency Inversion Principle
* The DIP states that high level modules should not depend on low level modules but should depend on abstractions.
* This principle promotes loose coupling between components, making the code more maintainable and testable.

### Example
* Consider a simple example where a NotificationService sends notifications using an email service
```
type EmailService struct{}

func(e *EmailService) Send(to, message string) {
    // sends an email
}

type NotificationService struct {
    emailService *EmailService
}

func (n *NotificationService) Notify(to, message string) {
    n.emailService.Send(to, message)
}
```

* In this example the Notification Service directly depends on the Email Service making it difficult to switch to another notification method (e.g: SMS) or test the notification service in isolation.
* To follow the dependency inversion principle we should introduce a interface and depends on that instead.

```
type NotificationSender interface {
    Send(to, message string)
}

type EmailService struct {}

func (e *EmailService) Send(to, message string) {
    // sends email notification
}

type SMSService struct {}

func (s *SMSService) Send(to, message string) {
    // sends SMS notification
}

type NotificationService struct {
    sender NotificationSender
}

func(s *NotificationSender) Notify(to, message string) {
    s.sender.Send(to, message)
}
```

* Now the notification depends on the interface and makes the code more flexible to use and we can add a new type as well.