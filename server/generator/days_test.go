package generator

import (
	"testing"
	"time"
)

// Test_RandomizeDays : Tests simple scenario with fewer hours are there are days. The first days must be working.
func Test_RandomizeDays_LessHours(t *testing.T) {
	// we have hours for 5 separate days
	hours := [5]int{4, 5, 3, 1, 10}
	// let set fixed now
	now := time.Date(2020, 4, 13, 0, 0, 0, 0, time.UTC)

	// lets generate which days will be working
	days := randomizeDays(hours[:], now)

	// count working days
	var workingDaysCont int = 0
	for _, workingDay := range days {
		if workingDay.IsWorking {
			workingDaysCont++
		}
	}

	// for this case lengths must match
	if workingDaysCont != len(hours) {
		t.Errorf("days is not the same length as hours")
	}
}

// Test_RandomizeDays : Tests scenario with more hours are there are days. All days should be working.
func Test_RandomizeDays_MoreHours(t *testing.T) {
	// we have hours for 40 separate days
	hours := [40]int{4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10, 4, 5, 3, 1, 10}
	// let set fixed now
	now := time.Date(2020, 4, 13, 0, 0, 0, 0, time.UTC)

	// lets generate which days will be working
	days := randomizeDays(hours[:], now)

	// check working days
	for _, workingDay := range days {
		if !workingDay.IsWorking {
			t.Errorf("not all days are working days")
			break
		}
	}
}
