package sync_pool

// https://mp.weixin.qq.com/s/6Nx7IGFU_FbM5AOdUzmvcw
import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Demo1() {
	var pool = &sync.Pool{
		New: func() interface{} {
			return "new world"
		},
	}
	value := "Hello,xuzhiyong"
	pool.Put(value)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	pool.Put("Hello,xuzhiyong2")
	fmt.Println(pool.Get())
}

// 用来统计实例真正创建的次数
var numCalcsCreated int32



func Demo2() {
	bufferPool := &sync.Pool{
		New: func() interface{} {
			atomic.AddInt32(&numCalcsCreated, 1)
			buffer := make([]byte, 1024)
			return &buffer
		},
	}
	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0 ; i < numWorkers ; i++ {
		go func() {
			defer wg.Done()
			buffer := bufferPool.Get()
			_ = buffer.(*[]byte)
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()
	fmt.Printf("%d buffer objects were created.\n", numCalcsCreated)
}
