package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestDemo1(t *testing.T) {
	var count uint64
	var wg sync.WaitGroup
	for i := 0 ; i < 10 ; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddUint64(&count,1)
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
	countVal.Store([]int{1,3,5,7,9})
}
