package services

import (
	"context"
	pb "example/project/go-grpc-example/proto"
)

type ProdService struct{}

func (p ProdService) GetProductStock(ctx context.Context, req *pb.ProdRequest) (*pb.ProdResponse, error) {
	return &pb.ProdResponse{
		ProdStock: 100,
	}, nil
}
