package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Demo1(){
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

