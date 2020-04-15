package generator

import (
	"math/rand"
	"time"
	"workingtimeweb/server/core"
)

// Generate : Generates random sequence of working hours.
func Generate(totalHours int) []*core.WorkingDay {

	// reinitialize the seed
	rand.Seed(time.Now().UnixNano())

	// get hours
	var hours []int = GenerateHours(totalHours)

	// get days
	var days []*core.WorkingDay = RandomizeDays(hours, time.Now())

	// create list of working days with appropriate working hours
	return fillDays(days, hours)
}

// fillDays : creates list of working days from hours and days
func fillDays(workingDays []*core.WorkingDay, hours []int) []*core.WorkingDay {

	// if there are more "instances" of hours than days then we must combine and make as much hours as days
	countDifference := len(hours) - len(workingDays)
	hours = Reduce(hours, countDifference)

	// fill with data
	for i, e := range hours {
		workingDays[i] = &core.WorkingDay{
			Date:      time.Now(),
			Hours:     e,
			IsWorking: true,
		}
	}

	return workingDays
}
