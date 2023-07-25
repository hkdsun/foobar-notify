package main

import (
	"context"
	"flag"
	"foobar/proto"
	"log"
	"net"
	"os/exec"

	"google.golang.org/grpc"
)

type server struct {
	scriptPath string

	proto.UnimplementedFoobarServiceServer
}

func (s *server) TriggerRescan(ctx context.Context, req *proto.RescanRequest) (*proto.RescanResponse, error) {
	cmd := exec.Command(s.scriptPath)
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return &proto.RescanResponse{Success: true}, nil
}

func main() {
	// Define command line flags
	scriptPath := flag.String("scriptPath", "", "path to foobar notification script")

	// Parse command line flags
	flag.Parse()

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterFoobarServiceServer(s, &server{
		scriptPath: *scriptPath,
	})

	// Serve the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
