package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) LicencePlateCheck(context.Context, *rpc.LPCheckRequest) (*rpc.LPCheckResponse, error) {
	return &rpc.LPCheckResponse{Result: true, Balance: 100.01}, nil
}
