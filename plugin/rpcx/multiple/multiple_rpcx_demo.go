package multiple

import (
	"context"
	"example/plugin/rpcx"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"log"
	"time"
)

func ServiceDemo1() {
	// 定义两个服务端，客户端采用负载均衡的方式调用
	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(rpcx.Arith), "")
		s.Serve("tcp", "localhost:8972")
	}()
	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(rpcx.Arith), "")
		s.Serve("tcp", "localhost:8973")
	}()

	select {}
}

func ClientDemo1() {
	d, _ := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: "localhost:8972"}, {Key: "localhost:8973"}})
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
	// 更新注册的实例
	//dd := d.(*client.MultipleServersDiscovery)
	//dd.Update([]*client.KVPair{{Key: "localhost:8972"}, {Key: "localhost:8973"}})
	defer xclient.Close()
	args := &rpcx.Args{
		A: 10,
		B: 20,
	}
	for {
		reply := &rpcx.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		log.Printf("%d * %d = %d\n", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}
}
