package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"foobar/proto"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewFoobarServiceClient(conn)

	// Send a request to trigger a rescan.
	req := &proto.RescanRequest{FoobarPath: "/path/to/foobar"}
	res, err := c.TriggerRescan(context.Background(), req)
	if err != nil {
		log.Fatalf("could not trigger rescan: %v", err)
	}
	log.Printf("rescan triggered: %v", res.Success)
}
