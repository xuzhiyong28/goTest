package go_crawler

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestDemo1(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		body := fetch("https://movie.douban.com/top250?start=" + strconv.Itoa(25*i))
		parseBody(body)
	}
	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
}

func TestDemo2(t *testing.T) {
	//使用并发
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			body := fetch("https://movie.douban.com/top250?start=" + strconv.Itoa(25*i))
			parseBody(body)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
}
