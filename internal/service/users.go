package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// Login users login.
func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return &v1.LoginReply{
		User: &v1.User{
			Username: "gin",
		},
	}, nil
}

// Register users registration.
func (s *RealWorldService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.RegisterReply, error) {
	_, err := s.uuc.CreateUser(ctx, &biz.User{Name: "tinggwan", Age: 18})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		User: &v1.User{
			Username: "tinggwan",
		},
	}, nil
}
