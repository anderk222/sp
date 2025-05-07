/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/anderk222/sp/snippet"
	"github.com/spf13/cobra"
)

// injectCmd represents the inject command
var injectCmd = &cobra.Command{
	Use:     "inject [identifier] [file]",
	Aliases: []string{"ij"},
	Short:   "Inject code snippet to a file",
	Long:    `This commando will add a code snippet to the specified file, if this file does not exist, it will create a new one.`,
	Run:     runInject,
}

func init() {
	rootCmd.AddCommand(injectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// injectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// injectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runInject(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide a snippet name and a file name.")
		return
	}

	snippetName := args[0]
	fileName := args[1]

	err := snippet.InjectSnippet(dataFile, fileName, snippetName)
	if err != nil {
		fmt.Println("Error injecting snippet:", err)
		return
	}

	fmt.Println("Snippet injected successfully")
}
