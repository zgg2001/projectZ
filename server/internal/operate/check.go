package operate

import (
	"context"
	"log"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

const (
	FrontCamera int32 = 0
	RearCamera  int32 = 1
)

func (ss *serverService) LicencePlateCheck(con context.Context, request *rpc.LPCheckRequest) (*rpc.LPCheckResponse, error) {
	log.Println(request.Model, request.ParkingId, request.ParkingSpaceId, request.License)

	model := request.GetModel()
	license := request.GetLicense()
	pid := request.GetParkingId()
	sid := request.GetParkingSpaceId()

	// find car
	ok, uptr := ss.uMgr.GetUser(license)
	if !ok {
		return &rpc.LPCheckResponse{Result: false, Balance: 1}, nil
	}
	cptr, err := uptr.GetCarPtr(license)
	if err != nil {
		return &rpc.LPCheckResponse{Result: false, Balance: 0}, err
	}
	balance := uptr.GetBalance()

	if model == FrontCamera {
		if balance <= 0 {
			return &rpc.LPCheckResponse{Result: false, Balance: balance}, nil
		}
		pptr, sptr, err := ss.pMgr.MgrGetParkingPtrPair(pid, sid)
		if err != nil {
			return &rpc.LPCheckResponse{Result: false, Balance: 0}, err
		}
		cptr.SetParkingSpace(pptr, sptr, 0)
		// 登记entry time
	} else {

	}

	return &rpc.LPCheckResponse{Result: true, Balance: balance}, nil
}
