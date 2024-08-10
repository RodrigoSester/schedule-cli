package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"encoding/json"
	"github.com/RodrigoSester/schedule-cli/configs"
	"github.com/RodrigoSester/schedule-cli/types"
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

	if name == "" {
		return "", errors.New("name is required")
	}

	

	_, err := os.OpenFile(configs.FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(configs.FilePath)

	if err != nil {
		return "", err
	}

	schedules := &types.ScheduleList{}

	if err := json.Unmarshal(data, &schedules); err != nil {
		return "", err
	}

	os.Remove(configs.FilePath)

	file, err := os.Create(configs.FilePath)

	if err != nil {
		return "", err
	}

	newSchedule := types.Schedule{
		Name: name,
		Date: date,
	}

	schedules.Schedules.Tasks = append(schedules.Schedules.Tasks, newSchedule)

	jsonTaskContent, err := json.Marshal(&schedules)

	if err != nil {
		return "", err
	}

	io.WriteString(file, string(jsonTaskContent))

	fmt.Printf("Schedule %s created on %s\n", name, date)

	return newSchedule.ToString(), nil
}
