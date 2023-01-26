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

// Test_EmptyGroups : Tests if code does not fail when request has no groups.
func Test_EmptyGroups(t *testing.T) {

	request := &Request{Groups: []*Group{}}
	result := Generate(request)

	if len(result) != 0 {
		t.Errorf("There should be no generated hours.")
	}
}

// Test_NilGroups : Tests if code does not fail when request has no groups.
func Test_NilGroups(t *testing.T) {

	request := &Request{Groups: nil}
	result := Generate(request)

	if len(result) != 0 {
		t.Errorf("There should be no generated hours.")
	}
}
