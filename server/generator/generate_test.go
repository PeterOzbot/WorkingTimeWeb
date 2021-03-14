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
