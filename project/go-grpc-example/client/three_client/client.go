package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"example/project/go-grpc-example/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	clientPem := dir + "\\project\\go-grpc-example\\client\\three_client\\cert\\client.pem"
	clientKey := dir + "\\project\\go-grpc-example\\client\\three_client\\cert\\client.key"
	cert, err := tls.LoadX509KeyPair(clientPem, clientKey)
	if err != nil {
		log.Fatalf("加载客户端证书失败, err: %v\n", err)
	}
	certPool := x509.NewCertPool()
	caPem := dir + "\\project\\go-grpc-example\\client\\three_client\\cert\\ca.pem"
	ca, err := ioutil.ReadFile(caPem)
	if err != nil {
		log.Fatalf("读取公钥文件失败: %v\n", err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}
	defer conn.Close()
	prodClient := proto.NewProductServiceClient(conn)
	prodRes, err := prodClient.GetProductStock(context.Background(), &proto.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(prodRes.ProdStock)
}
