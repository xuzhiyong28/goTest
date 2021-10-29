package loadbalance

import (
	"fmt"
	"testing"
)

func TestLoadBalanceFactory(t *testing.T) {
	weightLb := LoadBalanceFactory(2)
	weightLb.Add("a", "1")
	weightLb.Add("b", "2")
	weightLb.Add("c", "100")
	var count = make(map[string]int)
	for i := 0; i < 200000; i++ {
		weightRs, _ := weightLb.Get()
		count[weightRs]++
	}
	fmt.Println(count)
}
