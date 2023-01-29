package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) UploadParkingInfo(context.Context, *rpc.UploadInfoRequest) (*rpc.UploadInfoResponse, error) {
	return &rpc.UploadInfoResponse{Result: 1}, nil
}
