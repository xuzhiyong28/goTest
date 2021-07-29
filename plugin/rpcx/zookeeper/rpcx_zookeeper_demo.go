package zookeeper

import (
	"context"
	"example/plugin/rpcx"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

const (
	zk_addr = "localhost:2181"
	zk_path = "/rpc/xuzyService"
)

func ServiceDemo1(addr string) {
	s := server.NewServer()
	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + addr,
		ZooKeeperServers: []string{zk_addr},
		BasePath:         zk_path,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   5 * time.Second,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
	s.RegisterName("Arith", new(rpcx.Arith), "")
	s.Serve("tcp", addr)
}

func ServiceMuitDemo(){
	go ServiceDemo1("localhost:8972")
	go ServiceDemo1("localhost:8973")
	select {}
}



func ClientDemo1() {
	d, _ := client.NewZookeeperDiscovery(zk_path, "Arith", []string{"localhost:2181"}, nil)
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
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
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}
}
