package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/shanth1/golang-microservices/grpc/api"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedAdderServer
}

func (s *GRPCServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdderServer(s, &GRPCServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
