package types

type Schedule struct {
	Name string `yaml:"name"`
	Date string `yaml:"date"`
	Time string `yaml:"time"`
}

type ScheduleList struct {
	Schedules struct {
		Tasks []Schedule `yaml:"tasks"`
	}
}
