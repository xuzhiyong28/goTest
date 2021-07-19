package go_grpc_example

import "testing"
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
