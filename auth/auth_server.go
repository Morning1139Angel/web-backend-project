package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	"github.com/go-redis/redis/v8"
)

type pqData struct {
	MessageId uint64
	P         uint64
	G         uint64
}

type authServer struct {
	pb.UnimplementedAuthServiceServer
	rdb *redis.Client
	ctx context.Context
}

func NewAuthServer(rdb *redis.Client, ctx context.Context) *authServer {
	return &authServer{rdb: rdb, ctx: ctx}
}

func (s *authServer) PqReq(
	ctx context.Context,
	in *pb.PQRequest,
) (*pb.PQResponse, error) {

	var messageId uint64 = 0
	nonceServer := generateNonce(20)

	p, g := generatePandG()

	pqData := pqData{in.MessageId, p, g}
	s.storePQdata(pqData, in.Nonce, nonceServer, 20*time.Minute)

	completeNonces := &pb.CompleteNonces{Nonce: in.Nonce, NonceServer: nonceServer}
	return &pb.PQResponse{MessageId: messageId, Nonces: completeNonces}, nil
}

func (s *authServer) storePQdata(pqData pqData, nonce, nonceServer string, exp time.Duration) {
	encodingData, err := json.Marshal(pqData)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}
	key := getStringSha1(nonce + nonceServer)
	s.rdb.SetEX(s.ctx, key, encodingData, exp)
}

func (s *authServer) Req_DHParams(
	ctx context.Context,
	in *pb.DHParamsRequest,
) (*pb.DHParamsResponse, error) {

	// function body
	var messageId uint64 = 0
	var b uint64 = 0

	return &pb.DHParamsResponse{MessageId: messageId, Nonces: in.Nonces, B: b}, nil
}
