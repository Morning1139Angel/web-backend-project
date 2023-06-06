package authserver

import (
	"context"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	"github.com/go-redis/redis/v8"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
	rdb *redis.Client
	ctx context.Context
}

func NewAuthServer(rdb *redis.Client, ctx context.Context) *authServer {
	return &authServer{rdb: rdb, ctx: ctx}
}

func (s *authServer) readFromRedis(key string) *redis.StringCmd {
	return s.rdb.Get(s.ctx, key)
}
