package generator

import (
	"math/rand"
	"time"
)

// Generate : Generates random sequence of working hours.
func Generate(request *Request) []*WorkingDay {

	// reinitialize the seed
	rand.Seed(time.Now().UnixNano())

	// index 0 is for 1h, while index 9 is for 10h
	choiceWeight := [10]int{0, 8, 8, 6, 4, 4, 4, 4, 1, 1}

	// get hours
	totalHours := request.A_Hours + request.B_Hours
	hours := generateHours(totalHours, choiceWeight)

	// skip day chance
	skipDay := [7]int{6, -8, -8, -8, -8, -8, 4}

	// get days
	days := fillDays(hours, request.Year, time.Month(request.Month), skipDay)

	// split into groups
	return splitDaysIntoGroups(days, request.A_Hours)
}
