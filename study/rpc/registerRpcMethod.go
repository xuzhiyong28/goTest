package rpc

import "net/rpc"

// rpc服务1
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// rpc服务2
type DoServiceInterface = interface {
	Do(request string, reply *string) error
}

func RegisterRpc(rpcName string,svc interface{}) error {
	return rpc.RegisterName(rpcName,svc)
}

