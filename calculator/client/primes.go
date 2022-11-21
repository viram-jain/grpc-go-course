package main

import (
	"context"
	"io"
	"log"

	pb "grpc-go-course/calculator/proto"
)

func doPrimes(client pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := client.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in calling Primes %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %v\n", err)
		}

		log.Printf("Primes : %d\n", res.Result)
	}

}
