package cmd

import (
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

			if title == "" || author == "" || year <= 0 {
				cmd.PrintErrln("Error: all flags --title, --author, --year are required")
				cmd.Usage()
				return nil
			}

			err = bookStore.Add(title, author, year)
			if err != nil {
				return err
			}

			cmd.PrintErrln("Book added successfully")
			return nil
		},
	}

	command.Flags().StringP("title", "t", "", "Adds book title")
	command.Flags().StringP("author", "a", "", "Adds book author")
	command.Flags().IntP("year", "y", 0, "Adds book publication year")

	return command
}