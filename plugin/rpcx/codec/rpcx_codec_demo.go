package codec

import (
	"context"
	"example/plugin/rpcx"
	jsoniter "github.com/json-iterator/go"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
	"log"
)

type JsoniterCodec struct {
}

func (c *JsoniterCodec) Decode(data []byte, i interface{}) error {
	return jsoniter.Unmarshal(data, i)
}

func (c *JsoniterCodec) Encode(i interface{}) ([]byte, error) {
	return jsoniter.Marshal(i)
}

func ServiceDemo() {
	share.Codecs[protocol.SerializeType(4)] = &JsoniterCodec{}
	s := server.NewServer()
	s.Register(new(rpcx.Arith), "")
	s.Serve("tcp", "localhost:8972")
}

func ClientDemo() {
	share.Codecs[protocol.SerializeType(4)] = &JsoniterCodec{}
	option := client.DefaultOption
	option.SerializeType = protocol.SerializeType(4)

	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &rpcx.Args{
		A: 10,
		B: 20,
	}

	reply := &rpcx.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
