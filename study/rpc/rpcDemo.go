package rpc

import (
	"fmt"
	"net"
	"net/rpc"
)

// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-01-rpc-intro.html

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

type DoService struct {

}

func (p *DoService) Do(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

/***
rpc服务端
*/
func SevletDemo() {
	rpc.RegisterName("HelloService", new(HelloService))
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println(err)
	}
	rpc.ServeConn(conn)
}

/****
客户端请求rpc服务
*/
func ClientDemo() {
	//通过rpc.Dial拨号RPC服务
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println(err)
	}
	var reply string
	//通过client.Call调用具体的RPC方法
	err = client.Call("HelloService.Hello", "hello", reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}

/***
基于设计模式的rpc服务
*/
func SevletDemo2() {
	RegisterRpc("HelloService", new(HelloService))
	RegisterRpc("DoService", new(DoService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go rpc.ServeConn(conn)
	}
}
