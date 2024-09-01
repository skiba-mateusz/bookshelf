package cmd

import (
	"strings"

	"github.com/skiba-mateusz/bookshelf/store"
	"github.com/spf13/cobra"
)

// Command for searching books
func SearchBooks(bookStore *store.BookStore) *cobra.Command {
	command := &cobra.Command{
		Use: "search",
		Short: "Searches for books by title, author or year",
		RunE: func(cmd *cobra.Command, args []string) error {
			title, err := getStringFlag(cmd, "title")
			if err != nil {
				return err
			}
			author, err := getStringFlag(cmd, "author")
			if err != nil {
				return err
			}
			read, err := parseReadFlag(cmd) 
			if err != nil {
				return err
			}
			year, err := getIntFlag(cmd, "year")
			if err != nil {
				return err
			}

			if title == "" && author == "" && year == 0 && read == nil {
				cmd.PrintErrln("Error: at least one flag --title, --author or --year must be specified")
				cmd.Usage()
				return nil
			}

			query := store.SearchQuery{
				Title: strings.ToLower(title),
				Author: strings.ToLower(author),
				Year: year,
				Read: read,
			}

			results := bookStore.Search(query)

			printBooks(cmd, results)

			return nil
		},
	}

	command.Flags().StringP("title", "t", "", "Adds book title for querying")
	command.Flags().StringP("author", "a", "", "Adds book author for querying")
	command.Flags().StringP("read", "r", "", "Adds book read status ('yes' or 'no') for querying")
	command.Flags().IntP("year", "y", 0, "Adds book year for querying")

	return command
}