package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type ScheduleList struct {
	Schedules struct {
		Tasks []Schedule `yaml:"tasks"`
	}
}

type Schedule struct {
	Name string `yaml:"name"`
	Date string `yaml:"date"`
	Time string `yaml:"time"`
}

var (
	fOrderBy string
	fGroupBy string
	fGroupByValue string

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
	listScheduleCmd.Flags().StringVarP(&fOrderBy, "order-by", "o", "", "Order by name, date or time")
	listScheduleCmd.Flags().StringVarP(&fGroupBy, "group-by", "g", "", "Group by name or date")
	listScheduleCmd.Flags().StringVarP(&fGroupByValue, "group-by-value", "gv", "", "Group by value")

	viper.BindPFlag("order-by", listScheduleCmd.Flags().Lookup("order-by"))

	rootCmd.AddCommand(listScheduleCmd)
}

func listSchedules() (ScheduleList, error) {
	data, err := os.ReadFile("./schedules.yml")

	if err != nil {
		return ScheduleList{}, err
	}

	var schedules ScheduleList

	if err := yaml.Unmarshal(data, &schedules); err != nil {
		return ScheduleList{}, err
	}

	if (fGroupBy != "") {
		fmt.Println("Group by: ", fGroupBy)

		var schedulesFiltered ScheduleList

		for idx, schedule := range schedules.Schedules.Tasks {
			if (fGroupBy == "date" && schedule.Date == fGroupByValue) {
				fmt.Printf("Schedule index %v: %v\n", idx, schedule)
				schedulesFiltered.Schedules.Tasks = append(schedulesFiltered.Schedules.Tasks, schedule)
			}
		}

		fmt.Println("Schedules Filtered: ", schedulesFiltered)

		return schedulesFiltered, nil
	}

	fmt.Println("Schedules: ", schedules.Schedules.Tasks)

	return schedules, nil
}
