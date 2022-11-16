package main

import (
	"context"
	"log"

	pb "grpc-go-course/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Printf("doGreet method invoked")

	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Viram",
	})

	if err != nil {
		log.Fatalf("Could not send greet request %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
