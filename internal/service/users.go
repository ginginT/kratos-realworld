package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

// Login users login.
func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{
		User: &v1.User{
			Username: "gin",
		},
	}, nil
}

// Users registration.
func (s *RealWorldService) Users(ctx context.Context, in *v1.UsersRequest) (*v1.UsersReply, error) {
	return &v1.UsersReply{
		User: &v1.User{
			Username: "tinggwan",
		},
	}, nil
}
