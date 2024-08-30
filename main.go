package main

import (
	"log"

	"github.com/skiba-mateusz/bookshelf/cmd"
	"github.com/skiba-mateusz/bookshelf/paths"
	"github.com/spf13/cobra"
)

func main() {
	_, err := paths.GetBooksJsonFile()
	if err != nil {
		log.Fatalf("Error retrieving json file: %v", err)
	}

	rootCmd := &cobra.Command{
		Use: "bookshelf",
		Short: "CLI app for managing your book collection",
	}

	rootCmd.AddCommand(cmd.AddBookCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
