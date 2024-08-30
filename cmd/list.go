package cmd

import (
	"strings"

	"github.com/skiba-mateusz/bookshelf/store"
	"github.com/spf13/cobra"
)

func ListBooksCommand(bookStore *store.BookStore) *cobra.Command {
	command := &cobra.Command{
		Use: "list",
		Short: "Displays collection of your books",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(bookStore.Books()) == 0{
				cmd.Println("Book collection is empty")
				return nil
			}
			
			cmd.Println("Your book collection:")
			cmd.Printf("%-4s %-30s %-30s %-6s %-6s\n", "#", "Title", "Author", "Year", "Read")
			cmd.Println(strings.Repeat("-", 78))
			for _, book := range bookStore.Books() {
				isRead := "[✘]"
				if book.Read {
					isRead = "[✔]"
				}
				cmd.Printf("%-4d %-30s %-30s %-6d %-6s\n", book.ID, book.Title, book.Author, book.Year, isRead)
			}
			return nil
		},
	}

	return command
}