package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/adenierosorto/grpc-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("sqrt was invoked with: %v\n", in)
	number := in.Number

	if number < 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}
	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
