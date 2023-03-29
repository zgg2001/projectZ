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
	ok, uid := data.GetUserByLicense(license)
	if !ok {
		// 未登记车辆
		return &rpc.LPCheckResponse{Result: false, Balance: 1}, nil
	}
	balance := data.GetBalanceByUid(uid)
	requestTime := time.Now().Unix()

	if model == FrontCamera {
		err := data.CheckCarIsEntered(license)
		if err != nil {
			// 车辆登记数据错误
			return &rpc.LPCheckResponse{Result: false, Balance: 0}, err
		}
		if balance <= 0 {
			// 余额不足
			return &rpc.LPCheckResponse{Result: false, Balance: balance}, nil
		}
		ss.SetParkingSpace(license, pid, sid, requestTime, FrontCamera)
	} else {
		// 消费校验
		etime, err := data.SelectRecordTbl(license)
		if err != nil {
			// 无进入记录
			return &rpc.LPCheckResponse{Result: false, Balance: 0}, err
		}
		balance = calculateBalance(balance, etime, requestTime)
		if balance <= -50 {
			// 余额不足
			return &rpc.LPCheckResponse{Result: false, Balance: balance}, nil
		}
		ss.SetBalance(uid, balance)
		ss.SetParkingSpace(license, -1, -1, 0, RearCamera)
	}

	return &rpc.LPCheckResponse{Result: true, Balance: balance}, nil
}

func calculateBalance(balance int32, entryTime, requestTime int64) int32 {
	parkingTime := requestTime - entryTime
	return balance - int32(parkingTime/60)
}
