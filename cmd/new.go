/*
Copyright © 2025 Anderson Macias <anderk222@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/anderk222/sp/snippet"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [identifier] [content]",
	Short: "Create a new snippet",
	Long:  `Create will add a new snippet to your collection of snippets.`,
	Run:   runNew,
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runNew(cmd *cobra.Command, args []string) {

	lenArhgs := len(args)

	if lenArhgs < 2 {
		panic("Args are required example sp new <identifier snippet> <content snippet>")
	}
	item := snippet.Snippet{Name: args[0], Content: args[1]}

	if lenArhgs >= 3 {
		item.Description = args[2]
	}

	if err := snippet.SaveSnippetContent(snippet.GenerateFilenameContentSnippet(dataFile, item), &item); err != nil {

		fmt.Printf("Error saving snippet content: %v\n", err)
		return
	}

	snippets, err := snippet.ReadItems(dataFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	snippets[item.Name] = item

	if err := snippet.SaveSnippets(dataFile, snippets); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Snippet saved successfully!")

}
