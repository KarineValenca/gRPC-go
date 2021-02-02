package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/KarineValenca/gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Sum function was invocked with %v", req)
	firstNumber := req.GetValue().GetFirstNumber()
	secondNumber := req.GetValue().GetSecondNumber()

	result := firstNumber + secondNumber

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("started server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("could not start the server %v", err)
	}
}
