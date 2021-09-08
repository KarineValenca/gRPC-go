package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/KarineValenca/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Started client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	//fmt.Printf("Created cliente %f", c)
	//doUnary(c)
	//doServerStreaming(c)
	doClienteStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do an Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Karine",
			LastName:  "Valença",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Karine",
			LastName:  "Valença",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTime: %v", msg.GetResult())
	}

}

func doClienteStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LoongGreetRequest{
		&greetpb.LoongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Karine",
			},
		},
		&greetpb.LoongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Raj",
			},
		},
		&greetpb.LoongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lilloo",
			},
		},
		&greetpb.LoongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Kety",
			},
		},
		&greetpb.LoongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Beck",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling long greet: %v", err)
	}

	// we iterate over our slice and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from long greet: %v", err)

	}

	fmt.Printf("LongGreet response %v\n", resp)
}
