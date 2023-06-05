package main

import (
	"context"
	"github.com/Morning1139Angel/web-hw1/biz/myPkgName"
	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
	myPkgName.UnimplementedGetUsersServiceServer
}

func (s *server) GetUser(ctx context.Context, req *myPkgName.GetUserRequest) (*myPkgName.GetUserResponse, error) {
	// Authenticate the user using the auth_key if necessary
	// ...

	// Retrieve the user from the database using GORM
	var users []*myPkgName.User
	s.db.Find(&users, "id = ?", req.UserId)

	// Construct the response
	resp := &myPkgName.GetUserResponse{
		Users:     users,
		MessageId: req.MessageId,
	}

	return resp, nil
}

func (s *server) GetUserWithInject(ctx context.Context, req *myPkgName.GetUserWithInjectRequest) (*myPkgName.GetUserResponse, error) {
	var users []*myPkgName.User
	s.db.Raw("SELECT * from USERS where id ="+req.UserId, users)

	// Construct the response
	resp := &myPkgName.GetUserResponse{
		Users:     users,
		MessageId: req.MessageId,
	}

	return resp, nil
}
