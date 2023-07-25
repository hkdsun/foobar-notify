package main

import (
	"context"
	"log"
	"net"
	"os/exec"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) TriggerRescan(ctx context.Context, req *RescanRequest) (*RescanResponse, error) {
	cmd := exec.Command(req.FoobarPath, "/command", "Rescan")
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return &RescanResponse{Success: true}, nil
}

func main() {
	// Start the gRPC serve5r
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterFoobarServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
