package library

import (
	"fmt"
	"simple-library/domain"
	"simple-library/storage"
)

type Library struct {
	Books        map[int]*domain.Book
	Readers      map[int]*domain.Reader
	lastBookID   int
	lastReaderID int
}

// Фабричная функция для создания новой библиотеки
func New() *Library {
	return &Library{
		Books:        make(map[int]*domain.Book),
		Readers:      make(map[int]*domain.Reader),
		lastBookID:   0,
		lastReaderID: 0,
	}
}

// Метод для добавления читателя
func (l *Library) AddReader(firstName, lastName string) *domain.Reader {
	l.lastReaderID++
	reader := &domain.Reader{
		ID:        l.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	l.Readers[reader.ID] = reader
	fmt.Printf("Добавлен читатель: %s %s (ID: %d)\n", firstName, lastName, reader.ID)
	return reader
}

// Метод для добавления книги
func (l *Library) AddBook(title, author string, year int) *domain.Book {
	l.lastBookID++
	book := &domain.Book{
		ID:       l.lastBookID,
		Title:    title,
		Author:   author,
		Year:     year,
		IsIssued: false,
		ReaderID: nil,
	}
	l.Books[book.ID] = book
	fmt.Printf("Добавлена книга: %s (%s, %d)\n", title, author, year)
	return book
}

// Метод для поиска книги по ID
func (l *Library) FindBookByID(id int) *domain.Book {
	return l.Books[id]
}

// Метод для поиска читателя по ID
func (l *Library) FindReaderByID(id int) *domain.Reader {
	return l.Readers[id]
}

// Метод для выдачи книги читателю
func (l *Library) IssueBookToReader(bookID, readerID int) {
	book := l.FindBookByID(bookID)
	reader := l.FindReaderByID(readerID)
	
	if book == nil {
		fmt.Printf("Книга с ID %d не найдена.\n", bookID)
		return
	}
	
	if reader == nil {
		fmt.Printf("Читатель с ID %d не найден.\n", readerID)
		return
	}
	
	book.IssueBook(reader)
}

// Метод для возврата книги
func (l *Library) ReturnBook(bookID int) {
	book := l.FindBookByID(bookID)
	if book == nil {
		fmt.Printf("Книга с ID %d не найдена.\n", bookID)
		return
	}
	
	book.ReturnBook()
}

// Метод для получения всех книг
func (l *Library) GetAllBooks() []*domain.Book {
	books := make([]*domain.Book, 0, len(l.Books))
	for _, book := range l.Books {
		books = append(books, book)
	}
	return books
}

// Метод для получения всех читателей
func (l *Library) GetAllReaders() []*domain.Reader {
	readers := make([]*domain.Reader, 0, len(l.Readers))
	for _, reader := range l.Readers {
		readers = append(readers, reader)
	}
	return readers
}

// Метод для вывода статистики библиотеки
func (l *Library) PrintStats() {
	fmt.Printf("=== Статистика библиотеки ===\n")
	fmt.Printf("Книг в библиотеке: %d\n", len(l.Books))
	fmt.Printf("Читателей: %d\n", len(l.Readers))
	
	issuedBooks := 0
	for _, book := range l.Books {
		if book.IsIssued {
			issuedBooks++
		}
	}
	fmt.Printf("Книг выдано: %d\n", issuedBooks)
}

// SaveToCSV сохраняет все книги библиотеки в CSV файл
func (l *Library) SaveToCSV(filename string) error {
	books := l.GetAllBooks()
	return storage.SaveBooksToCSV(books, filename)
}

// LoadFromCSV загружает книги из CSV файла в библиотеку
func (l *Library) LoadFromCSV(filename string) error {
	books, err := storage.LoadBooksFromCSV(filename)
	if err != nil {
		return err
	}

	// Очищаем текущие книги и заменяем их загруженными
	l.Books = make(map[int]*domain.Book)
	for _, book := range books {
		l.Books[book.ID] = book
		// Обновляем lastBookID если нужно
		if book.ID > l.lastBookID {
			l.lastBookID = book.ID
		}
	}

	return nil
}