package main

import (
	"github.com/Morning1139Angel/web-hw1/biz/myPkgName"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	//TODO: change this and create database first
	dsn := "host=localhost user=postgres password=test dbname=users port=5432 sslmode=disable TimeZone=Asia/Tehran"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create the gRPC server
	s := grpc.NewServer()
	myPkgName.RegisterGetUsersServiceServer(s, &server{db: DB})

	// Start listening on the specified port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start serving requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
