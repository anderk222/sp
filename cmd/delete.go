/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/anderk222/sp/snippet"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete [identifier]",
	Aliases: []string{"del"},
	Short:   "Delete a snippet by its identifier",
	Long:    `This command allows you to delete a snippet using its unique identifier.`,
	Run:     runDelete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runDelete(cmd *cobra.Command, args []string) {

	if len(args) < 1 {

		cmd.Println("Error: Identifier is required.")
		return
	}

	err := snippet.DeleteSnippet(dataFile, args[0])

	if err != nil {
		cmd.Println("Error deleting snippet:", err)
		return
	}

	cmd.Println("Snippet deleted successfully.")
}
