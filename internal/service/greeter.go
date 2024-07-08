package service

import (
	"context"
	v1 "kratos-realwd/api/realwd/v1"
	"kratos-realwd/internal/biz"
)

// SayHello implements helloworld.RealwdServer.
func (r *RealwdService) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := r.uc.CreateRealwd(ctx, &biz.Realwd{Hello: req.Name})
	return &v1.HelloReply{
		Message: "Hello " + g.Hello,
	}, err
}
