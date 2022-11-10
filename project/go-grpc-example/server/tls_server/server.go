package main

import (
	"context"
	pb "example/project/go-grpc-example/proto"
	"example/project/go-grpc-example/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func main() {
	serverPemPath := utils.GetFileRealPath("project\\go-grpc-example\\conf\\server.pem")
	serveKeyPath := utils.GetFileRealPath("project\\go-grpc-example\\conf\\server.key")
	c, err := credentials.NewServerTLSFromFile(serverPemPath, serveKeyPath)
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
