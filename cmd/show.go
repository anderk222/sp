/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/anderk222/sp/snippet"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show [identifier]",
	Short: "Show code content of a snippet",
	Long:  `It will show the code snippet of the code snippet passed`,
	Run:   runShow,
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runShow(cmd *cobra.Command, args []string) {

	if len(args) < 1 {
		fmt.Println("Please provide a snippet name.")
		return
	}

	item, err := snippet.RetrieveSnippet(dataFile, args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	snippetContent, err := snippet.RetrieveSnippetContent(item)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(snippetContent)

}
