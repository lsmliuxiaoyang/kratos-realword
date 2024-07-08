package biz

import (
	"github.com/google/wire"
	"kratos-realwd/internal/biz/user"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(user.NewUserUseCase, NewSocialUsecase)
