package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTicker(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			{
				fmt.Println("doing....")
			}
		}
	}
}

func TestNewTimer(t *testing.T) {
	timer := time.NewTimer(2 * time.Second)
	conn := make(chan interface{}, 1)
	go func() {
		time.Sleep(20 * time.Second)
		conn <- struct{}{}
	}()
	for {
		select {
		case <-timer.C:
			{
				fmt.Println("doing....")
				timer.Reset(2 * time.Second)
			}
		case <-conn:
			timer.Stop()
			return
		}
	}
}
