package schedule

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	orderBy string

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all schedules",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List all schedules")
			listSchedules()
		},
	}
)

func listSchedules() []Schedule {
	return []Schedule{
		{
			Name: "Work",
			Date: "2021-02-01",
			Time: "10:00",
		},
		{
			Name: "Sleep",
			Date: "2021-02-01",
			Time: "22:00",
		},
	}
}
