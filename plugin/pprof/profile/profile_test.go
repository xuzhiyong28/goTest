package profile

import (
	"fmt"
	"testing"
)
import "github.com/pkg/profile"


// https://www.toutiao.com/a6724648030403822084/?channel=&source=search_tab
func TestPorfile(t *testing.T) {
	defer profile.Start().Stop()
	sl := MakeSlice()
	fmt.Printf("sum = %d\n", SumSlice(sl))
}


func MakeSlice() []int {
	sl := make([]int, 10000000)
	for idx := range sl {
		sl[idx] = idx
	}
	return sl
}
func SumSlice(sl []int) int {
	sum := 0
	for _, x := range sl {
		sum += x
	}
	return sum
}