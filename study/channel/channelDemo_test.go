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

// 错误示范1 ： 在channel关闭时再从channel获取数据会获取到空值
func TestDemo7(t *testing.T) {
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				task := <-ch // 如果ch关闭了，这里还能获取到值，只是是类型的默认值，他并不会停止运行
				// 正确做法，需要判断是否关闭
				//task , ok := <-ch
				//if !ok {
				//	break
				//}
				fmt.Println(task)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// channel泄露
// 由于 select 命中了超时逻辑，导致通道没有消费者（无接收操作），而其定义的通道为无缓冲通道，因此 goroutine 中的ch <- "job result"操作会一直阻塞，最终导致 goroutine 泄露。
func TestDemo8(t *testing.T) {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()
	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		return
	}
}
