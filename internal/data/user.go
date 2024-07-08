package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realwd/internal/biz"
)

// userRepo 操作用户相关代码
type userRepo struct {
	data *Data
	log  *log.Helper
}

// biz里 userRepo 方法 的实现
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 操作数据库
func (u *userRepo) CreateUser(ctx context.Context, User *biz.User) error {
	//tx := u.data.Db.Create(&U)
	return nil
}

func (u *userRepo) SearchUser(ctx context.Context, User *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) UpdateUser(ctx context.Context, User *biz.User) error {
	//TODO implement me
	panic("implement me")
}
