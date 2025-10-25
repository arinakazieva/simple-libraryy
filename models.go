package main

import "fmt"

type Library struct {
	Books        []*Book
	Readers      []*Reader
	lastBookID   int
	lastReaderID int
}

type Book struct {
	ID       int
	Year     int
	Title    string
	Author   string
	IsIssued bool
	ReaderID int
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (r *Reader) String() string {
	return fmt.Sprintf("%s %s (ID: %d)", r.FirstName, r.LastName, r.ID)
}

func (b *Book) String() string {
	status := "в наличии"
	if b.IsIssued {
		status = "выдана"
	}
	return fmt.Sprintf("'%s' by %s (%d) [%s]", b.Title, b.Author, b.Year, status)
}

// создает нового читателя
func (lib *Library) AddReader(firstName, lastName string) (*Reader, error) {
	lib.lastReaderID++
	newReader := &Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	lib.Readers = append(lib.Readers, newReader)
	return newReader, nil
}

// создает новую книгу
func (lib *Library) AddBook(title, author string, year int) (*Book, error) {
	// Проверяем, нет ли уже книги с таким же автором и названием
	for _, book := range lib.Books {
		if book.Title == title && book.Author == author {
			return nil, fmt.Errorf("книга '%s' автора '%s' уже существует в библиотеке", title, author)
		}
	}

	lib.lastBookID++
	newBook := &Book{
		ID:       lib.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
		ReaderID: 0,
	}
	lib.Books = append(lib.Books, newBook)
	return newBook, nil
}

// находит книгу по ID
func (lib *Library) FindBookByID(id int) (*Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

// находит читателя по ID
func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

// выдает книгу читателю
func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	if book.IsIssued {
		return fmt.Errorf("книга '%s' уже выдана", book.Title)
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	if !reader.IsActive {
		return fmt.Errorf("читатель %s не активен", reader)
	}

	book.IsIssued = true
	book.ReaderID = readerID

	return nil
}

// возвращает книгу в библиотеку
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
