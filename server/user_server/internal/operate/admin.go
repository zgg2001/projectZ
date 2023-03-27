package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/user_server/internal/data"
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

func (ss *serverService) AdminLogin(ctx context.Context, request *rpc.AdminLoginRequest) (*rpc.AdminLoginResponse, error) {
	pid := request.GetPId()
	password := request.GetPassword()
	ret := data.ParkingLoginAuth(pid, password)
	return &rpc.AdminLoginResponse{Result: ret}, nil
}
