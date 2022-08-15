package service

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type SocketServer struct {
	BaseService
	proto    string
	addr     string
	listener net.Listener
}

// 重写 OnStart 方法
func (s *SocketServer) OnStart() error {
	ln, err := net.Listen(s.proto, s.addr)
	if err != nil {
		return err
	}
	s.listener = ln
	go s.acceptConnectionsRoutine()
	select {}
	return nil
}

func (s *SocketServer) acceptConnectionsRoutine() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn)
	}
}

// 重写 OnStop 方法
func (s *SocketServer) OnStop() {
	if err := s.listener.Close(); err != nil {
		log.Fatal(err)
	}
}

func NewSocketServer(protoAddr string) Service {
	proto, addr := ProtocolAndAddress(protoAddr)
	s := &SocketServer{
		proto:    proto,
		addr:     addr,
		listener: nil,
	}
	s.BaseService = *NewBaseService("ABCIServer", s)
	return s
}

func ProtocolAndAddress(listenAddr string) (string, string) {
	protocol, address := "tcp", listenAddr
	parts := strings.SplitN(address, "://", 2)
	if len(parts) == 2 {
		protocol, address = parts[0], parts[1]
	}
	return protocol, address
}
