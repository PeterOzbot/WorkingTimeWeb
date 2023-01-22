package excel

import "time"

// request represents a single day with working hours for specific group
type Request struct {
	Date    time.Time `json:"date"`
	Hours   int       `json:"hours"`
	GroupId string    `json:"groupId"`
}

type RequestList struct {
	Days    []*Request `json:"days"`
	GroupId string     `json:"groupId"`
}
