package cmd

import (
	"github.com/skiba-mateusz/bookshelf/store"
	"github.com/spf13/cobra"
)

func DeleteBookCommand(bookStore *store.BookStore) *cobra.Command {
	command := &cobra.Command{
		Use: "delete",
		Short: "Deletes book from your collection",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := getIntFlag(cmd, "id")
			if err != nil {
				return err
			}

			if id <= 0 {
				cmd.PrintErrln("Error: invalid ID Value: ID must be a positive integer")
				cmd.Usage()
				return nil
			}
			err = bookStore.Delete(int64(id))
			if err != nil {
				return err
			}

			cmd.Println("Book deleted successfully")
			return nil
		},
	}

	command.Flags().Int("id", 0, "Adds ID of the book to delete")

	return command
}