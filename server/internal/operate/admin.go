package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) AdminLogin(ctx context.Context, request *rpc.AdminLoginRequest) (*rpc.AdminLoginResponse, error) {
	pid := request.GetPId()
	password := request.GetPassword()
	ret := ss.pMgr.LoginAuth(pid, password)
	return &rpc.AdminLoginResponse{Result: ret}, nil
}
