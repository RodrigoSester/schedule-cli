package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Schedule struct {
	Name string
	Date string
	Time string
}

var (
	orderBy string

	listScheduleCmd = &cobra.Command{
		Use:     "list",
		Short:   "List all schedules",
		Aliases: []string{"l"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("List all schedules")
			if _, err := listSchedules(); err != nil {
				fmt.Println("Error:", err)
				return
			}
		},
	}
)

func init() {
	listScheduleCmd.Flags().StringVarP(&orderBy, "order-by", "o", "", "Order by name, date or time")

	viper.BindPFlag("order-by", listScheduleCmd.Flags().Lookup("order-by"))

	rootCmd.AddCommand(listScheduleCmd)
}

func listSchedules() ([]Schedule, error) {
	schedules := []Schedule{
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

	fmt.Println("Schedules: ", schedules)

	return schedules, nil
}
