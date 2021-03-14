package generator

import (
	"math/rand"
	"time"
)

// Fills working days. Result is array with true if day should have working hours.
func fillDays(hours []int, year int, month time.Month, skipDayChance [7]int) []*WorkingDay {

	// get available days
	var availableDays int = time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// initialize array of working days
	var days []*WorkingDay = make([]*WorkingDay, availableDays)

	// if there are more "instances" of hours than days then we must combine and make as much hours as days
	countDifference := len(hours) - availableDays
	hours = reduce(hours, countDifference)

	// initialize workign date in case the input is not first of the month
	workingDate := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	// used to go through hours array
	var hoursIndex int = 0

	// fill working days array with days and data about it
	// skip sunday/saturday depending on chance
	for dayIndex := range days {

		currentHours := 0

		// figure it out if its working day only if there are enough hours
		// and day is not skiped
		if len(hours)-1 >= hoursIndex && !skipDay(workingDate, skipDayChance) {
			currentHours = hours[hoursIndex]
			hoursIndex++
		}

		// create the working day struct
		days[dayIndex] = &WorkingDay{
			Date:  workingDate,
			Hours: currentHours,
		}

		// update the date with next day
		workingDate = workingDate.AddDate(0, 0, 1)
	}

	// check if all hours have been used
	// if not go through days again and fill non working days
	if len(hours)-1 >= hoursIndex {

		for dayIndex := range days {
			if days[dayIndex].Hours == 0 && len(hours)-1 >= hoursIndex {
				days[dayIndex].Hours = hours[hoursIndex]
				hoursIndex++
			}
		}
	}

	return days
}

// Should day be skiped when setting hours.
func skipDay(date time.Time, skipChance [7]int) bool {

	// get day of the week
	dayOfTheWeek := int(date.Weekday())

	// if chance is zero always skip
	if skipChance[dayOfTheWeek] == 0 {
		return false
	}

	// if chance is 1 always
	if skipChance[dayOfTheWeek] == 1 {
		return true
	}

	if skipChance[dayOfTheWeek] < 0 {
		// get randomized result
		randResult := rand.Intn(skipChance[dayOfTheWeek] * -1)

		// run rand and match
		return randResult == 0
	} else {
		// get randomized result
		randResult := rand.Intn(skipChance[dayOfTheWeek])

		// run rand and match
		return randResult > 0
	}

}
