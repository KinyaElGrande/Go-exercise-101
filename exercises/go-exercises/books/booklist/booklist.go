package booklist

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Book exported
type Book struct {
	Author   string
	Title    string
	Filename string
}

const (
	dataFileName = "_booklist.csv"
)

// Service is exported
type Service struct {
	books []*Book
}

//NewService is exported
func NewService(dataDirectory string) (*Service, error) {
	books, err := loadFile(dataDirectory)
	if err != nil {
		return nil, err
	}
	return &Service{books: books}, nil
}

// GetByAuthor returns all the books by the given author.
func (s *Service) GetByAuthor(authorName string) []*Book {
	var books []*Book
	for _, b := range s.books {
		if strings.ToLower(authorName) == strings.ToLower(b.Author) {
			books = append(books, b)
		}
	}
	return books
}

func loadFile(dataDirectory string) ([]*Book, error) {

	f, err := os.Open(filepath.Join(dataDirectory, dataFileName))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.ReuseRecord = true
	reader.FieldsPerRecord = 0
	reader.TrimLeadingSpace = true
	_, err = reader.Read() // header row
	if err != nil {
		return nil, err
	}

	var readErr error
	var books []*Book
	for {
		row, readErr := reader.Read()
		if readErr != nil {
			break
		}
		books = append(books, &Book{
			Author:   row[0],
			Filename: row[1],
			Title:    row[2],
		})

	}

	if readErr == io.EOF {
		readErr = nil
	}
	if readErr != nil {
		return nil, readErr
	}

	return books, nil
}
