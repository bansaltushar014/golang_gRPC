package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/example/golanggrpc/proto"
)

func CallSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	ctx, cancel := context.WithTimeout(context.Background(), 13*time.Second)
	defer cancel()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("Failed1 %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished!")
}
