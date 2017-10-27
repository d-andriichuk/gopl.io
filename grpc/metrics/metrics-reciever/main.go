package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "gopl.io/grpc/metrics/metrics"
)

const (
	port = ":50051"
)

type server struct{}

// Exchange method
func (s *server) Exchange(ctx context.Context, in *pb.MetricsModel) (*pb.MetricsResponse, error) {
	log.Println(in.Metrics)
	return &pb.MetricsResponse{Status: "Metrics received"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterMetricsServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
