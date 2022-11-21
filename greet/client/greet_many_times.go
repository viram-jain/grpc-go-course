package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Viram",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in GreetManyTimes %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("GreetManyTimes %s\n", msg.Result)
	}
}
