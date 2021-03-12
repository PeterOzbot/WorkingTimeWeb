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
	choiceWeight := [10]int{0, 1, 2, 4, 5, 5, 6, 8, 3, 1}

	// get hours
	var hours []int = generateHours(request.TotalHours, choiceWeight)

	// get days
	return fillDays(hours, time.Now().Year(), time.Now().Month(), [7]int{8, 3, 3, 3, 3, 3, 4})
}
