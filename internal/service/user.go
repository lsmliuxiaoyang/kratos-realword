package service

import (
	"context"
	v1 "kratos-realwd/api/realwd/v1"
)

func (s *RealwdService) Login(context.Context, *v1.LoginRequest) (*v1.LoginReply, error) {
	//req := v1.LoginRequest{}
	rsp := &v1.LoginReply{
		Token: "123456",
	}
	return rsp, nil

}
