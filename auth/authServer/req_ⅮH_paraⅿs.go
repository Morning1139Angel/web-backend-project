package authserver

import (
	"context"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
)

func (s *authServer) Req_DHParams(
	ctx context.Context,
	in *pb.DHParamsRequest,
) (*pb.DHParamsResponse, error) {

	// function body
	var messageId uint64 = 0
	var b uint64 = 0

	return &pb.DHParamsResponse{MessageId: messageId, Nonces: in.Nonces, B: b}, nil
}
