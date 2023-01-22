package generator

import (
	"testing"
	"workingtimeweb/server/core"
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

// Test_Reduce : Test if hours get reduced if there is a difference. The two lowest hours should be combined. The rest must remain the same. Combine 1 and 2 into second place.
func Test_Reduce_Simple(t *testing.T) {
	var hours []int = []int{6, 1, 7, 10, 2}
	var difference int = 1

	result := reduce(hours, difference)

	if len(result) != len(hours)-1 {
		t.Errorf("input hours should be reduced by one")
		return
	}
	if result[0] != 6 {
		t.Errorf("Result does not have correct amount of hours.")
		return
	}
	if result[1] != 3 {
		t.Errorf("Result does not have correct amount of hours.")
		return
	}
	if result[2] != 7 {
		t.Errorf("Result does not have correct amount of hours.")
		return
	}
	if result[3] != 10 {
		t.Errorf("Result does not have correct amount of hours.")
		return
	}
}

// Test_Reduce : Test if hours get reduced when the first hour is the lowest.
func Test_Reduce_First_Lowest(t *testing.T) {
	var hours []int = []int{1, 6}
	var difference int = 1

	var original_sum = core.Sum(hours)

	result := reduce(hours, difference)

	var result_sum = core.Sum(result)

	if original_sum != result_sum {
		t.Errorf("Total hours sum does not match.")
	}
}

// Test_Reduce : Test if hours get reduced when the second hour is the lowest.
func Test_Reduce_Second_Lowest(t *testing.T) {
	var hours []int = []int{6, 1}
	var difference int = 1

	var original_sum = core.Sum(hours)

	result := reduce(hours, difference)

	var result_sum = core.Sum(result)

	if original_sum != result_sum {
		t.Errorf("Total hours sum does not match.")
	}
}

// Test_Reduce : Test if hours get reduced when there is large difference.
func Test_Reduce_Large_Difference(t *testing.T) {
	var hours []int = []int{6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2, 6, 1, 7, 10, 2}
	var difference int = len(hours) - 31

	var original_sum = core.Sum(hours)

	result := reduce(hours, difference)

	var result_sum = core.Sum(result)

	if original_sum != result_sum {
		t.Errorf("Total hours sum does not match.")
	}
}

// Test_GenerateHours : Test if zero hours as input are handled.
func Test_GenerateHours_Zero(t *testing.T) {
	totalHours := 0
	choiceWeight := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	result := generateHours(totalHours, choiceWeight)

	if len(result) != 0 {
		t.Errorf("zero input should return NO hours")
	}
}

// Test_GenerateHours : Test if sum of generated hours matches requested input.
func Test_GenerateHours_Sum(t *testing.T) {
	totalHours := 12
	choiceWeight := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	result := generateHours(totalHours, choiceWeight)

	if len(result) == 0 {
		t.Errorf("should return some hours")
	}
	if core.Sum(result) != totalHours {
		t.Errorf("sum of generated hours do not match requested total hours")
	}
}
