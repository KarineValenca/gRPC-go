package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	//calculateSum(c)
	//calculatePrimeDecomposition(c)
	//calculateAverage(c)
	calculateMaximum(c)
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

func calculateAverage(c calculatorpb.SumServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*calculatorpb.AverageRequest{
		{
			Value: 1,
		},
		{
			Value: 2,
		},
		{
			Value: 3,
		},
		{
			Value: 4,
		},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error while calling average %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from average: %v", err)
	}

	fmt.Printf("Average is: %v \n", resp.Average)

}

func calculateMaximum(c calculatorpb.SumServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("error while calling average %v", err)
	}

	requests := []*calculatorpb.FindMaximumRequest{
		{
			Number: 1,
		},
		{
			Number: 5,
		},
		{
			Number: 3,
		},
		{
			Number: 6,
		},
		{
			Number: 2,
		},
		{
			Number: 20,
		},
	}

	waitc := make(chan struct{})
	//we send a buch of messages to the client (go routine)
	go func() {
		//func to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req.Number)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	//we receive a bunch of messages from the client (go routine)
	go func() {
		//func to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Maximum value is %v\n", res.GetMaximum())
		}
		close(waitc)
	}()

	//block until everythins is done
	<-waitc
}
