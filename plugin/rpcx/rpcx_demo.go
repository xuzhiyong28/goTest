package rpcx

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"log"
)

type Arith int
type Args struct {
	A int
	B int
}
type Reply struct {
	C int
}

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func ServiceDemo1() {
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", "localhost:8972")
}

func ClientDemo1() {
	d, err := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "") //点对点的方式-客户端直连服务器来获取服务地址
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	args := &Args{
		A: 10,
		B: 20,
	}
	reply := &Reply{}
	err = xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}

// 异步方式调用rpc
func ClientDemo1_2(){
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "") //点对点的方式-客户端直连服务器来获取服务地址
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	args := &Args{
		A: 10,
		B: 20,
	}
	reply := &Reply{}
	call, err := xclient.Go(context.Background(),"Mul",args,reply,nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	replyCall := <-call.Done //等待结果返回
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
}
