package main

import "fmt"

func main() {
	myLibrary := Library{
		Books: []Book{
			{ID: 1, Title: "Война и мир", Author: "Лев Толстой", Year: 1869, IsIssued: false},
			{ID: 2, Title: "Преступление и наказание", Author: "Федор Достоевский", Year: 1866, IsIssued: false},
		},
		Readers: []Reader{
			{ID: 1, Name: "Иван Иванов"},
			{ID: 2, Name: "Петр Петров"},
		},
	}

	fmt.Println("=== Тестирование выдачи книги ===")
	err := myLibrary.IssueBookToReader(1, 1)
	if err != nil {
		fmt.Printf("Ошибка при выдаче книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно выдана!")
	}

	fmt.Println("\n=== Тестирование возврата книги ===")
	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("Ошибка при возврате книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно возвращена в библиотеку!")
	}

	fmt.Println("\n=== Тестирование повторного возврата книги ===")
	err = myLibrary.ReturnBook(1)
	if err != nil {
		fmt.Printf("Ошибка при возврате книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно возвращена в библиотеку!")
	}

	fmt.Println("\n=== Тестирование с несуществующей книгой ===")

	err = myLibrary.ReturnBook(999)
	if err != nil {
		fmt.Printf("Ошибка при возврате книги: %v\n", err)
	} else {
		fmt.Println("Книга успешно возвращена в библиотеку!")
	}

	fmt.Println("\n=== Тестирование функции GetPortFromConfig ===")

	configWithPort := map[string]string{"PORT": "8080", "HOST": "localhost"}
	configWithoutPort := map[string]string{"HOST": "localhost", "TIMEOUT": "30s"}

	port1, err1 := GetPortFromConfig(configWithPort)
	if err1 != nil {
		fmt.Printf("Ошибка: %v\n", err1)
	} else {
		fmt.Printf("PORT найден: %s\n", port1)
	}

	port2, err2 := GetPortFromConfig(configWithoutPort)
	if err2 != nil {
		fmt.Printf("Ошибка: %v\n", err2)
	} else {
		fmt.Printf("PORT найден: %s\n", port2)
	}
}
