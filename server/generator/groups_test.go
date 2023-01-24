package generator

import (
	"testing"
)

// Test_SplitDaysIntoGroups : Tests if GroupId is set to all days.
func Test_SplitDaysIntoGroups_GroupId(t *testing.T) {

	days := []*WorkingDay{{Hours: 5}, {Hours: 1}, {Hours: 3}}
	groups := []*Group{{GroupId: "A", Hours: 9}}

	result := splitDaysIntoGroups(days, groups)

	for _, day := range result {
		if day.GroupId != groups[0].GroupId {
			t.Errorf("Day does not have correct GroupId.")
		}
	}
}

// Test_SplitDaysIntoGroups : Tests if GroupId is set to all days for multiple groups.
func Test_SplitDaysIntoGroups_GroupId_Multiple_Groups(t *testing.T) {

	days := []*WorkingDay{{Hours: 5}, {Hours: 1}, {Hours: 3}}
	groups := []*Group{{GroupId: "A", Hours: 5}, {GroupId: "B", Hours: 4}}

	result := splitDaysIntoGroups(days, groups)

	if result[0].GroupId != groups[0].GroupId {
		t.Errorf("Day does not have correct GroupId.")
	}
	if result[1].GroupId != groups[1].GroupId {
		t.Errorf("Day does not have GroupId.")
	}
	if result[2].GroupId != groups[1].GroupId {
		t.Errorf("Day does not have GroupId.")
	}
}

// Test_SplitDaysIntoGroups : Tests if hours get split when part of the day is one group and the rest in another group.
func Test_SplitDaysIntoGroups_Split_Days(t *testing.T) {

	days := []*WorkingDay{{Hours: 5}, {Hours: 1}}
	groups := []*Group{{GroupId: "A", Hours: 3}, {GroupId: "B", Hours: 3}}

	result := splitDaysIntoGroups(days, groups)

	if result[0].GroupId != groups[0].GroupId {
		t.Errorf("Day does not have correct GroupId.")
	}
	if result[0].Hours != groups[0].Hours {
		t.Errorf("Day does not have correct Hours.")
	}
	if result[1].Hours != groups[1].Hours {
		t.Errorf("Day does not have Hours.")
	}
}
