package main

import (
	"context"
	"io"
	"log"

	pb "github.com/adenierosorto/grpc-go/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPromes was invoked")

	req := &pb.PrimesRequest{
		Number: 12390392840,
	}
	steam, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling primes: %v\n", err)
	}
	for {
		res, err := steam.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}
}
