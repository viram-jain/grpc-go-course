package main

import (
	"context"
	"log"

	pb "grpc-go-course/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		log.Printf("Sending number: %v\n", number)

		stream.Send(&pb.AverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error in receiving response: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)
}
