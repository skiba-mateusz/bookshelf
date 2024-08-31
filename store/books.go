package store

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

type SearchQuery struct {
	Title  string
	Author string
	Year   int
	Read   *bool
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

// Searches for book by SearchQuery and returns them
func (s *BookStore) Search(query SearchQuery) []Book {
	results := []Book{}
	
	for _, book := range s.books {
		if query.Title != "" && !strings.Contains(strings.ToLower(book.Title), query.Title) {
			continue
		}
		if query.Author != "" && !strings.Contains(strings.ToLower(book.Author), query.Author) {
			continue
		}
		if query.Year > 0  && book.Year != query.Year {
			continue
		}
		if query.Read != nil && book.Read != *query.Read {
			continue
		}
		results = append(results, book)
	}

	return results
}

// Adds a new book to the store and saves it in the file
func (s *BookStore) Add(book Book) error {
	book.ID = s.nextID()
	s.books = append(s.books, book)
	return s.save()
}

func (s *BookStore) Delete(bookID int64) error {
	updatedBooks := []Book{}
	found := false

	for _, book := range s.books {
		if bookID == book.ID {
			found = true
			continue
		}
		updatedBooks = append(updatedBooks, book)
	}

	if !found {
		return fmt.Errorf("book with ID %d not found", bookID)
	}

	s.books = updatedBooks
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

// Returns a unique ID that is never 0
func (s *BookStore) nextID() int64 {
	var maxID int64
	for _, book := range s.books {
		if book.ID > maxID {
			maxID = book.ID
		}
	}

	if maxID == 0 {
		maxID = 1
	}

	return maxID + 1
}