package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/user_server/internal/data"
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

func (ss *serverService) GetUserData(ctx context.Context, request *rpc.GetUserDataRequest) (*rpc.GetUserDataResponse, error) {
	var ret []*rpc.CarInfo
	uid := request.GetUId()
	licenseArr := data.UserGetLicenseArr(uid)
	for _, license := range licenseArr {
		if info := data.RedisGetLicenseInfo(license); info != nil {
			ret = append(ret, info)
		}
	}
	return &rpc.GetUserDataResponse{CarInfoArr: ret}, nil
}
