package main

import (
	"context"
	"net/http"

	pb "github.com/Morning1139Angel/web-hw1/gateway/grpc"
	"github.com/gin-gonic/gin"
)

type PQRequest struct {
	MessageId uint64 `json:"messageId" binding:"required"` //TODO:add custom odd validator
	Nonce     string `json:"nonce" binding:"required,len=20"`
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
