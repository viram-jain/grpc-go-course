package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-course/greet/proto"
)

var addr string = "localhost:5001"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	doGreetManyTimes(client)
	doLongGreet(client)
}
