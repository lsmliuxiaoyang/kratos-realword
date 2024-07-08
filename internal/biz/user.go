package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// 领域对象
type User struct {
	UserName string
}

// biz层的 repo 是一个接口，定义各种方法，具体的实现在data层
// UserRepo 用户相关的接口
type UserRepo interface {
	SearchUser(ctx context.Context, User *User) error
	UpdateUser(ctx context.Context, User *User) error
	CreateUser(ctx context.Context, user *User) error
}

// xx相关的接口
type ProfileRepo interface {
	SearchProfile(ctx context.Context, User *User) error
}

// UserUseCase user 实现哪些Repo，repo 可能是功能的合集，u 、p 是对合集的分组？
type UserUseCase struct {
	uc  UserRepo
	pr  ProfileRepo
	log *log.Helper
}

// NewUserUseCase  函数是工厂函数，它创建了一个新的 UserUseCase 实例，并将传入的 UserRepo 和 ProfileRepo 实例分配给新实例的成员变量。
//
//	UserUseCase 实例就可以通过自己的成员变量来访问和操作数据存储。
func NewUserUseCase(u UserRepo, p ProfileRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		uc:  u,
		pr:  p,
		log: log.NewHelper(logger),
	}
}

// Register  对外提供接口调用data层createuser 代码
func (u *UserUseCase) Register(ctx context.Context, user *User) error {
	if err := u.uc.CreateUser(ctx, user); err != nil {

	}
	return nil
}
func (u *UserUseCase) Profile(ctx context.Context, user *User) error {
	return nil
}

func NewProfileUseCase() {

}
