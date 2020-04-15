package generator

import "testing"

// Test : tests generator method
func Test(t *testing.T) {

	var targetHours int = 37
	result := Generate(targetHours)

	var generatedHours int = 0
	for _, workingDay := range result {
		generatedHours += workingDay.Hours
	}

	if generatedHours != targetHours {
		t.Errorf("Generated hours not the same as target hours.")
	}
}
