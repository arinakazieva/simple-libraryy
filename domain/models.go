package domain

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус: %s", b.Title, b.Author, b.Year, status)
}

// Метод для выдачи книги читателю
func (b *Book) IssueBook(reader *Reader) {
	if !reader.IsActive {
		fmt.Printf("Читатель %s %s не активен и не может получить книгу.\n", reader.FirstName, reader.LastName)
		return
	}

	if b.IsIssued {
		fmt.Printf("Книга '%s' уже выдана.\n", b.Title)
		return
	}

	b.IsIssued = true
	b.ReaderID = &reader.ID
	fmt.Printf("Книга '%s' выдана читателю %s %s.\n", b.Title, reader.FirstName, reader.LastName)
}

// Метод для возврата книги
func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга '%s' уже в библиотеке.\n", b.Title)
		return
	}

	b.IsIssued = false
	b.ReaderID = nil
	fmt.Printf("Книга '%s' возвращена в библиотеку.\n", b.Title)
}

// Метод для деактивации читателя
func (r *Reader) Deactivate() {
	r.IsActive = false
	fmt.Printf("Читатель %s %s деактивирован.\n", r.FirstName, r.LastName)
}

// Метод для активации читателя
func (r *Reader) Activate() {
	r.IsActive = true
	fmt.Printf("Читатель %s %s активирован.\n", r.FirstName, r.LastName)
}

// Метод для вывода информации о читателе
func (r Reader) String() string {
	status := "активен"
	if !r.IsActive {
		status = "не активен"
	}
	return fmt.Sprintf("%s %s (ID: %d), статус: %s", r.FirstName, r.LastName, r.ID, status)
}

// Метод: читатель берет книгу
func (r *Reader) AssignBook(book *Book) {
	fmt.Printf("Читатель %s %s взял книгу '%s'.\n", r.FirstName, r.LastName, book.String())
}
