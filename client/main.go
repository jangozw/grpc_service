package main

import (
	pb "github.com/jangozw/grpc_service/source"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCallServiceMethodClient(conn)
	req := &pb.Request{ServiceName:"Test", MethodName:"Test", Params:"test"}

	r, err := c.Invoke(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)

}
