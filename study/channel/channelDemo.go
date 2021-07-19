package channel

import (
	"fmt"
	"math/rand"
	"time"
)

// 非缓冲channel : 如果没 goroutine 读取接收者<-ch ，那么发送者ch<- 就会一直阻塞
// 这里会报错 fatal error: all goroutines are asleep - deadlock! 就是因为 ch<-1 发送了，但是同时没有接收者，所以就发生了阻塞
func Demo1() {
	ch := make(chan int) //后面没有指定管道数量的就是非缓冲的channel
	//ch <- 1
	go func() {
		for {
			select {
			case i := <-ch:
				fmt.Println("this  value of unbuffer channel", i)
			}
		}
	}()
	ch <- 1 //放在下面就不会报错
	time.Sleep(1 * time.Second)
}

// 缓冲channel
func Demo2() {
	ch := make(chan int, 3)
	go func() {
		for {
			ch <- rand.Intn(10)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case i := <-ch:
				fmt.Println("this  value of unbuffer channel", i)
			}
		}
	}()
	time.Sleep(10 * time.Second)
}
