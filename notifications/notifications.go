package notifications

import "fmt"

// интерфейс для уведомлений
type Notifier interface {
	SendNotification(message string) error
}

// реализация уведомлений по email
type EmailNotifier struct {
	Email string
}

func (e *EmailNotifier) SendNotification(message string) error {
	fmt.Printf("Отправка email на %s: %s\n", e.Email, message)
	return nil
}

//реализация SMS уведомлений
type SMSNotifier struct {
	Phone string
}

func (s *SMSNotifier) SendNotification(message string) error {
	fmt.Printf("Отправка SMS на %s: %s\n", s.Phone, message)
	return nil
}
