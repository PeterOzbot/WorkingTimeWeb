package generator

import (
	"math/rand"
	"workingtimeweb/server/core"
)

// Generates list for days with random hours.
func generateHours(totalHours int) []int {

	// index 0 is for 1h, while index 9 is for 10h
	choiceWeight := [10]int{0, 1, 2, 4, 5, 5, 6, 8, 3, 1}
	sumWeight := 35
	numChoices := 10

	// randomize hours until all are used
	var hours []int
	for totalHours > 0 {
		var workingHours int = randomize(choiceWeight[:], sumWeight, numChoices)
		// index 0 is for 1h, while index 9 is for 10h so it should be incremented by 1
		workingHours++

		workingHours = core.Min(totalHours, workingHours)
		totalHours -= workingHours
		hours = append(hours, workingHours)
	}

	return hours
}

// Calculates weighted random from list of weights. Result is selected index. Caller should handle which weight index is the actual selection value.
func randomize(choiceWeight []int, sumWeight int, numChoices int) int {

	// get random number from sum of all weights
	var rnd int = rand.Intn(sumWeight)

	// search for randomized index
	// https://stackoverflow.com/questions/1761626/weighted-random-numbers
	var index int
	for index = 0; index < numChoices; index++ {
		if rnd < choiceWeight[index] {
			break
		}

		rnd -= choiceWeight[index]
	}

	return index
}

// Reduces number of hours instances by countDifference number
func reduce(inputHours []int, countDifference int) []int {
	// if there are no differences just return input
	// negative differences means that theere are less hours as required, so also no reducing need
	if countDifference <= 0 {
		return inputHours
	}

	// clone the hours list to prevent changin the original
	hours := make([]int, len(inputHours))
	copy(hours[:], inputHours[:])

	for countDifference > 0 {
		// find two lowest hours and combine
		var minIndex1 int = 0
		var minIndex2 int = 0

		for i, e := range hours {
			if i == 0 || e < hours[minIndex1] {
				minIndex1 = i
			}

			if (minIndex2 != minIndex1 && minIndex1 != i) && e < hours[minIndex2] {
				minIndex2 = i
			}
		}

		// decrease the number of hours we must reduce
		countDifference--

		// combine the two minimal hours
		hours[minIndex2] = hours[minIndex1] + hours[minIndex2]

		// remove one of the combine hours
		copy(hours[minIndex1:], hours[minIndex1+1:])
		hours = hours[:len(hours)-1]
	}

	return hours
}
