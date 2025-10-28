package main

import (
	"fmt"
	"simple-library/library"
)

func main() {
	myLibrary := library.New()

	// Добавляем читателя
	reader := myLibrary.AddReader("Арина", "Казиева")
	fmt.Println("Зарегистрирован новый читатель:", reader.FirstName, reader.LastName)

	book := myLibrary.AddBook("Александр Толстой", "Александр Пушкин", 1833)
	fmt.Printf("Книга '%s' успешно добавлена\n", book.Title)

	// Выдаем книгу читателю
	myLibrary.IssueBookToReader(book.ID, reader.ID)
	fmt.Println("Состояние книги после выдачи:", book)

	myLibrary.PrintStats()
}
