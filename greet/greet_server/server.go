package main

import (
	"com.grpc.tleu/greet/greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"net"
	"strconv"
	"time"
)

type Server struct {
	greetpb.UnimplementedGreetServiceServer
}


func (s *Server) GreetManyTimes(request *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	number := int(request.GetGreeting().GetNumber())
	for number > 1 {
		for i := 2; number >= i;{
			if big.NewInt(int64(i)).ProbablyPrime(0){
				if number % i == 0 {
					number = number / i
					response := &greetpb.GreetManyTimesResponse{Result: fmt.Sprintf(strconv.Itoa(i))}
					if err := stream.Send(response); err != nil {
						log.Fatalf("error while sending greet many times responses: %v", err.Error())
					}
					time.Sleep(time.Second)
				}else{
					i++
				}
			}else{
				i++
			}
		}
	}
	return nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
