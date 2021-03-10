package core

import "testing"

// Test_Sum_CorrectAmount : Tests if array sum is correctly calculated.
func Test_Sum_CorrectAmount(t *testing.T) {
	var array []int = []int{5, 1, 2, 4}

	result := Sum(array)

	if result != 12 {
		t.Errorf("sum of the array is not correct")
		return
	}
}
