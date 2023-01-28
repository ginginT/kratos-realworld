package service

import (
	"github.com/google/wire"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a RealWorld service.
type RealWorldService struct {
	v1.UnimplementedRealworldServer

	uuc *biz.UserUsecase
}

// NewRealWorldService new a RealWorld service.
func NewRealWorldService(
	uuc *biz.UserUsecase,
) *RealWorldService {
	return &RealWorldService{
		uuc: uuc,
	}
}
