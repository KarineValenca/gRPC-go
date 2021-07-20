package main

import (
	"context"
	"fmt"
	"io"
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

	calculateSum(c)
	calculatePrimeDecomposition(c)
}

func calculateSum(c calculatorpb.SumServiceClient) {
	req := &calculatorpb.CalculatorRequest{
		Value: &calculatorpb.Values{
			FirstNumber:  1,
			SecondNumber: 2,
		},
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to do the request: %v", err)
	}
	log.Println("Response from calculator:", res.Result)
}

func calculatePrimeDecomposition(c calculatorpb.SumServiceClient) {
	req := &calculatorpb.PrimeDecompositionRequest{
		Number: 120,
	}

	resStream, err := c.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reaind stream: %v", err)
		}
		log.Printf("Reponse from PrimeDecomposition: %v", msg.GetResult())
	}
}
