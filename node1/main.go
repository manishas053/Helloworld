//client

package main

import (
	"log"

	pb "github.com/manishas053/Helloworld/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Data: "Hello, world",
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
