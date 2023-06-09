package main

import (
	"context"
	"net/http"

	pb "github.com/Morning1139Angel/web-hw1/gateway/grpc"
	"github.com/gin-gonic/gin"
)

type DHRequest struct {
	A         string         `json:"A" binding:"required,number"`
	MessageId uint64         `json:"messageId" binding:"required"`
	Nonces    CompleteNonces `json:"nonces" binding:"required"`
}

type CompleteNonces struct {
	Nonce        string `json:"nonce" binding:"required,len=20"`
	Nonce_server string `json:"nonceServer" binding:"required,len=20"`
}

func DHhandler(c *gin.Context) {
	//do field checks on the request body
	body, err := fieldCheckDHBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//send request to server
	completeNonces := &pb.CompleteNonces{Nonce: body.Nonces.Nonce, NonceServer: body.Nonces.Nonce_server}
	request := &pb.DHParamsRequest{MessageId: body.MessageId, Nonces: completeNonces, A: body.A}
	dhResponse, err := authClient.Req_DHParams(context.Background(), request)

	//send the gRPC response back to the client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dhResponse)

}

func fieldCheckDHBody(c *gin.Context) (DHRequest, error) {
	body := DHRequest{}
	//do the checks specified by the struct tags of PQRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		return DHRequest{}, err
	}

	//check message Id being odd
	if err := checkMessageIdOddness(body.MessageId); err != nil {
		return DHRequest{}, err
	}
	return body, nil
}
