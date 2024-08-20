/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiates the tool to sync dot and configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		printIntro()
	},
}

func printIntro() {
	fmt.Println(`>>>>> Welcome to Homenv!\n
							This tool will help you synchronize your configuration files (dotfiles, scripts, etc)`)
	// Init an empty git repo
	// Show next steps
}
func init() {
	rootCmd.AddCommand(initCmd)
}
