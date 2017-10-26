package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "gopl.io/grpc/habr_api/habrapi"
)

var port = flag.String("port", ":50051", "Server port")

type server struct{}

// GetKarma method
func (s *server) GetKarma(ctx context.Context, in *pb.KarmaRequest) (*pb.KarmaResponse, error) {
	return &pb.KarmaResponse{Username: in.Username, Karma: 500}, nil
}

// PostArticle method
func (s *server) PostArticle(ctx context.Context, in *pb.PostArticleRequest) (*pb.PostArticleResponse, error) {
	return &pb.PostArticleResponse{
		Posted:    true,
		Url:       "www.habrahabr.com/article/326347",
		Time:      time.Now().String(),
		ErrorCode: "OK",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHabrApiServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
