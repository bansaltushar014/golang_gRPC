package main

import (
	"context"
	"log"
	"time"

	pb "github.com/example/golanggrpc/proto"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalf("Failed1 %v", err)
	}
	log.Printf("%s", res.Message)
}
