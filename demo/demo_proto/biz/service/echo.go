package service

import (
	"context"

	pbapi "github.com/99n/gomall/demo/demo_proto/kitex_gen/pbapi"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Responce, err error) {
	// Finish your business logic.

	return &pbapi.Responce{Message: req.Message}, nil
}
