package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"simple-library/domain"
)

// SaveBooksToCSV сохраняет список книг в CSV файл
func SaveBooksToCSV(books []*domain.Book, filename string) error {
	// Создаем или перезаписываем файл
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %v", err)
	}
	defer file.Close()

	// Создаем CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Записываем заголовок
	header := []string{"ID", "Title", "Author", "Year", "IsIssued", "ReaderID"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("ошибка записи заголовка: %v", err)
	}

	// Записываем данные книг
	for _, book := range books {
		// Конвертируем ReaderID в строку (может быть nil)
		readerIDStr := ""
		if book.ReaderID != nil {
			readerIDStr = strconv.Itoa(*book.ReaderID)
		}

		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			strconv.FormatBool(book.IsIssued),
			readerIDStr,
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("ошибка записи данных книги: %v", err)
		}
	}

	fmt.Printf("Книги успешно сохранены в файл: %s\n", filename)
	return nil
}

// LoadBooksFromCSV загружает книги из CSV файла
func LoadBooksFromCSV(filename string) ([]*domain.Book, error) {
	// Открываем файл
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	// Создаем CSV reader
	reader := csv.NewReader(file)
	
	// Читаем все записи
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения CSV: %v", err)
	}

	// Проверяем, что файл не пустой
	if len(records) < 2 {
		return nil, fmt.Errorf("файл не содержит данных или заголовка")
	}

	var books []*domain.Book

	// Пропускаем заголовок и обрабатываем данные
	for i, record := range records[1:] {
		// Проверяем количество полей
		if len(record) != 6 {
			return nil, fmt.Errorf("некорректное количество полей в строке %d", i+2)
		}

		// Парсим ID
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга ID в строке %d: %v", i+2, err)
		}

		// Парсим Year
		year, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга Year в строке %d: %v", i+2, err)
		}

		// Парсим IsIssued
		isIssued, err := strconv.ParseBool(record[4])
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга IsIssued в строке %d: %v", i+2, err)
		}

		// Обрабатываем ReaderID (может быть пустым)
		var readerID *int
		if record[5] != "" {
			rid, err := strconv.Atoi(record[5])
			if err != nil {
				return nil, fmt.Errorf("ошибка парсинга ReaderID в строке %d: %v", i+2, err)
			}
			readerID = &rid
		}

		// Создаем книгу
		book := &domain.Book{
			ID:       id,
			Title:    record[1],
			Author:   record[2],
			Year:     year,
			IsIssued: isIssued,
			ReaderID: readerID,
		}

		books = append(books, book)
	}

	fmt.Printf("Книги успешно загружены из файла: %s\n", filename)
	return books, nil
}