package main

import (
	"context"
	pb "example/project/go-grpc-example/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

const (
	PORT = "9001"
)

type SearchService struct {
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":" + PORT)
	if err != nil {
		fmt.Printf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
