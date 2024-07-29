package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createScheduleCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new schedule",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := createSchedule(); err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
	Aliases: []string{"c"},
}

func init() {
	createScheduleCmd.Flags().StringP("name", "n", "", "Name of the schedule")
	createScheduleCmd.Flags().StringP("date", "d", "", "Date of the schedule")
	createScheduleCmd.Flags().StringP("time", "t", "", "Time of the schedule")

	viper.BindPFlag("name", createScheduleCmd.Flags().Lookup("name"))
	viper.BindPFlag("date", createScheduleCmd.Flags().Lookup("date"))
	viper.BindPFlag("time", createScheduleCmd.Flags().Lookup("time"))

	rootCmd.AddCommand(createScheduleCmd)
}

func createSchedule() (string, error) {
	name := viper.GetString("name")
	date := viper.GetString("date")
	time := viper.GetString("time")

	if name == "" {
		return "", errors.New("name is required")
	}

	if date == "" {
		return "", errors.New("date is required")
	}

	if time == "" {
		return "", errors.New("time is required")
	}

	fmt.Printf("Schedule %s created on %s at %s\n", name, date, time)

	return name, nil
}
