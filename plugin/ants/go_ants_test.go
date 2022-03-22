package ants

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type task struct {
	index int
	wg    *sync.WaitGroup
}

func (t *task) Do() {
	fmt.Println("任务编号 = ", t.index)
	t.wg.Done()
}

func TestDemo1(t *testing.T) {
	//创建一个容量为10的线程池
	p, _ := ants.NewPoolWithFunc(10, func(data interface{}) {
		t := data.(*task)
		t.Do()
	})

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		p.Invoke(&task{
			index: i,
			wg:    &wg,
		})
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
}

// 官方测试案例1 ： https://github.com/panjf2000/ants/blob/master/README_ZH.md
// 使用的是ants自带的默认线程池 defaultAntsPool,看了下 容量 = 2147483647
func TestDemo0(t *testing.T) {
	defer ants.Release()
	runTimes := 1000
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Hello World!")
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
}

// 官方测试案例2 ： https://github.com/panjf2000/ants/blob/master/README_ZH.md
func TestDemo2(t *testing.T) {
	var sum int32
	var wg sync.WaitGroup
	runTimes := 1000
	myFunc := func(i interface{}) {
		n := i.(int32)
		atomic.AddInt32(&sum, n)
		fmt.Printf("run with %d\n", n)
	}
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}
