package generator

import "testing"

// Test_TotalHours : Tests if total hours generated are the same as requested.
func Test_TotalHours(t *testing.T) {

	request := &Request{TotalHours: 37}
	result := Generate(request)

	var generatedHours int = 0
	for _, workingDay := range result {
		generatedHours += workingDay.Hours
	}

	if generatedHours != request.TotalHours {
		t.Errorf("Generated hours not the same as target hours.")
	}
}

// Test_TotalHours : Tests if generated working days have hours more than 0 if the day is working.
func Test_WorkingDasHaveHours(t *testing.T) {

	request := &Request{TotalHours: 37}
	result := Generate(request)

	for _, workingDay := range result {
		if workingDay.IsWorking && workingDay.Hours <= 0 {
			t.Errorf("Generated working day is working and does not have hours.")
			break
		}
	}
}
