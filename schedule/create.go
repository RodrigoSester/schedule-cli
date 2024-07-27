package schedule

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	createScheduleCmd.Flags().String("name", name, "Name of the schedule")
	viper.BindPFlag("name", createScheduleCmd.Flags().Lookup("name"))

	rootCmd.AddCommand(createScheduleCmd)
}

var (
	name string
	date string
	time string

	createScheduleCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new schedule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := createSchedule(); err == "" {
				fmt.Println("Name is required")
			}

			fmt.Printf("Schedule %s created\n", name)
		},
		Aliases: []string{"c", "-c"},
	}
)

func createSchedule() string {
	name = viper.GetString("name")

	if name == "" {
		return errors.New("name is required").Error()
	}

	return name
}
