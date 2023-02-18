package operate

import (
	"context"
	"log"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) GetUserData(ctx context.Context, request *rpc.GetUserDataRequest) (*rpc.GetUserDataResponse, error) {

	log.Println(request)

	var ret []*rpc.CarInfo

	return &rpc.GetUserDataResponse{CarInfoArr: ret}, nil
}
