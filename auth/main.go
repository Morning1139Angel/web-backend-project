package main

import (
	"log"
	"net"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	"google.golang.org/grpc"
)

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()

	// Register implementation of the service
	authService := &authServer{}
	pb.RegisterAuthServiceServer(server, authService)

	// Create a TCP listener
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start serving requests
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
