package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

var UploadParkingInfoService = &uploadParkingInfoService{}

type uploadParkingInfoService struct {
	rpc.UnimplementedProjectServiceServer
}

func (us *uploadParkingInfoService) UploadParkingInfo(context.Context, *rpc.UploadInfoRequest) (*rpc.UploadInfoResponse, error) {
	return &rpc.UploadInfoResponse{Result: 1}, nil
}
