package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-realwd/api/realwd/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// 一种将数据访问逻辑从业务逻辑中分离出来的设计模式
type ArticleRepo interface {
	CreateArticle(ctx context.Context) error
	SearchArticle(ctx context.Context) error
}

type CommentRepo interface {
	CreateComment(ctx context.Context) error
	SearchComment(ctx context.Context) error
}

type TagRepo interface {
	CreateTag(ctx context.Context) error
	SearchTag(ctx context.Context) error
}

// SocialUsecase 应用服务层或业务逻辑层
type SocialUsecase struct {
	at  ArticleRepo
	co  CommentRepo
	ta  TagRepo
	log *log.Helper
}

// NewSocialUsecase new a Realwd usecase.
func NewSocialUsecase(at ArticleRepo, co CommentRepo, tag TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{
		at:  at,
		co:  co,
		ta:  tag,
		log: log.NewHelper(logger),
	}
}

// Usecase 的方法
func (s *SocialUsecase) CreateArticle(ctx context.Context) error {
	return nil
}

func (s *SocialUsecase) SearchArticle(ctx context.Context) error {
	return nil
}

func (s *SocialUsecase) CreateComment(ctx context.Context) error {
	return nil
}
func (s *SocialUsecase) SearchComment(ctx context.Context) error {
	return nil
}

func (s *SocialUsecase) CreateTag(ctx context.Context) error {
	return nil
}
func (s *SocialUsecase) SearchTag(ctx context.Context) error {
	return nil
}
