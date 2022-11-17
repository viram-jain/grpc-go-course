package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 15,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Println("Sum =>", res.Result)
}
