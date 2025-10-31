package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Простая запись в файл")

	data := []byte("Это первая строка лога.\nА это вторая строка лога\n")
	err := os.WriteFile("log.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Файл log.txt успешно создан и записан")
}
