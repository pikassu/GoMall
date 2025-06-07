package service
//存放业务逻辑
import (
	"context"
	api "github.com/pikassu/Gomall/gomall/demo/demo_thrift/kitex_gen/api"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *api.Requset) (resp *api.Response, err error) {
	// Finish your business logic.

	return &api.Response{Message: req.Message},nil
}
