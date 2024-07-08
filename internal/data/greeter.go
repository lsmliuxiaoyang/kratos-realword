package data

import (
	"context"

	"kratos-realwd/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type RealwdRepo struct {
	data *Data
	log  *log.Helper
}

// NewRealwdRepo .
func NewRealwdRepo(data *Data, logger log.Logger) biz.RealwdRepo {
	return &RealwdRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RealwdRepo) Save(ctx context.Context, g *biz.Realwd) (*biz.Realwd, error) {
	return g, nil
}

func (r *RealwdRepo) Update(ctx context.Context, g *biz.Realwd) (*biz.Realwd, error) {
	return g, nil
}

func (r *RealwdRepo) FindByID(context.Context, int64) (*biz.Realwd, error) {
	return nil, nil
}

func (r *RealwdRepo) ListByHello(context.Context, string) ([]*biz.Realwd, error) {
	return nil, nil
}

func (r *RealwdRepo) ListAll(context.Context) ([]*biz.Realwd, error) {
	return nil, nil
}
