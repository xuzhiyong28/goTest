package panic

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	panicDemo1()
}

func TestDemo2(t *testing.T) {
	panicDemo2()
}

func TestDemo3(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			if i == 5 {
				panic("错误 5")
			} else {
				fmt.Println(i)
			}

		}(i)
	}
	wg.Wait()
}

func TestDemo4(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("happend error")
		}
	}()
	x := 1
	y := 0
	result := x / y
	fmt.Println(result)
}

// 方法里面发生panic
func TestDemo5(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("happend error")
		}
	}()
	result, err := division(1, 0)
	if err == nil {
		fmt.Println(result)
	}
}

// 协程出现panic,这个程序退出，所以协程里面也要做panic处理
func TestDemo6(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("happend error")
		}
	}()
	go func() {
		result := 1 / 0
		fmt.Println(result)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("End ...")
}

func division(x, y int) (int, error) {
	z := x / y
	return z, nil
}
