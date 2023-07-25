package main

import (
	"context"
	"flag"
	"foobar/proto"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	foobarPath string

	proto.UnimplementedFoobarServiceServer
}

func (s *server) TriggerRescan(ctx context.Context, req *proto.RescanRequest) (*proto.RescanResponse, error) {
	// Write to /tmp/debug
	err := ioutil.WriteFile("/tmp/debug", []byte(s.foobarPath), 0644)
	if err != nil {
		return nil, err
	}
	return &proto.RescanResponse{Success: true}, nil
}

func main() {
	// Define command line flags
	foobarPath := flag.String("foobarPath", "", "path to foobar2000.exe")

	// Parse command line flags
	flag.Parse()

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterFoobarServiceServer(s, &server{
		foobarPath: *foobarPath,
	})

	// Serve the gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
