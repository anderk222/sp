/*
Copyright © 2025 Anderson Macias <anderk222@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/anderk222/sp/snippet"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List info of snippets",
	Long:  `It will show te info of de code snippets and their details.`,
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listRun(cmd *cobra.Command, args []string) {

	snippets, err := snippet.ReadItems(dataFile)
	if err != nil {
		fmt.Println("Error reading snippets:", err)
		return
	}

	for key, val := range snippets {
		fmt.Printf("\033[33mSnippet %s\033[0m \n\n Description: %s \n Location: %s\n\n", key, val.Description, val.FileContentName)
	}
}
