/*
Copyright © 2024 Rodrigo Weber Sesterheim <rodrigowsesterheim@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "schedule-cli",
	Short: "A CLI to control your schedule",
	Long:  `schedule-cli is a CLI tool to manage your schedule. You can add, list, update and delete your schedule`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Schedule CLI")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
