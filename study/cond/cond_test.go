package cond

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestDemo0(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()
			log.Printf("运动员#%d已准备就绪\n", i)
			// 广播
			c.Broadcast()
		}(i)
	}
	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()
	log.Println("所有运动员准备就绪。比赛开始...")
}
