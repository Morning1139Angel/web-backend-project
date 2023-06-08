package main

import (
	"context"
	"log"
	"net/http"
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

type PQRequest struct {
	MessageId uint64 `json:"messageId" binding:"required"` //TODO:add custom odd validator
	Nonce     string `json:"nonce" binding:"required,len=20"`
}

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
	engine.POST("/auth/pq", PQhandler) //TODO: add DH handler
	engine.Run(":8080")
}

func PQhandler(c *gin.Context) {
	//do field checks on the request body
	body, err := fieldCheckPQBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//send request to server
	request := pb.PQRequest{MessageId: body.MessageId, Nonce: body.Nonce}
	pqResponse, err := authClient.PqReq(context.Background(), &request)

	//send the gRPC response back to the client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pqResponse)

}

func fieldCheckPQBody(c *gin.Context) (PQRequest, error) {
	body := PQRequest{}
	//do the checks specified by the struct tags of PQRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		return PQRequest{}, err
	}

	//check message Id being odd
	if err := checkMessageIdOddness(body.MessageId); err != nil {
		return PQRequest{}, err
	}
	return body, nil
}

func checkMessageIdOddness(messageId uint64) error {
	if messageId%2 != 1 {
		return status.Errorf(codes.InvalidArgument, "messageId must be an odd number")
	}
	return nil
}
