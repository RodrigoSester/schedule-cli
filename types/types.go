package types

type Schedule struct {
	Name      string `json:"name"`
	Date      string `json:"date,omitempty"`
	Time      string `json:"time,omitempty"`
	Type      string `json:"type"`
	Frequency string `json:"frequency"`
}

func (s *Schedule) ToString() string {
	return "Schedule"
}

type ScheduleList struct {
	Schedules struct {
		Tasks []Schedule `json:"tasks"`
	}
}
