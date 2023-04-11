package main

import (
	"log"

	pb "github.com/example/golanggrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	// grpc dial creates the client connection to provided string
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Ankit", "Alice", "Bob"},
	}

	// CallSayHello(client)
	// CallSayHelloServerStreaming(client, names)
	// callSayHelloClientStreaming(client, names)
	callHelloBidirectionalStream(client, names)
}
