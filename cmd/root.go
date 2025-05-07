/*
Copyright © 2025 Anderson Macias <anderk222@gmail.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sp",
	Short: "Sp is a command-line code snippet tool",
	Long: `Sp is a interactive commando-line code snippet, you can 
perform various operations such as creating, editing, and deleting snippets.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set datafile using --datafile")
	}

	rootCmd.PersistentFlags().StringVar(&dataFile,
		"datafile", home+string(os.PathSeparator)+".snippets/snippets.json",
		"datafile to store info about snippets")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
