package types

type Schedule struct {
	Name      string `yaml:"name"`
	Date      string `yaml:"date"`
	Time      string `yaml:"time"`
	Type      string `yaml:"type"`
	Frequency string `yaml:"frequency"`
}

type ScheduleList struct {
	Schedules struct {
		Tasks []Schedule `yaml:"tasks"`
	}
}
