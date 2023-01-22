package generator

import "testing"

// Test_TotalHours : Tests if total hours generated are the same as requested.
func Test_TotalHours(t *testing.T) {

	request := &Request{A_Hours: 20, B_Hours: 17}
	result := Generate(request)

	var generatedHours int = 0
	for _, workingDay := range result {
		generatedHours += workingDay.Hours
	}

	if generatedHours != (request.A_Hours + request.B_Hours) {
		t.Errorf("Generated hours not the same as target hours.")
	}
}
