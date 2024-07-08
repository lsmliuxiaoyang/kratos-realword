package service

import (
	"github.com/google/wire"
	v1 "kratos-realwd/api/realwd/v1"
	"kratos-realwd/internal/biz/user/user"
)

// ProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewRealwdService,
)

// RealwdService is a Realwd service.
type RealwdService struct {
	v1.UnimplementedReadwdServer

	uc *user.UserUseCase
	//sc *biz.SocialUsecase
}

// NewRealwdService new a Realwd service.
func NewRealwdService(uc *user.UserUseCase) *RealwdService {
	return &RealwdService{uc: uc}
}
