package service

import (
	"context"

	v1 "kratos-demo/api/realWord/v1"
	"kratos-demo/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedRealWordServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements realWord.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

func (s *GreeterService) SayHi(ctx context.Context, in *v1.HiRequest) (*v1.HiReplay, error) {
	g, err := s.uc.Test(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HiReplay{Message: "hi " + g.Hello}, nil
}
