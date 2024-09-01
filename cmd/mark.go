package cmd

import (
	"github.com/skiba-mateusz/bookshelf/store"
	"github.com/spf13/cobra"
)

// Command for marking book as read/unread
func MarkReadCommand(bookStore *store.BookStore) *cobra.Command {
	command := &cobra.Command{
		Use: "mark",
		Short: "Marks the book as read/unread",
		RunE: func(cmd *cobra.Command, args []string) error {
			bookID, err := getIntFlag(cmd, "id")
			if err != nil {
				return err
			}
			read, err := parseReadFlag(cmd)
			if err != nil {
				return err
			}

			if bookID <= 0 {
				cmd.PrintErrln("Error: id must be a positive integer")
				cmd.Usage()
				return nil
			}

			err = bookStore.Mark(int64(bookID), read)
			if err != nil {
				return err
			}

			cmd.Println("Book marked successfully")
			return nil
		},
	}

	command.Flags().Int("id", 0, "Adds ID of book to mark")
	command.Flags().String("read", "", "Adds read status for book ('yes' or 'no')")

	return command
}