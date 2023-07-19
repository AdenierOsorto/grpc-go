package main

import (
	"io"
	"log"

	pb "github.com/adenierosorto/grpc-go/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()
		res := "hello " + req.FisrtName + "!"
		if err != io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receiving: %v\n", req.FisrtName)
	}
}
