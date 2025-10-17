package main

import "fmt"

type Book struct {
	ID       int
	Year     int
	Title    string
	Author   string
	IsIssued bool
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (b *Book) IssueBook(reader *Reader) {
	if b.IsIssued {
		fmt.Printf("Книга '%s' уже выдана\n", b.Title)
		return
	}
	b.IsIssued = true
	fmt.Printf("Книга '%s' выдана читателю %s %s\n", b.Title, reader.FirstName, reader.LastName)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга '%s' уже возвращена\n", b.Title)
		return
	}
	b.IsIssued = false
	fmt.Printf("Книга '%s' возвращена в библиотеку\n", b.Title)
}

func (r *Reader) AssignBook(book *Book) {
	fmt.Printf("Книга '%s' назначена читателю %s %s\n", book.Title, r.FirstName, r.LastName)
}
