package multiple

import (
	"fmt"
	"testing"
	"time"
)

func TestServiceDemo1(t *testing.T) {
	ServiceDemo1()
}

func TestClientDemo1(t *testing.T) {
	ClientDemo1()
}


func TestOther(t *testing.T) {
	var stopCh chan struct{}
	stopCh = make(chan struct{})
	go func(c <-chan struct{}) {
		fmt.Println("ready1")
		<-c //get
		fmt.Println("ready2")
	}(stopCh)
	close(stopCh)
	time.Sleep(10 * time.Second)
}