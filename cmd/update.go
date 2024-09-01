package cmd

import (
	"github.com/skiba-mateusz/bookshelf/store"
	"github.com/spf13/cobra"
)

func UpdateBookCommand(bookStore *store.BookStore) *cobra.Command {
	command := &cobra.Command{
		Use: "update",
		Short: "Updates information of a specific book",
		RunE: func(cmd *cobra.Command, args []string) error {
			bookID, err := getIntFlag(cmd, "id")
			if err != nil {
				return err
			}
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
			read, err := parseReadFlag(cmd)
			if err != nil {
				return err
			}

			if bookID <= 0 {
				cmd.PrintErrln("Error: flag --id must be a positive integer")
				cmd.Usage()
				return nil
			}

			book := store.Book{
				ID: int64(bookID),
				Title: title,
				Author: author,
				Year:  year,
				Read: read,
			}

			err = bookStore.Update(book)
			if err != nil {
				return err
			}

			cmd.Println("Book updated successfully")
			return nil
		},
	}

	command.Flags().Int("id", 0, "Adds ID of the book to update")
	command.Flags().String("title", "", "Adds new title of the book")
	command.Flags().String("author", "", "Adds new author of the book")
	command.Flags().String("read", "", "Adds new read status of the book")
	command.Flags().Int("year", 0, "Adds new publication year of the book")

	return command
}