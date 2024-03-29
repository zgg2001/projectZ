package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/user_server/internal/data"
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

func (ss *serverService) AdminLogin(ctx context.Context, request *rpc.AdminLoginRequest) (*rpc.AdminLoginResponse, error) {
	pid := request.GetPId()
	password := request.GetPassword()
	count, ret := data.ParkingLoginAuth(pid, password)
	return &rpc.AdminLoginResponse{Result: ret, Count: count}, nil
}

func (ss *serverService) AdminGetSpaceInfo(ctx context.Context, request *rpc.AdminGetSpaceInfoRequest) (*rpc.AdminGetSpaceInfoResponse, error) {
	pid := request.GetPId()
	sid := request.GetSId()
	isUse, license, entryTime := data.ParkingGetSpaceInfo(pid, sid)
	return &rpc.AdminGetSpaceInfoResponse{IsUse: isUse, License: license, Entrytime: entryTime}, nil
}
