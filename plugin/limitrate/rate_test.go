package limitrate

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"testing"
	"time"
)
// https://zhuanlan.zhihu.com/p/89820414

func TestDemo1(t *testing.T) {
	// r - 代表每秒可以向Token桶中产生多少token
	// b - b代表Token桶的容量大小
	limiter := rate.NewLimiter(2, 5)
	ctx := context.Background()
	start := time.Now()
	// 要处理二十个事件
	for i := 0; i < 20; i++ {
		//limiter.Wait(ctx)
		limiter.WaitN(ctx,1)	// 如果桶内的token小于1则等待
		//TODO do something
	}
	fmt.Println(time.Since(start)) // output: 7.501262697s （初始桶内5个和每秒2个token）
}


func TestDemo(t *testing.T) {
	limiter := rate.NewLimiter(1, 5)
	ctx := context.Background()
	start := time.Now()
	for i := 0 ; i < 100 ; i++ {
		fmt.Println(fmt.Sprintf("result = %v" , i))
		limiter.Wait(ctx)
	}
	fmt.Println(time.Since(start))
}