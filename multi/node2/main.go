// SERVER SIDE

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
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	message := in.Data
  return &pb.HelloReply{Name: (message)}, nil

	const (
	address = "localhost:50052"
	)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	r, err := c.SendHello(context.Background(), &pb.HelloRequest{
		Data : message,
  })

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
