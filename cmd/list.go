package cmd

import (
	"github.com/skiba-mateusz/bookshelf/store"
	"github.com/spf13/cobra"
)

func ListBooksCommand(bookStore *store.BookStore) *cobra.Command {
	command := &cobra.Command{
		Use: "list",
		Short: "Displays collection of your books",
		RunE: func(cmd *cobra.Command, args []string) error {
			printBooks(cmd, bookStore.Books())
			return nil
		},
	}

	return command
}