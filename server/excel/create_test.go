package excel

import (
	"testing"
	"time"
)

// Test_
func Test_(t *testing.T) {
	request := []*Request{{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Hours: 1, GroupId: "1"}}

	Create(request)
	//if result != nil {
	t.Errorf("Implement test")
	//}
}
