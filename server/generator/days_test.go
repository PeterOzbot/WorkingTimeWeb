package generator

import (
	"testing"
	"time"
)

// Test_FillDays_LessHours : Tests simple scenario with fewer hours are there are days. The first days must be working.
func Test_FillDays_LessHours(t *testing.T) {
	// we have hours for 5 separate days
	hours := [5]int{4, 5, 3, 1, 10}
	skipChance := [7]int{1, 1, 1, 1, 1, 1, 1}

	// lets generate which days will be working
	days := fillDays(hours[:], 2020, time.April, skipChance)

	// count working days
	var workingDaysCont int = 0
	for _, workingDay := range days {
		if workingDay.Hours > 0 {
			workingDaysCont++
		}
	}

	// for this case lengths must match
	if workingDaysCont != len(hours) {
		t.Errorf("days is not the same length as hours")
	}
}

// Test_FillDays_MoreHours : Tests scenario with more hours that there are days. All days should be working.
func Test_FillDays_MoreHours(t *testing.T) {
	// we have hours for 40 separate days
	hours := [40]int{4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10}
	skipChance := [7]int{1, 1, 1, 1, 1, 1, 1}

	// lets generate which days will be working
	days := fillDays(hours[:], 2020, time.April, skipChance)

	// check working days
	for _, workingDay := range days {
		if workingDay.Hours == 0 {
			t.Errorf("not all days are working days")
			break
		}
	}
}

// Test_FillDays_AllDays : Tests if all days in a month are returned.
func Test_FillDays_AllDays(t *testing.T) {
	skipChance := [7]int{1, 1, 1, 1, 1, 1, 1}
	hours := [5]int{4, 5, 3, 1, 10}

	// lets generate which days will be working
	days := fillDays(hours[:], 2021, time.March, skipChance)

	// in march there are 31 days
	marchDaysCount := 31
	// check working days
	if len(days) != marchDaysCount {
		t.Errorf("not all days are returned")
	}
}

// Test_Skip : Tests if day is skipped more than not regarding input chance.
func Test_Skip(t *testing.T) {

	// go through all days
	for dayIndex := range [7]int{0, 1, 2, 3, 4, 5, 6} {

		// date to check
		date := time.Date(2021, 3, dayIndex, 0, 0, 0, 0, time.UTC)

		// chance
		skipChance := [7]int{0, 0, 0, 0, 0, 0, 0}
		skipChance[dayIndex] = 3 // chance 1 in 3

		// go through it 10k times to get approximate average
		result := [2]int{0, 0}
		for runIndex := 0; runIndex < 10000; runIndex++ {
			currentResult := skipDay(date, skipChance)
			if currentResult {
				result[0]++
			} else {
				result[1]++
			}
		}

		// validate
		if result[0] < result[1] {
			t.Errorf("day %d should be skiped more times than not", dayIndex)
			break
		}
	}
}

// Test_Skip_Negative : Tests if day is skipped less than not regarding input chance when negative inputs.
func Test_Skip_Negative(t *testing.T) {

	// go through all days
	for dayIndex := range [7]int{0, 1, 2, 3, 4, 5, 6} {

		// date to check
		date := time.Date(2021, 3, dayIndex, 0, 0, 0, 0, time.UTC)

		// chance
		skipChance := [7]int{0, 0, 0, 0, 0, 0, 0}
		skipChance[dayIndex] = -3 // chance 1 in 3

		// go through it 10k times to get approximate average
		result := [2]int{0, 0}
		for runIndex := 0; runIndex < 10000; runIndex++ {
			currentResult := skipDay(date, skipChance)
			if currentResult {
				result[0]++
			} else {
				result[1]++
			}
		}

		// validate
		if result[0] > result[1] {
			t.Errorf("day %d should be skiped less times than not", dayIndex)
			break
		}
	}
}

// Test_Skip_Never : Tests if day is never skiped when chance is 0.
func Test_Skip_Never(t *testing.T) {

	// date to check
	dayIndex := 1
	date := time.Date(2021, 3, dayIndex, 0, 0, 0, 0, time.UTC)

	// chance
	daySkipChance := 0
	skipChance := [7]int{}
	skipChance[dayIndex] = daySkipChance

	// go through it 10k times to get approximate average
	result := [2]int{0, 0}
	for runIndex := 0; runIndex < 10000; runIndex++ {
		currentResult := skipDay(date, skipChance)
		if currentResult {
			result[0]++
		} else {
			result[1]++
		}
	}

	// validate
	if result[0] != 0 || result[1] != 10000 {
		t.Errorf("day should never be skiped - chance is 0")
	}
}

// Test_Skip_Always : Tests if day is always skiped when chance is 1.
func Test_Skip_Always(t *testing.T) {

	// date to check
	dayIndex := 1
	date := time.Date(2021, 3, dayIndex, 0, 0, 0, 0, time.UTC)

	// chance
	daySkipChance := 1
	skipChance := [7]int{}
	skipChance[dayIndex] = daySkipChance

	// go through it 10k times to get approximate average
	result := [2]int{0, 0}
	for runIndex := 0; runIndex < 10000; runIndex++ {
		currentResult := skipDay(date, skipChance)
		if currentResult {
			result[0]++
		} else {
			result[1]++
		}
	}

	// validate
	if result[0] != 10000 || result[1] != 0 {
		t.Errorf("day should always be skiped - chance is 1")
	}
}
