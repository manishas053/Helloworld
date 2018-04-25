// CLIENT SIDE

package main

import (
	"log"

	pb "github.com/manishas053/Helloworld/single/helloworld"


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
	message := "Helloworld!"

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{

		Data : message,

	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("Greeting: ", r.Name)

}
