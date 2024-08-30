package cmd

import "github.com/spf13/cobra"

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