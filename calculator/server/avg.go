package main

import (
	"io"
	"log"

	pb "github.com/adenierosorto/grpc-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("avg function was invoked")

	sum := 0
	count := 0

	for {
		req, err := stream.Recv()

		if err != io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receiving number: %d\n", req.Number)

		sum += int(req.Number)
		count++
	}
}
