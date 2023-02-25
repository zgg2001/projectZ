package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) UserLogin(ctx context.Context, request *rpc.UserLoginRequest) (*rpc.UserLoginResponse, error) {
	username := request.GetUsername()
	password := request.GetPassword()
	ret := ss.uMgr.LoginAuth(username, password)
	return &rpc.UserLoginResponse{Result: ret, UId: 1}, nil
}
