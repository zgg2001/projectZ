package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) AdminLogin(ctx context.Context, request *rpc.AdminLoginRequest) (*rpc.AdminLoginResponse, error) {
	return &rpc.AdminLoginResponse{Result: rpc.LoginResult_LOGIN_SUCCESS}, nil
}
