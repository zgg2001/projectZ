package operate

import (
	"context"
	"log"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) UploadParkingInfo(con context.Context, request *rpc.UploadInfoRequest) (*rpc.UploadInfoResponse, error) {
	log.Println(request)
	return &rpc.UploadInfoResponse{Result: 1}, nil
}
