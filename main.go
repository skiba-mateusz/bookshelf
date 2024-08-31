package main

import (
	"log"

	"github.com/skiba-mateusz/bookshelf/cmd"
	"github.com/skiba-mateusz/bookshelf/paths"
	"github.com/skiba-mateusz/bookshelf/store"
	"github.com/spf13/cobra"
)

func main() {
	booksFile, err := paths.GetBooksJsonFile()
	if err != nil {
		log.Fatalf("Error retrieving json file: %v", err)
	}
	bookStore, err := store.NewBookStore(booksFile)
	if err != nil {
		log.Fatalf("Error initializing BookStore: %v", err)
	}

	rootCmd := &cobra.Command{
		Use: "bookshelf",
		Short: "CLI app for managing your book collection",
	}

	rootCmd.AddCommand(cmd.AddBookCommand(bookStore))
	rootCmd.AddCommand(cmd.ListBooksCommand(bookStore))
	rootCmd.AddCommand(cmd.SearchBooks(bookStore))
	rootCmd.AddCommand(cmd.DeleteBookCommand(bookStore))

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
