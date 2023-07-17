package main

import (
	"context"
	"log"

	pb "github.com/adenierosorto/grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 1,
	})
	if err != nil {
		log.Fatalf("could not sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}
