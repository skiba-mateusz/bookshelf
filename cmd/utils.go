package cmd

import (
	"fmt"
	"strings"

	"github.com/skiba-mateusz/bookshelf/store"
	"github.com/spf13/cobra"
)

// Displays collection of books
func printBooks(cmd *cobra.Command, books []store.Book) {
	if len(books) == 0{
		cmd.Println("Book collection is empty")
		return
	}
	
	cmd.Println("Your book collection:")
	cmd.Printf("%-4s %-30s %-30s %-6s %-6s\n", "#", "Title", "Author", "Year", "Read")
	cmd.Println(strings.Repeat("-", 78))
	for _, book := range books {
		isRead := "[✘]"
		if *book.Read {
			isRead = "[✔]"
		}
		cmd.Printf("%-4d %-30s %-30s %-6d %-6s\n", book.ID, book.Title, book.Author, book.Year, isRead)
	}
}

// Parses read satus flag
func parseReadFlag(cmd *cobra.Command) (*bool, error) {
	readStr, err := cmd.Flags().GetString("read")
	if err != nil {
		cmd.PrintErrln("Error retrieving read flag:", err)
		return nil, err
	}

	var read *bool
	if readStr == "yes" {
		value := true
		read = &value
	} else if readStr == "no" {
		value := false
		read = &value
	}  else if readStr != "" {
		return nil, fmt.Errorf("invalid value for --read flag: must be 'yes' or 'no'")
	}
	
	return read, nil
}

// Retrievs flag as string
func getStringFlag(cmd *cobra.Command, name string) (string, error) {
	value, err := cmd.Flags().GetString(name)
	if err != nil {
		cmd.PrintErrln("Error retrieving", name, "flag:", err)
		return "", err
	}
	return value, nil
} 

// Retrievs flag as int
func getIntFlag(cmd *cobra.Command, name string) (int, error) {
	value, err := cmd.Flags().GetInt(name)
	if err != nil {
		cmd.PrintErrln("Error retrieving", name, "flag:", err)
		return 0, err
	}

	return value, nil
}