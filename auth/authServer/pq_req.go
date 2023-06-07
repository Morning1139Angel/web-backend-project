package authserver

import (
	"context"
	"strconv"
	"time"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	utils "github.com/Morning1139Angel/web-hw1/auth/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) PqReq(
	ctx context.Context,
	in *pb.PQRequest,
) (*pb.PQResponse, error) {

	fieldsHasIssues, err := checkPQFields(in)
	if fieldsHasIssues {
		return nil, err
	}

	messageId, _ := utils.GenerateRandomOddNumber()
	nonceServer := utils.GenerateNonce(20)
	clientNonce := in.Nonce
	p, g := utils.GetPandG()

	key := utils.StorageKeyFromNonces(clientNonce, nonceServer)
	s.rdb.SetEX(s.ctx, key, strconv.FormatUint(in.MessageId, 10), 20*time.Minute)

	completeNonces := &pb.CompleteNonces{Nonce: in.Nonce, NonceServer: nonceServer}
	return &pb.PQResponse{
		MessageId: messageId,
		Nonces:    completeNonces,
		P:         p,
		G:         g}, nil
}

func checkPQFields(in *pb.PQRequest) (bool, error) {
	// Check for missing required fields
	if in.Nonce == "" || in.MessageId == 0 {
		return true, status.Errorf(codes.InvalidArgument, "Missing required field(s)")
	}
	// Validate nonce length
	if len(in.Nonce) != 20 {
		return true, status.Errorf(codes.InvalidArgument, "Invalid nonce length")
	}
	// Validate messageId as an odd number
	if in.MessageId%2 != 1 {
		return true, status.Errorf(codes.InvalidArgument, "messageId must be an odd number")
	}
	return false, nil
}
