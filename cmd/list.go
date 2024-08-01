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

func listSchedules() (ScheduleList, error) {
	data, err := os.ReadFile("./schedules.yml")

	if err != nil {
		return ScheduleList{}, err
	}

	var schedules ScheduleList

	if err := yaml.Unmarshal(data, &schedules); err != nil {
		return ScheduleList{}, err
	}

	fmt.Println("Schedules: ", schedules.Schedules.Tasks)

	return schedules, nil
}
