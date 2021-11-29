package panic

import (
	"fmt"
	"sync"
	"testing"
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
				if err := recover() ; err != nil {fmt.Println(err)}
			}()
			if i == 5 {
				panic("错误 5")
			}else {
				fmt.Println(i)
			}

		}(i)
	}
	wg.Wait()
}
