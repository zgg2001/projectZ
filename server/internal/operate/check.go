package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

var CheckService = &checkService{}

type checkService struct {
	rpc.UnimplementedProjectServiceServer
}

func (cs *checkService) LicencePlateCheck(context.Context, *rpc.LPCheckRequest) (*rpc.LPCheckResponse, error) {

	return &rpc.LPCheckResponse{Result: 1}, nil
}
