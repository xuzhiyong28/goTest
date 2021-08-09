package zhihu_queue

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
	知乎的一个发布订阅 https://zhuanlan.zhihu.com/p/255207686
**/

func TestClient_Simple(t *testing.T) {
	b := NewClient()
	b.SetConditions(10)
	topic := "主题001"
	ch1, _ := b.Subscribe(topic)
	ch2, _ := b.Subscribe(topic)
	go func() {
		for {
			e := b.GetPayLoad(ch1)
			fmt.Println("ch1 = " , e)
		}
	}()
	go func() {
		for {
			e := b.GetPayLoad(ch2)
			fmt.Println("ch2 = " , e)
		}
	}()
	b.Publish(topic,"主题001 - Go语言交流-1") // 指定主题发送消息
	b.Publish(topic,"主题001 - Go语言交流-2")
	b.Publish(topic,"主题001 - Go语言交流-3")
	b.Publish(topic,"主题001 - Go语言交流-4")
	b.Publish(topic,"主题001 - Go语言交流-5")
	b.Publish(topic,"主题001 - Go语言交流-6")
	b.Publish(topic,"主题001 - Go语言交流-7")
	time.Sleep(10 * time.Second)
}


func TestClient(t *testing.T) {
	b := NewClient()
	b.SetConditions(100)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		topic := fmt.Sprintf("Golang梦工厂%d", i)
		payload := fmt.Sprintf("asong%d", i)
		ch, err := b.Subscribe(topic)
		if err != nil {
			t.Fatal(err)
		}
		wg.Add(1)
		go func() {
			e := b.GetPayLoad(ch)
			if e != payload {
				t.Fatalf("%s expected %s but get %s", topic, payload, e)
			}else {
				fmt.Println("value = " , e)
			}
			if err := b.Unsubscribe(topic, ch); err != nil {
				t.Fatal(err)
			}
			wg.Done()
		}()

		if err := b.Publish(topic, payload); err != nil {
			t.Fatal(err)
		}
	}
	wg.Wait()
}
