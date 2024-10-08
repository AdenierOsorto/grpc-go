package main

import (
	"context"
	"log"

	pb "github.com/adenierosorto/grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Print("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FisrtName: "Adenier",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
