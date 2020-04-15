package generator

import (
	"time"
	"workingtimeweb/server/core"
)

// RandomizeDays : Randomizes working days. Result is array with true if day should have working hours.
func RandomizeDays(hours []int, desiredDate time.Time) []*core.WorkingDay {

	// get available days
	var availableDays int = time.Date(desiredDate.Year(), desiredDate.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// initialize array of working days
	var days []*core.WorkingDay = make([]*core.WorkingDay, availableDays)

	// initialize workign date in case the input is not first of the month
	workingDate := time.Date(desiredDate.Year(), desiredDate.Month(), 1, 0, 0, 0, 0, time.UTC)

	// fill working days array with days and if its working days
	// for now we just fill first x days if there are hours for them
	for dayIndex := range days {

		// figure it out if its working day
		var isWorking bool = false
		if len(hours)-1 >= dayIndex {
			isWorking = true
		}

		// create the working day struct
		days[dayIndex] = &core.WorkingDay{
			Date:      workingDate,
			IsWorking: isWorking,
		}

		// update the date with next day
		workingDate = workingDate.AddDate(0, 0, 1)
	}

	return days
}
