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
	var days []*WorkingDay = randomizeDays(hours, time.Now())

	// create list of working days with appropriate working hours
	return fillDays(days, hours)
}

// fillDays : creates list of working days from hours and days
func fillDays(workingDays []*WorkingDay, hours []int) []*WorkingDay {

	// if there are more "instances" of hours than days then we must combine and make as much hours as days
	countDifference := len(hours) - len(workingDays)
	hours = reduce(hours, countDifference)

	// fill with data
	for i, e := range hours {
		workingDays[i] = &WorkingDay{
			Date:      time.Now(),
			Hours:     e,
			IsWorking: true,
		}
	}

	return workingDays
}
