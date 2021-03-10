package generator

import (
	"testing"
)

// Test_Reduce : Test if hours remain unchanged if there is zero difference.
func Test_Reduce_ZeroDifference(t *testing.T) {
	var hours []int = []int{5, 1, 7, 10}
	var difference int = 0

	result := reduce(hours, difference)

	if len(result) != len(hours) {
		t.Errorf("result and hours do not match in length")
		return
	}
	for i, v := range result {
		if v != hours[i] {
			t.Errorf("Result do not match input hours.")
			return
		}
	}
}

// Test_Reduce : Test if hours get reduced if there is a difference. The two lowest hours should be combined. The rest must remain the same.
func Test_Reduce_Simple(t *testing.T) {
	var hours []int = []int{6, 1, 7, 10, 2}
	var difference int = 1

	result := reduce(hours, difference)

	if len(result) != len(hours)-1 {
		t.Errorf("input hours should be reduced by one")
		return
	}
	if result[0] != hours[0] {
		t.Errorf("Result does not have correct amount of hours.")
		return
	}
	if result[1] != hours[2] {
		t.Errorf("Result does not have correct amount of hours.")
		return
	}
	if result[2] != hours[3] {
		t.Errorf("Result does not have correct amount of hours.")
		return
	}
	if result[3] != hours[1]+hours[4] {
		t.Errorf("Result does not have correct amount of hours.")
		return
	}
}

// Test_GenerateHours : Test ...
func Test_GenerateHours(t *testing.T) {
	t.Errorf("implement")
}
