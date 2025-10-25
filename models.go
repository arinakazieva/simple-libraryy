package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID int
}

type Reader struct {
	ID   int
	Name string
}

type Library struct {
	Books   []Book
	Readers []Reader
}

// находим книгу по ID
func (lib *Library) FindBookByID(bookID int) (*Book, error) {
	for i := range lib.Books {
		if lib.Books[i].ID == bookID {
			return &lib.Books[i], nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", bookID)
}

// находим читателя по ID
func (lib *Library) FindReaderByID(readerID int) (*Reader, error) {
	for i := range lib.Readers {
		if lib.Readers[i].ID == readerID {
			return &lib.Readers[i], nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", readerID)
}

// выдаем книгу читателю
func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	if book.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана", book.Title)
	}

	_, err = lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	book.IsIssued = true
	book.ReaderID = readerID

	return nil
}

// возвращаем книгу в библиотеку
func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}

	b.IsIssued = false
	b.ReaderID = 0

	return nil
}

// для возврата книги
func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	err = book.ReturnBook()
	if err != nil {
		return err
	}

	return nil
}
