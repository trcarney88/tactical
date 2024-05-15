/*
Copyright © 2024 Todd Carney <toddcarney44@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tactical",
	Short: "Todd's Awesome Cli for Triage, Incident Correction, and AnaLytics",
	Long: `
	This is a CLI I'm building to just be my little handy dandy CLI to get news,
	create issues in Linear for work and personal projects, run scripts for things
	at work and personal projects, and a ChatGPT interface.

	The goal is to use the CLI to get all the information I want from the internet
	without going on a browser and to learn go, cobra, and bubble tea.`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tactical.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
