package authserver

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	"github.com/Morning1139Angel/web-hw1/auth/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) Req_DHParams(
	ctx context.Context,
	in *pb.DHParamsRequest,
) (*pb.DHParamsResponse, error) {

	fieldsHasIssues, err := s.checkDHFields(in)
	if fieldsHasIssues {
		return nil, err
	}

	messageId, _ := utils.GenerateRandomOddNumber()
	clientNonce, nonceServer := in.Nonces.Nonce, in.Nonces.NonceServer
	key := utils.StorageKeyFromNonces(clientNonce, nonceServer)

	// Define the prime number p and generator g
	p, g, _ := utils.GetPandGBigInts()

	B, commonSecret := diffieHellman(in.A, p, g)

	s.rdb.Set(s.ctx, key, commonSecret, 0)

	return &pb.DHParamsResponse{MessageId: messageId, Nonces: in.Nonces, B: B}, nil
}

func diffieHellman(AString string, p, g *big.Int) (string, string) {
	A, _ := new(big.Int).SetString(AString, 10)

	// Generate private key for server
	b, _ := rand.Int(rand.Reader, p)

	// Compute public key
	B := new(big.Int).Exp(g, b, p)

	// Perform Diffie-Hellman key exchange
	commonSecret := new(big.Int).Exp(A, b, p)
	return B.String(), commonSecret.String()
}

func (s *authServer) checkDHFields(in *pb.DHParamsRequest) (bool, error) {
	if err := validateRequiredFields(in); err != nil {
		return true, err
	}

	if err := validateNonceLengths(in.Nonces.Nonce, in.Nonces.NonceServer); err != nil {
		return true, err
	}

	if err := validateMessageId(in.MessageId); err != nil {
		return true, err
	}

	key := utils.GetStringSha1(in.Nonces.Nonce + in.Nonces.NonceServer)
	storedMessageIdStr, err := s.rdb.Get(s.ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return true, status.Errorf(codes.FailedPrecondition, "pq request not yet sent")
		} else {
			return true, status.Errorf(codes.Internal, "Failed to get value from Redis: %v", err)
		}
	}

	storedMessageIdInt, _ := strconv.ParseUint(storedMessageIdStr, 10, 64)
	if in.MessageId < storedMessageIdInt {
		return true, status.Errorf(codes.InvalidArgument, "messageId must be larger than the number specified in the pq request")
	}

	return false, nil
}

func validateRequiredFields(in *pb.DHParamsRequest) error {
	if in.Nonces == nil || in.MessageId == 0 || in.A == "" {
		return status.Errorf(codes.InvalidArgument, "Missing required field(s)")
	}
	return nil
}

func validateNonceLengths(clientNonce, serverNonce string) error {
	if len(clientNonce) != 20 || len(serverNonce) != 20 {
		return status.Errorf(codes.InvalidArgument, "Invalid nonce lengths")
	}
	return nil
}

func validateMessageId(messageId uint64) error {
	if messageId%2 != 1 {
		return status.Errorf(codes.InvalidArgument, "messageId must be an odd number")
	}
	return nil
}
