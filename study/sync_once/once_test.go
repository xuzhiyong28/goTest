package sync_once

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var count int64
var wg sync.WaitGroup

func TestDemo(t *testing.T) {
	var once sync.Once
	once.Do(func() {
		count = 10
	})
	for i := 0 ; i < 100 ; i++ {
		once.Do(func() {
			count = rand.Int63n(100)
		})
	}
	time.Sleep(10 * time.Second)
	log.Println(fmt.Sprintf("count = %d" , count))
}
