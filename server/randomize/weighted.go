package randomize

import (
	"math/rand"
	"workingtimeweb/server/core"
)

// Weighted : Calculates weighted random from list of weights. Result is selected index. Caller should handle which weight index is the actual selection value.
// https://stackoverflow.com/questions/1761626/weighted-random-numbers
func Weighted(choiceWeight []int) int {

	// sum and number of choices
	sumWeight := core.Sum(choiceWeight[:])
	numChoices := len(choiceWeight)

	// get random number from sum of all weights
	var rnd int = rand.Intn(sumWeight)

	// search for randomized index
	var index int
	for index = 0; index < numChoices; index++ {
		if rnd < choiceWeight[index] {
			break
		}

		rnd -= choiceWeight[index]
	}

	return index
}
