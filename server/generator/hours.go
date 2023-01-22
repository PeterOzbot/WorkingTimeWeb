package generator

import (
	"workingtimeweb/server/core"
	"workingtimeweb/server/randomize"
)

// Generates list for days with random hours.
func generateHours(totalHours int, choiceWeight [10]int) []int {

	// randomize hours until all are used
	var hours []int

	// generate until total hours have been used
	for totalHours > 0 {

		// get random hours
		var workingHours int = randomize.Weighted(choiceWeight[:])
		// index 0 is for 1h, while index 9 is for 10h so it should be incremented by 1
		workingHours++

		// handle if we reach the total hours
		workingHours = core.Min(totalHours, workingHours)
		totalHours -= workingHours

		// add to list
		hours = append(hours, workingHours)
	}

	return hours
}

// Reduces number of hours instances by countDifference number
func reduce(inputHours []int, countDifference int) []int {
	// if there are no differences just return input
	// negative differences means that there are less hours as required, so also no reducing need
	if countDifference <= 0 {
		return inputHours
	}

	// clone the hours list to prevent changing the original
	hours := make([]int, len(inputHours))
	copy(hours[:], inputHours[:])

	for countDifference > 0 {

		// find two lowest hours and combine

		var firstIndex int = -1
		var secondIndex int = -1

		for index := 0; index < len(hours); index++ {

			if firstIndex == -1 {
				firstIndex = index
			} else if firstIndex != -1 && secondIndex == -1 {
				secondIndex = index
			} else if hours[index] < hours[firstIndex] {
				firstIndex = index
			} else if secondIndex == -1 || hours[index] < hours[secondIndex] {
				secondIndex = index
			}
		}

		// decrease the number of hours we must reduce
		countDifference--

		// combine the two minimal hours
		hours[secondIndex] = hours[firstIndex] + hours[secondIndex]

		// remove one of the combine hours
		copy(hours[firstIndex:], hours[firstIndex+1:])
		hours = hours[:len(hours)-1]
	}

	return hours
}
