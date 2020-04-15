package core

import "time"

// WorkingDay represents a single days in a month with working data.
type WorkingDay struct {
	Date      time.Time
	IsWorking bool
	Hours     int
}
