// +build server

package main

import (
	"consultest/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedCountingServer
}

func (s *server) GetMessage(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Received: %v", in.GetText())
	return &pb.Message{Text: "Hello " + in.GetText(), Number: in.Number}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCountingServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
