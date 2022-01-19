package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDemo1(t *testing.T) {
	Demo1()
}

func TestDemo2(t *testing.T) {
	Demo2()
}

func TestDemo3(t *testing.T) {
	strChan := make(chan string, 3)
	syncChan := make(chan struct{}, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		<-syncChan
		fmt.Println("Received a sync signal and wait a second... [receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				//如果管道关闭则返回false
				break
			}
		}
		fmt.Println("Stopped..")
		wg.Done()
	}()
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Send:", elem, "[sender]")
			if elem == "c" {
				syncChan <- struct{}{}
				fmt.Println("sent a sync signal. [sender]")
			}
		}
		fmt.Println("Wait 2 seconds...[sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		wg.Done()
	}()
	wg.Wait()
}

func TestDemo4(t *testing.T) {
	mapChan := make(chan map[string]int, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"] ++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		wg.Done()
	}()
	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v .[sender]\n", countMap)
		}
		close(mapChan)
		wg.Done()
	}()
	wg.Wait()
}

func TestDemo5(t *testing.T) {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Present time: %v. \n", time.Now())
	expirationTime := <-timer.C
	fmt.Printf("Now time: %v. \n", expirationTime)
}

func TestDemo6(t *testing.T) {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()
	timeout := time.Millisecond * 1000
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}
		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Printf("Received:%v\n", e)
		case <-timer.C:
				fmt.Println("Timeout!")
		}
	}
}
