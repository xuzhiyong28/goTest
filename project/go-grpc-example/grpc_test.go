package go_grpc_example

import (
	"fmt"
	"testing"
	"time"
)
import cli "example/project/go-grpc-example/client"
import sli "example/project/go-grpc-example/server"

func TestDemo1_client(t *testing.T) {
	cli.DemoClient1()
}

func TestDemo1_server(t *testing.T) {
	sli.DemoServer1()
}

func TestDemo2_client(t *testing.T) {
	cli.DemoClient2()
}

func TestDemo2_server(t *testing.T) {
	sli.DemoServer2()
}


func TestT(t *testing.T) {
	ticker := time.NewTicker(time.Second).C
	for {
		select {
			case now := <-ticker:
				unix := now.Unix()
				offset := unix % 86400
				//offset /= 10
				fmt.Println(unix, offset)
		}
	}
}
