package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/Morning1139Angel/web-hw1/biz/myPkgName"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = map[string]string{
	"host":     os.Getenv("POSTGRES_HOST"),
	"user":     os.Getenv("POSTGRES_USER"),
	"password": os.Getenv("POSTGRES_PASSWORD"),
	"dbname":   os.Getenv("POSTGRES_DBNAME"),
	"port":     os.Getenv("POSTGRES_PORT"),
	"sslmode":  os.Getenv("POSTGRES_SSLMODE"),
	"TimeZone": os.Getenv("POSTGRES_TIMEZONE"),
}

func main() {
	//TODO: change this and create database first
	DB, err := gorm.Open(postgres.Open(mapToString(dsn)), &gorm.Config{})
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

func mapToString(m map[string]string) string {
	var parts []string
	for key, value := range m {
		parts = append(parts, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(parts, " ")
}
