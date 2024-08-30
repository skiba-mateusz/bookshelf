package store

import (
	"encoding/json"
	"os"
	"time"
)

type Book struct {
	ID				int64 `json:"id"`
	Title			string `json:"title"`
	Author 		string `json:"author"`
	Year    	int `json:"year"`
	Read    	bool `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}

type BookStore struct {
	filename string
	books []Book
}

// Initializes a new BookStore and loads existing books
func NewBookStore(filename string) (*BookStore, error) {
	store := &BookStore{
		filename: filename,
		books: []Book{},
	}

	if err := store.load(); err != nil {
		return nil, err
	}

	return store, nil
}

// Returns the list of books in the store
func (s *BookStore) Books() []Book {
	return s.books
}

// Adds a new book to the store and saves it in the file
func (s *BookStore) Add(title, author string, year int) error {
	book := Book{
		ID: s.nextID(),
		Title: title,
		Author: author,
		Year: year,
		Read: false,
		CreatedAt: time.Now(),
	}

	s.books = append(s.books, book)
	return s.save()
}

// Loads books from the file and populates the store
func (s *BookStore) load() error {
	dataBytes, err := os.ReadFile(s.filename)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			// File doesn't exist, so there are no books to load
			return nil
		default:
			return err
		}
	}

	if len(dataBytes) == 0 {
		return nil
	}

	return json.Unmarshal(dataBytes, &s.books)
}

// Saves list of the books to the file
func (s *BookStore) save() error {
	dataBytes, err := json.MarshalIndent(s.books, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, dataBytes, 0644)
}

// Returns a unique ID
func (s *BookStore) nextID() int64 {
	var maxID int64
	for _, book := range s.books {
		if book.ID > maxID {
			maxID = book.ID
		}
	}

	return maxID + 1
}