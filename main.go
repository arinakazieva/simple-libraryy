package main

import "fmt"

func main() {
	user1 := Reader{
		ID:        1,
		FirstName: "Арина",
		LastName:  "Казиева",
		IsActive:  true,
	}

	book1 := Book{
		ID:       1,
		Year:     1867,
		Title:    "Война и мир",
		Author:   "Лев Толстой",
		IsIssued: false,
	}
	fmt.Println(user1)
	fmt.Println(book1)
	book1.IssueBook(&user1)
	fmt.Println(book1)
	book1.ReturnBook()
	fmt.Println(book1)
	user1.AssignBook(&book1)

	fmt.Println("------------------------------------")

	n := []Notifier{}
	em := EmailNotifier{EmailAddress: "arinakazieva01@gmail.com"}
	sms := SMSNotifier{PhoneNumber: "+79633787686"}
	n = append(n, em, sms)
	for i := 0; i < len(n); i++ {
		n[i].Notify("Ваша книга просрочена!")
	}
}
