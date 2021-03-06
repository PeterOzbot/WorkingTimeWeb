package generator

import (
	"math/rand"
	"time"
)

// Fills working days. Result is array with true if day should have working hours.
func fillDays(hours []int, desiredDate time.Time) []*WorkingDay {

	// get available days
	var availableDays int = time.Date(desiredDate.Year(), desiredDate.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// initialize array of working days
	var days []*WorkingDay = make([]*WorkingDay, availableDays)

	// if there are more "instances" of hours than days then we must combine and make as much hours as days
	countDifference := len(hours) - availableDays
	hours = reduce(hours, countDifference)

	// initialize workign date in case the input is not first of the month
	workingDate := time.Date(desiredDate.Year(), desiredDate.Month(), 1, 0, 0, 0, 0, time.UTC)

	// used to go through hours array
	var hoursIndex int = 0

	// fill working days array with days and data about it
	// skip sunday/saturday depending on chance
	for dayIndex := range days {

		isWorking := false
		currentHours := 0

		// figure it out if its working day only if there are enough hours
		// and day is not skipped
		if len(hours)-1 >= hoursIndex && !skipDay(workingDate) {
			isWorking = true
			currentHours = hours[hoursIndex]
			hoursIndex++
		}

		// create the working day struct
		days[dayIndex] = &WorkingDay{
			Date:      workingDate,
			IsWorking: isWorking,
			Hours:     currentHours,
		}

		// update the date with next day
		workingDate = workingDate.AddDate(0, 0, 1)
	}

	// check if all hours have been used
	// if not go through days again and fill non working days
	if len(hours)-1 >= hoursIndex {

		for dayIndex := range days {
			if !days[dayIndex].IsWorking && len(hours)-1 >= hoursIndex {
				days[dayIndex].IsWorking = true
				days[dayIndex].Hours = hours[hoursIndex]
				hoursIndex++
			}
		}
	}

	return days
}

// Should day be skipped when setting hours.
func skipDay(date time.Time) bool {

	// get day of the week
	dayOfTheWeek := int(date.Weekday())

	if dayOfTheWeek == int(time.Saturday) {
		return rand.Intn(4) != 0
	}

	if dayOfTheWeek == int(time.Sunday) {
		return rand.Intn(8) != 0
	}

	return rand.Intn(3) == 0
}
