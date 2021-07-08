package gopsutil

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)
//http://godoc.org/github.com/shirou/gopsutil
//https://www.cnblogs.com/nickchen121/p/11517451.html

func Demo1(){
	//获取cpu信息
	cpuInfos, _ := cpu.Info()
	for _ , ci := range cpuInfos {
		fmt.Println(ci)
	}
}


func Demo2(){
	//定时任务输出cpu使用率
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
			case <-ticker.C : {
				per, _ := cpu.Percent(1 * time.Second, true)
				fmt.Printf("CPU Percent: %f\n", per)
			}
		}
	}
}
