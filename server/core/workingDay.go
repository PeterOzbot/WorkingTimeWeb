package core

import "time"

// WorkingDay represents a single days in a month with working data.
type WorkingDay struct {
	Date      time.Time
	IsWorking bool
	Hours     int

	// TODO:PETERO review this
	UserID string `json:"user_id" bson:"userId"`
	RepoID string `json:"id" bson:"repoId"`
}
