package generator

import (
	"math/rand"
	"time"
)

// Generate : Generates random sequence of working hours.
func Generate(request *Request) []*WorkingDay {

	// reinitialize the seed
	rand.Seed(time.Now().UnixNano())

	// get hours
	var hours []int = generateHours(request.TotalHours)

	// get days
	return fillDays(hours, time.Now())
}
