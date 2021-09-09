package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) PrimeDecomposition(req *calculatorpb.PrimeDecompositionRequest, stream calculatorpb.SumService_PrimeDecompositionServer) error {
	fmt.Printf("PrimeDecomposition function was invocked with %v", req)

	number := req.Number
	divisor := int32(2)

	var res calculatorpb.PrimeDecompositionReponse
	for number > 1 {
		if number%divisor == 0 {
			res.Result = divisor
			stream.Send(&res)
			number = number / divisor
		} else {
			divisor++
		}
	}

	return nil
}

func (s *server) Average(stream calculatorpb.SumService_AverageServer) error {
	fmt.Printf("Average function was invocked with \n")

	var sum int32
	var qtd float32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			average := float32(sum) / qtd
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Average: float64(average),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}

		sum += req.GetValue()
		qtd += 1
	}
}

func (s *server) FindMaximum(stream calculatorpb.SumService_FindMaximumServer) error {
	fmt.Printf("Find Maximum function was invocked \n")

	maximum := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		if req.Number > int32(maximum) {
			maximum = int(req.Number)
			if err := stream.Send(&calculatorpb.FindMaximumResponse{
				Maximum: int32(maximum),
			}); err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
				return err
			}
		}

	}

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
