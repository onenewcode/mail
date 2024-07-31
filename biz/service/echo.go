package service

import (
	"context"

	pbapi "mail/kitex_gen/pbapi"
)

type EchoService struct {
	ctx context.Context
}

// NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// 执行命令create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Response, err error) {

	return &pbapi.Response{Message: req.Message}, nil
}
