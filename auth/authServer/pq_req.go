package authserver

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	utils "github.com/Morning1139Angel/web-hw1/auth/utils"
)

type pqData struct {
	MessageId uint64
	P         uint64
	G         uint64
}

func (s *authServer) PqReq(
	ctx context.Context,
	in *pb.PQRequest,
) (*pb.PQResponse, error) {

	var messageId uint64 = 0
	nonceServer := utils.GenerateNonce(20)

	p, g := utils.GeneratePandG()

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
	key := utils.GetStringSha1(nonce + nonceServer)
	s.rdb.SetEX(s.ctx, key, encodingData, exp)
}
