package main

import (
	"log"
	"os"

	_ "github.com/Morning1139Angel/web-hw1/gateway/docs"
	pb "github.com/Morning1139Angel/web-hw1/gateway/grpc"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var authServerHost = os.Getenv("AUTH_SERVER_HOST")
var authServerPort = os.Getenv("AUTH_SERVER_PORT")
var authClient pb.AuthServiceClient

// @title           Web Hw1
// @version         1.0
// @description     A sample homework service API in Go using Gin framework.
// @termsOfService  https://tos.santoshk.dev

// @contact.name   amir khazama
// @contact.email  amirkhazama1139@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:80
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

	//add swagger
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authGroup := engine.Group("/auth")

	rds := InitRedicClient()

	// Apply IP block middleware to the /auth route
	authGroup.Use(IPBlockMiddleware(rds))

	authGroup.Use(RequestCountMiddleware(rds, 100))

	authGroup.POST("/pq", PQhandler)
	authGroup.POST("/dh", DHhandler)

	engine.Run(":8080")
}

func checkMessageIdOddness(messageId uint64) error {
	if messageId%2 != 1 {
		return status.Errorf(codes.InvalidArgument, "messageId must be an odd number")
	}
	return nil
}
