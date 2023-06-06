package authserver

import (
	"context"
	"strconv"
	"time"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	utils "github.com/Morning1139Angel/web-hw1/auth/utils"
)

func (s *authServer) PqReq(
	ctx context.Context,
	in *pb.PQRequest,
) (*pb.PQResponse, error) {

	messageId, _ := utils.GenerateRandomOddNumber()
	nonceServer := utils.GenerateNonce(20)
	clientNonce := in.Nonce

	key := utils.StorageKeyFromNonces(clientNonce, nonceServer)
	s.rdb.SetEX(s.ctx, key, strconv.FormatUint(in.MessageId, 10), 20*time.Minute)

	completeNonces := &pb.CompleteNonces{Nonce: in.Nonce, NonceServer: nonceServer}
	return &pb.PQResponse{
		MessageId: messageId,
		Nonces:    completeNonces,
		P:         "115792089237316195423570985008687907853269984665640564039457584007913129639747",
		G:         "2"}, nil
}
