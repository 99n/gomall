package service

import (
	"context"

	api "github.com/99n/gomall/demo/demo_thrift/kitex_gen/api"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *api.Request) (resp *api.Responce, err error) {
	// Finish your business logic.

	return &api.Responce{Message: req.Message}, nil
}
