package main

import (
	"log"

	"github.com/skiba-mateusz/bookshelf/paths"
)

func main() {
	_, err := paths.GetBooksJsonFile()
	if err != nil {
		log.Fatalf("Error retrieving json file: %v", err)
	}
}