package main

// https://www.toutiao.com/a6724648030403822084/?channel=&source=search_tab

import (
	pro "example/plugin/pprof/pprof-profile"
	"fmt"
	"github.com/pkg/profile"
)

func main(){
	defer profile.Start().Stop()
	sl := pro.MakeSlice()
	fmt.Printf("sum = %d\n", pro.SumSlice(sl))
}
