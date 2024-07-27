package schedule

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "schedule",
	Short: "schedule is a CLI tool to manage your schedule",
	Long:  `schedule is a CLI tool to manage your schedule. You can add, list, update and delete your schedule`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Schedule CLI")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
