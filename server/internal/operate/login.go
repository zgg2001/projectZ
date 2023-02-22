package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) UserLogin(ctx context.Context, request *rpc.UserLoginRequest) (*rpc.UserLoginResponse, error) {
	return &rpc.UserLoginResponse{Result: 0}, nil
}
