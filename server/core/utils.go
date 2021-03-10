package core

// Min : Returns the smaller of a or b.
func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Sum : Returns sum of all items in array.
func Sum(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}
