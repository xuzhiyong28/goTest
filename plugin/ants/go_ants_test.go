package ants

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
	"testing"
)


type task struct {
	index int
	wg    *sync.WaitGroup
}

func (t *task) Do() {
	fmt.Println("任务编号 = " , t.index)
	t.wg.Done()
}

func TestDemo1(t *testing.T) {
	//创建一个容量为10的线程池
	p, _ := ants.NewPoolWithFunc(10, func(data interface{}) {
		t := data.(*task)
		t.Do()
	})

	//构造数据
	nums := make([]int, 10000, 10000)
	for i := range nums {
		nums[i] = rand.Intn(1000)
	}
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0 ; i < 100 ; i++ {
		p.Invoke(&task{
			index: i,
			wg: &wg,
		})
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
}
