package service

import (
	"context"
	"github.com/google/wire"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a RealWorld service.
type RealWorldService struct {
	v1.UnimplementedRealworldServer

	uc *biz.RealWorldUsecase
}

// NewRealWorldService new a RealWorld service.
func NewRealWorldService(uc *biz.RealWorldUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}

// SayHello implements realworld.RealWorldServer.
func (s *RealWorldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateRealWorld(ctx, &biz.RealWorld{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "RealWorld " + g.Hello}, nil
}
