package operate

import (
	"context"
	"log"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) GetUserData(ctx context.Context, request *rpc.GetUserDataRequest) (*rpc.GetUserDataResponse, error) {

	log.Println(request)

	var ret []*rpc.CarInfo
	uid := request.GetUId()

	uptr, err := ss.uMgr.GetUserById(uid)
	if err != nil {
		// 未登记用户
		return &rpc.GetUserDataResponse{CarInfoArr: ret}, err
	}

	cptrArr := uptr.GetCarPtrArr()
	for _, cptr := range cptrArr {

		ret = append(ret, cptr.GetCarPtrArr())
	}

	return &rpc.GetUserDataResponse{CarInfoArr: ret}, nil
}
