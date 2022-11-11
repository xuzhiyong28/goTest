package main

import (
	"crypto/tls"
	"crypto/x509"
	"example/project/go-grpc-example/proto"
	"example/project/go-grpc-example/server/three_server/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	serverPem := dir + "\\project\\go-grpc-example\\server\\three_server\\cert\\server.pem"
	serverKey := dir + "\\project\\go-grpc-example\\server\\three_server\\cert\\server.key"
	cert, err := tls.LoadX509KeyPair(serverPem, serverKey)
	if err != nil {
		log.Fatalf("加载服务端证书失败, err: %v\n", err)
	}
	certPool := x509.NewCertPool()
	caPem := dir + "\\project\\go-grpc-example\\server\\three_server\\cert\\ca.pem"
	ca, err := ioutil.ReadFile(caPem)
	if err != nil {
		log.Fatalf("读取公钥文件失败: %v\n", err)
	}
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterProductServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}
