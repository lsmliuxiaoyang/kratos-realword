package service

import (
	"context"

	pb "kratos-demo/api/realWord/v1"
)

type Greeter1Service struct {
	pb.UnimplementedGreeter1Server
}

func NewGreeter1Service() *Greeter1Service {
	return &Greeter1Service{}
}

func (s *Greeter1Service) CreateGreeter1(ctx context.Context, req *pb.CreateGreeter1Request) (*pb.CreateGreeter1Reply, error) {
	return &pb.CreateGreeter1Reply{}, nil
}
func (s *Greeter1Service) UpdateGreeter1(ctx context.Context, req *pb.UpdateGreeter1Request) (*pb.UpdateGreeter1Reply, error) {
	return &pb.UpdateGreeter1Reply{}, nil
}
func (s *Greeter1Service) DeleteGreeter1(ctx context.Context, req *pb.DeleteGreeter1Request) (*pb.DeleteGreeter1Reply, error) {
	return &pb.DeleteGreeter1Reply{}, nil
}
func (s *Greeter1Service) GetGreeter1(ctx context.Context, req *pb.GetGreeter1Request) (*pb.GetGreeter1Reply, error) {
	return &pb.GetGreeter1Reply{}, nil
}
func (s *Greeter1Service) ListGreeter1(ctx context.Context, req *pb.ListGreeter1Request) (*pb.ListGreeter1Reply, error) {
	return &pb.ListGreeter1Reply{}, nil
}
