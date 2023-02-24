package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) UserLogin(ctx context.Context, request *rpc.UserLoginRequest) (*rpc.UserLoginResponse, error) {
	username := request.GetUsername()
	password := request.GetPassword()
	err := ss.uMgr.LoginAuth(username, password)
	if err != nil {
		return &rpc.UserLoginResponse{Result: 0}, err
	}
	return &rpc.UserLoginResponse{Result: 1}, nil
}
