package main

import "fmt"

func main() {
	fmt.Println("Запуск системы управления библиотекой...")

	myLibrary := &Library{}

	fmt.Println("\n--- Наполняем библиотеку ---")

	reader1, err := myLibrary.AddReader("Агунда", "Кокойты")
	if err != nil {
		fmt.Printf("Ошибка при добавлении читателя: %v\n", err)
	} else {
		fmt.Printf("Зарегистрирован новый читатель: %s\n", reader1)
	}

	reader2, err := myLibrary.AddReader("Сергей", "Меняйло")
	if err != nil {
		fmt.Printf("Ошибка при добавлении читателя: %v\n", err)
	} else {
		fmt.Printf("Зарегистрирован новый читатель: %s\n", reader2)
	}

	book1, err := myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	if err != nil {
		fmt.Printf("Ошибка при добавлении книги: %v\n", err)
	} else {
		fmt.Printf("Добавлена новая книга: %s\n", book1)
	}

	book2, err := myLibrary.AddBook("Мастер и Маргарита", "Михаил Булгаков", 1967)
	if err != nil {
		fmt.Printf("Ошибка при добавлении книги: %v\n", err)
	} else {
		fmt.Printf("Добавлена новая книга: %s\n", book2)
	}

	_, err = myLibrary.AddBook("1984", "Джордж Оруэлл", 1949)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка (дубликат книги): %v\n", err)
	}

	fmt.Println("\n--- Библиотека готова к работе ---")
	fmt.Println("Количество читателей:", len(myLibrary.Readers))
	fmt.Println("Количество книг:", len(myLibrary.Books))

	fmt.Println("\n=== Успешная выдача книги ===")
	err = myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Printf("Ошибка выдачи: %v\n", err)
	} else {
		fmt.Printf("Книга '1984' успешно выдана читателю Агунда Кокойты!\n")
	}

	fmt.Println("\n=== Попытка выдать уже выданную книгу ===")
	err = myLibrary.IssueBookToReader(1, 2)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка: %v\n", err)
	} else {
		fmt.Println("Книга выдана!")
	}

	fmt.Println("\n=== Попытка выдать книгу несуществующему читателю ===")
	err = myLibrary.IssueBookToReader(2, 99)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка: %v\n", err)
	} else {
		fmt.Println("Книга выдана!")
	}

	fmt.Println("\n===Успешный возврат книги ===")
	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("Ошибка возврата: %v\n", err)
	} else {
		fmt.Println("Книга '1984' успешно возвращена в библиотеку!")
	}

	fmt.Println("\n===Попытка вернуть книгу, которая уже в библиотеке ===")
	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("Ожидаемая ошибка: %v\n", err)
	} else {
		fmt.Println("Книга возвращена!")
	}

	fmt.Println("\n--- Финальный статус библиотеки ---")
	fmt.Println("Количество читателей:", len(myLibrary.Readers))
	fmt.Println("Количество книг:", len(myLibrary.Books))

	fmt.Println("\nСтатус всех книг:")
	for _, book := range myLibrary.Books {
		status := "в библиотеке"
		if book.IsIssued {
			reader, _ := myLibrary.FindReaderByID(book.ReaderID)
			status = fmt.Sprintf("выдана читателю: %s", reader)
		}
		fmt.Printf("  - %s: %s\n", book.Title, status)
	}

	fmt.Println("\nРабота системы завершена.")
}
