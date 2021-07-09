package gopsutil

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
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


func Demo3(){
	//查看CPU负载
	avg, _ := load.Avg()
	fmt.Println(avg)
}

// 内存信息
func Demo4(){
	memory, _ := mem.VirtualMemory()
	fmt.Println(memory)
}

// host info
func Demo5() {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}