package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetSenbeisByNum(ctx context.Context, in *NumGetSenbei) (*Senbeis, error) {
	log.Printf("Received: %v", in.Num)
	return &Senbeis{
		Senbeis: []*Senbei{
			{Price: 100, SenbeiType: "first one"},
			{Price: 200, SenbeiType: "second one"},
		},
	}, nil
}

func (s *server) GetSenbeisByType(ctx context.Context, in *SenbeiTypes) (*Senbeis, error) {
	log.Printf("Received: %v", in.SenbeiTypes)
	return &Senbeis{
		Senbeis: []*Senbei{
			{Price: 100, SenbeiType: "first one"},
			{Price: 200, SenbeiType: "second one"},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterSenbeiServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
