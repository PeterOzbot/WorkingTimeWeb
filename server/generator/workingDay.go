package generator

import "time"

// WorkingDay represents a single days in a month with working data.
type WorkingDay struct {
	Date    time.Time `json:"date"`
	Hours   int       `json:"hours"`
	GroupId string    `json:"groupId"`
}
