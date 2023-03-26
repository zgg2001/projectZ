package operate

import (
	"context"
	"time"

	"github.com/zgg2001/projectZ/server/parking_server/internal/data"
	"github.com/zgg2001/projectZ/server/parking_server/pkg/rpc"
)

const (
	FrontCamera int32 = 0
	RearCamera  int32 = 1
)

func (ss *serverService) LicencePlateCheck(con context.Context, request *rpc.LPCheckRequest) (*rpc.LPCheckResponse, error) {

	model := request.GetModel()
	license := request.GetLicense()
	pid := request.GetParkingId()
	sid := request.GetParkingSpaceId()

	// find car
	ok, uptr := ss.uMgr.GetUserByLicense(license)
	if !ok {
		// 未登记车辆
		return &rpc.LPCheckResponse{Result: false, Balance: 1}, nil
	}
	cptr, err := uptr.GetCarPtrCheckEntered(license)
	if err != nil {
		// 车辆登记数据错误
		return &rpc.LPCheckResponse{Result: false, Balance: 0}, err
	}
	balance := uptr.GetBalance()
	requestTime := time.Now().Unix()

	if model == FrontCamera {
		if balance <= 0 {
			// 余额不足
			return &rpc.LPCheckResponse{Result: false, Balance: balance}, nil
		}
		pptr, sptr, err := ss.pMgr.MgrGetParkingPtrPair(pid, sid)
		if err != nil {
			// 请求id错误
			return &rpc.LPCheckResponse{Result: false, Balance: 0}, err
		}
		cptr.SetParkingSpace(pptr, sptr, requestTime)
		data.InsertRecordTbl(license, pid, sid, requestTime)
		data.InsertParkingRecordTbl(license, pid, sid, FrontCamera, requestTime)
	} else {
		// 消费校验
		etime, err := data.SelectRecordTbl(license)
		if err != nil {
			// 无进入记录
			return &rpc.LPCheckResponse{Result: false, Balance: 0}, err
		}
		balance = calculateBalance(balance, etime, requestTime)
		uptr.SetBalance(balance)
		if balance <= -50 {
			// 余额不足
			return &rpc.LPCheckResponse{Result: false, Balance: balance}, nil
		}
		cptr.SetParkingSpace(nil, nil, 0)
		data.DeleteRecordTbl(license)
		data.InsertParkingRecordTbl(license, pid, sid, RearCamera, requestTime)
	}

	return &rpc.LPCheckResponse{Result: true, Balance: balance}, nil
}

func calculateBalance(balance int32, entryTime, requestTime int64) int32 {
	parkingTime := requestTime - entryTime
	return balance - int32(parkingTime/60)
}
