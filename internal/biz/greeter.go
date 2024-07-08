package biz

import (
	"context"

	v1 "kratos-realwd/api/realwd/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Realwd is a Realwd model.
type Realwd struct {
	Hello string
}

// RealwdRepo is a Greater repo.
type RealwdRepo interface {
	Save(context.Context, *Realwd) (*Realwd, error)
	Update(context.Context, *Realwd) (*Realwd, error)
	FindByID(context.Context, int64) (*Realwd, error)
	ListByHello(context.Context, string) ([]*Realwd, error)
	ListAll(context.Context) ([]*Realwd, error)
}

// RealwdUsecase is a Realwd usecase.
type RealwdUsecase struct {
	repo RealwdRepo
	log  *log.Helper
}

// NewRealwdUsecase new a Realwd usecase.
func NewRealwdUsecase(repo RealwdRepo, logger log.Logger) *RealwdUsecase {
	return &RealwdUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateRealwd creates a Realwd, and returns the new Realwd.
func (uc *RealwdUsecase) CreateRealwd(ctx context.Context, g *Realwd) (*Realwd, error) {
	uc.log.WithContext(ctx).Infof("CreateRealwd: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
