package server

import (
	"context"
	pb "example/project/go-grpc-example/proto"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

const (
	PORT        = "9001"
	PORT_STREAM = "9002"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func DemoServer1() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		fmt.Printf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}



//实现StreamServiceServer接口
type StreamService struct{}

func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
	return nil
}
func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "gPRC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
	return nil
}

func DemoServer2() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamService{})
	lis, err := net.Listen("tcp", ":"+PORT_STREAM)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
