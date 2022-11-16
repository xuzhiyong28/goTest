package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestDemo1(t *testing.T) {
	var count uint64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func TestDemo2(t *testing.T) {
	var count int64 = 100
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, -1)
		}()
	}
	wg.Wait()
	fmt.Printf("count : %v \n", count)

	// 比较并替换
	bol := atomic.CompareAndSwapInt64(&count, 50, 20)
	fmt.Println(bol)
}

func TestDemo3(t *testing.T) {
	var c int32
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := atomic.LoadInt32(&c)
			if !atomic.CompareAndSwapInt32(&c, tmp, (tmp + 1)) {
				fmt.Println("修改失败")
			}
		}()
	}
	wg.Wait()
	fmt.Printf("count : %v\n", c)
}

func TestDemo4(t *testing.T) {
	var countVal atomic.Value
	countVal.Store([]int{1, 3, 5, 7, 9})
}

func TestDemo5(t *testing.T) {
	ch := make(chan int)
	for num := range ch {
		fmt.Println("num = ", num)
	}
}

func TestDemo6(t *testing.T) {
	newWorkCh := make(chan string)
	go func() {
		for {
			select {
			case req := <-newWorkCh:
				time.Sleep(time.Second * 2)
				fmt.Println(req)
			}
		}
	}()
	newWorkCh <- "1"
	newWorkCh <- "2"
	newWorkCh <- "3"
	newWorkCh <- "4"
	newWorkCh <- "5"
}
