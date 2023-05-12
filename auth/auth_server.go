package main

import (
	"context"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	"github.com/go-redis/redis/v8"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
	rdb *redis.Client
}

func (s *authServer) PqReq(
	ctx context.Context,
	in *pb.PQRequest,
) (*pb.PQResponse, error) {

	// function body
	var messageId uint64 = 0
	nonceServer := ""

	completeNonces := &pb.CompleteNonces{Nonce: in.Nonce, NonceServer: nonceServer}
	return &pb.PQResponse{MessageId: messageId, Nonces: completeNonces}, nil
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

/*
service AuthService{
    rpc pq_req(PQRequest) returns (PQResponse);
    rpc req_DH_params(DHParamsRequest) returns (DHParamsResponse);
}

message CompleteNonces {
    string nonce = 1;
    string nonce_server = 2;
}

message PQRequest {
    uint64 message_id = 1;
    string nonce = 2;
}

message PQResponse {
    uint64 message_id = 1;
    CompleteNonces nonces = 2;
    uint64 p = 3;
    uint64 g = 4;
}

message DHParamsRequest {
    uint64 message_id = 1;
    CompleteNonces nonces = 2;
}

message DHParamsResponse {
    uint64 message_id = 1;
    CompleteNonces nonces = 2;
    uint64 b = 3;
}

*/
