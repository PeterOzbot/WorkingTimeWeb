package generator

import "testing"

// Test_TotalHours : Tests if total hours generated are the same as requested.
func Test_TotalHours(t *testing.T) {

	request := &Request{Groups: []*Group{{Hours: 20}, {Hours: 17}}}
	result := Generate(request)

	var generatedHours int = 0
	for _, workingDay := range result {
		generatedHours += workingDay.Hours
	}

	if generatedHours != (request.Groups[0].Hours + request.Groups[1].Hours) {
		t.Errorf("Generated hours not the same as target hours.")
	}
}
