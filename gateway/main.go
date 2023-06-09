package main

import (
	"log"
	"os"

	pb "github.com/Morning1139Angel/web-hw1/gateway/grpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var authServerHost = os.Getenv("AUTH_SERVER_HOST")
var authServerPort = os.Getenv("AUTH_SERVER_PORT")
var authClient pb.AuthServiceClient

func main() {
	// Create a gRPC connection to the server
	targetPath := authServerHost + ":" + authServerPort
	conn, err := grpc.Dial(targetPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client using the connection
	authClient = pb.NewAuthServiceClient(conn)

	//start the gin server
	engine := gin.New()
	engine.POST("/auth/pq", PQhandler)
	engine.POST("/auth/dh", DHhandler)
	engine.Run(":8080")
}

func checkMessageIdOddness(messageId uint64) error {
	if messageId%2 != 1 {
		return status.Errorf(codes.InvalidArgument, "messageId must be an odd number")
	}
	return nil
}
