package testing

import (
	"fmt"
	"testing"
)

func sum(x, y int) int {
	return x + y
}

// 子测试
func TestChildRun(t *testing.T) {
	testCalss := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zuri", "13:31"},     // incorrect location name
		{"12:31", "America/New_York", "7:31"}, // should be 07:31
		{"08:08", "Australia/Sydney", "18:08"},
	}
	for _, tc := range testCalss {
		t.Run(fmt.Sprintf("%s_%s", tc.gmt, tc.loc), func(t *testing.T) {
			t.Logf("%s , %s , %s", tc.gmt, tc.loc, tc.want)
		})
	}
}
