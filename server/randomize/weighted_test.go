package randomize

import (
	"testing"
)

// Test_Randomize_Ordered : Tests ordered weights. First(the smallest) should be selected the least amount of times, while last(the largest) should be selected the most.
func Test_Randomize_Ordered(t *testing.T) {

	choiceWeight := [3]int{1, 2, 3}

	result := [3]int{0, 0, 0}
	for runIndex := 0; runIndex < 10000; runIndex++ {
		currentResult := Weighted(choiceWeight[:])
		result[currentResult]++
	}

	if result[0] > result[1] || result[0] > result[2] {
		t.Errorf("result[0] is not selected the least amount of times")
	}

	if result[1] > result[2] {
		t.Errorf("result[1] is not selected less times than result[2]")
	}
}

// Test_Randomize_Unordered : Tests unordered weights. Second(the smallest) should be selected the least amount of times, while first(the largest) should be selected the most.
func Test_Randomize_Unordered(t *testing.T) {

	choiceWeight := [3]int{3, 1, 2}

	result := [3]int{0, 0, 0}
	for runIndex := 0; runIndex < 10000; runIndex++ {
		currentResult := Weighted(choiceWeight[:])
		result[currentResult]++
	}

	if result[1] > result[0] || result[1] > result[2] {
		t.Errorf("result[1] is not selected the least amount of times")
	}

	if result[0] < result[1] || result[0] < result[2] {
		t.Errorf("result[0] is not selected the most amount of times")
	}
}

// Test_Randomize_Hours : Test used for testing hours distribution randomization.
func Test_Randomize_Hours(t *testing.T) {

	// index 0 is for 1h, while index 9 is for 10h
	choiceWeight := [10]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}

	result := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for runIndex := 0; runIndex < 10000; runIndex++ {
		currentResult := Weighted(choiceWeight[:])
		result[currentResult]++
	}

	// first and last selected less times than the rest
	if result[0] > result[1] || result[0] > result[2] || result[0] > result[3] || result[0] > result[4] || result[0] > result[5] || result[0] > result[6] || result[0] > result[7] || result[0] > result[8] {
		t.Errorf("result[0] is not selected the least amount of times")
	}
	if result[9] > result[1] || result[9] > result[2] || result[9] > result[3] || result[9] > result[4] || result[9] > result[5] || result[9] > result[6] || result[9] > result[7] || result[9] > result[8] {
		t.Errorf("result[9] is not selected the least amount of times")
	}

	// second and before last must be selected less times than the inner ones
	if result[1] > result[2] || result[1] > result[3] || result[1] > result[4] || result[1] > result[5] || result[1] > result[6] || result[1] > result[7] {
		t.Errorf("result[1] is not selected less amount of times then inner ones.")
	}
	if result[8] > result[2] || result[8] > result[3] || result[8] > result[4] || result[8] > result[5] || result[8] > result[6] || result[8] > result[7] {
		t.Errorf("result[8] is not selected less amount of times then inner ones.")
	}

	// third and third before last less than inner ones and so on...
	if result[2] > result[3] || result[2] > result[4] || result[2] > result[5] || result[2] > result[6] {
		t.Errorf("result[2] is not selected less amount of times then inner ones.")
	}
	if result[7] > result[3] || result[7] > result[4] || result[7] > result[5] || result[7] > result[6] {
		t.Errorf("result[7] is not selected less amount of times then inner ones.")
	}

	if result[3] > result[4] || result[3] > result[5] {
		t.Errorf("result[3] is not selected less amount of times then inner ones.")
	}
	if result[6] > result[4] || result[6] > result[5] {
		t.Errorf("result[6] is not selected less amount of times then inner ones.")
	}
}
