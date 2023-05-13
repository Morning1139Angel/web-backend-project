package main

import (
	"context"
	"log"
	"net"

	authServer "github.com/Morning1139Angel/web-hw1/auth/authServer"
	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	utils "github.com/Morning1139Angel/web-hw1/auth/utils"
	"google.golang.org/grpc"
)

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()

	//initialize redis client
	rdb := utils.InitRedicClient()

	// Register implementation of the service
	authService := authServer.NewAuthServer(rdb, context.Background())
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
