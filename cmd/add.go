package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/skiba-mateusz/bookshelf/store"
	"github.com/spf13/cobra"
)

// Command for adding books to collection
func AddBookCommand(bookStore *store.BookStore) *cobra.Command {
	command := &cobra.Command{
		Use: "add",
		Short: "Adds new book to your collection",
		RunE: func(cmd *cobra.Command, args []string) error {
			title, err := getStringFlag(cmd, "title")
			if err != nil {
				return err
			}
			author, err := getStringFlag(cmd, "author")
			if err != nil {
				return err
			}
			year, err := getIntFlag(cmd, "year")
			if err != nil {
				return err
			}
			readStr, err := getStringFlag(cmd, "read")
			if err != nil {
				return err
			}

			var read bool
			if strings.ToLower(readStr) == "yes" {
				read = true
			} else if strings.ToLower(readStr) == "no" {
				read = false
			} else if readStr != "" {
				return fmt.Errorf("invalid value for --read flag: it must be 'yes' or 'no'")
			}

			if title == "" || author == "" || year <= 0 {
				cmd.PrintErrln("Error: flags --title, --author, --year are required")
				cmd.Usage()
				return nil
			}

			book := store.Book{
				Title: title,
				Author: author,
				Year: year,
				Read: read,
				CreatedAt: time.Now(),
			}

			err = bookStore.Add(book)
			if err != nil {
				return err
			}

			cmd.PrintErrln("Book added successfully")
			return nil
		},
	}

	command.Flags().StringP("title", "t", "", "Adds book title")
	command.Flags().StringP("author", "a", "", "Adds book author")
	command.Flags().StringP("read", "r", "no", "Adds book read status ('yes' or 'no')")
	command.Flags().IntP("year", "y", 0, "Adds book publication year")

	return command
}