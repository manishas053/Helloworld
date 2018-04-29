package main

import (
	"log"
	"net"

	pb "github.com/manishas053/Helloworld/multi/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) SendHello(ctx context.Context, in2 *pb.HelloRequest) (*pb.HelloReply, error) {

	message := in2.Data

	return &pb.HelloReply{Name: (message)}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
