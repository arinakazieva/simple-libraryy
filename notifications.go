package main

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
}

func (en EmailNotifier) Notify(message string) {
	println("Отправляю email на", en.EmailAddress+":", "\""+message+"\"")
}

type SMSNotifier struct {
	PhoneNumber string
}

func (sn SMSNotifier) Notify(message string) {
	println("Отправляю SMS на номер", sn.PhoneNumber+":", "\""+message+"\"")
}
