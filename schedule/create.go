package schedule

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createScheduleCmd)
}

var createScheduleCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new schedule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create a new schedule")
	},
}
