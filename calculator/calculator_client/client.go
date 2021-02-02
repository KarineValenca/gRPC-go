package main

import (
	"context"
	"fmt"
	"log"

	"github.com/KarineValenca/gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Started client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewSumServiceClient(conn)

	req := &calculatorpb.CalculatorRequest{
		Value: &calculatorpb.Values{
			FirstNumber:  1,
			SecondNumber: 2,
		},
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalln("Failed to do the request: %v", err)
	}
	log.Println("Response from calculator:", res.Result)
}
