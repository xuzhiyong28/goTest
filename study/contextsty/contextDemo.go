package contextsty

import (
	"context"
	"fmt"
	"time"
)

/**
  context.WithCancel 调用cancel()方法后取消
**/
func ContextWithCancelDemo() {
	fmt.Println("start...")
	//可以使用带参数的context
	withValCtx := context.WithValue(context.Background(), "name", "xuzy")
	ctx, cancel := context.WithCancel(withValCtx)
	go func(ctx context.Context) {
		var i = 1
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("done", ctx.Value("name"))
				return
			default:
				fmt.Printf("work %d seconds \n", i)
			}
			i++
		}
	}(ctx)
	//模拟程序运行 - Sleep 5秒
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("end.")
}

func ContextWithTimeOutDemo(){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go func(ctx context.Context) {
		var i = 1
		for{
			time.Sleep(1 * time.Second)
			select {
			case <- ctx.Done():
				fmt.Println("done")
				return
			default:
				fmt.Printf("work %d seconds: \n", i)
			}
			i++
		}
	}(ctx)
	//模拟程序运行 - Sleep 10秒
	time.Sleep(10 * time.Second)
	cancel() // 3秒后将提前取消 doSth goroutine
}
