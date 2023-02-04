package operate

import (
	"context"
	"log"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) LicencePlateCheck(con context.Context, request *rpc.LPCheckRequest) (*rpc.LPCheckResponse, error) {
	log.Println(request.Model, request.ParkingId, request.License)
	return &rpc.LPCheckResponse{Result: true, Balance: 100.01}, nil
}
