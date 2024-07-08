package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realwd/internal/biz"
)

// repo 里面连接数据库
type articleRepo struct {
	data *Data
	log  *log.Helper
}

func (a articleRepo) CreateArticle(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a articleRepo) SearchArticle(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func (c commentRepo) CreateComment(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c commentRepo) SearchComment(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func (t tagRepo) CreateTag(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (t tagRepo) SearchTag(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
