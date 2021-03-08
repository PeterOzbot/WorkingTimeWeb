package generator

import "time"

// WorkingDay represents a single days in a month with working data.
type WorkingDay struct {
	Date      time.Time `json:"date"`
	IsWorking bool      `json:"isWorking"`
	Hours     int       `json:"hours"`
}
