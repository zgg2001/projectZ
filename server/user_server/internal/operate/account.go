package operate

import (
	"context"
	"time"

	"github.com/zgg2001/projectZ/server/user_server/internal/data"
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

func (ss *serverService) UserLogin(ctx context.Context, request *rpc.UserLoginRequest) (*rpc.UserLoginResponse, error) {
	username := request.GetUsername()
	password := request.GetPassword()
	uid, ret := data.LoginAuth(username, password)
	return &rpc.UserLoginResponse{Result: ret, UId: uid}, nil
}

func (ss *serverService) UserRegistration(ctx context.Context, request *rpc.UserRegistrationRequest) (*rpc.UserRegistrationResponse, error) {
	username := request.GetUsername()
	password := request.GetPassword()
	ret := data.UserRegistrationAuth(username)
	nowTime := time.Now().Unix()
	if ret == rpc.RegistrationResult_REGISTRATION_SUCCESS {
		ss.SqlRegisterUser(username, password, nowTime)
	}
	return &rpc.UserRegistrationResponse{Result: ret}, nil
}

func (ss *serverService) UserRecharge(ctx context.Context, request *rpc.UserRechargeRequest) (*rpc.UserRechargeResponse, error) {
	uid := request.GetUId()
	amount := request.GetAmount()
	ok := data.UserRechargeAuth(uid)
	if !ok {
		return &rpc.UserRechargeResponse{Balance: 0}, nil
	}
	balance := ss.SqlUserRecharge(uid, amount)
	return &rpc.UserRechargeResponse{Balance: balance}, nil
}

func (ss *serverService) CarOperation(ctx context.Context, request *rpc.CarOperationRequest) (*rpc.CarOperationResponse, error) {
	var ret rpc.CarOperationResult
	uid := request.GetUId()
	license := request.GetLicense()
	nowTime := time.Now().Unix()
	switch request.GetOperation() {
	case rpc.CarOperation_OPERATION_ADD:
		ret = data.UserAddCarAuth(uid, license)
		if ret == rpc.CarOperationResult_OPERATION_ADD_SUCCESS {
			ss.SqlAddCar(uid, license, nowTime)
		}
	case rpc.CarOperation_OPERATION_DELETE:
		ret = data.UserDeleteCarAuth(uid, license)
		if ret == rpc.CarOperationResult_OPERATION_DELETE_SUCCESS {
			ss.SqlDeleteCar(uid, license)
		}
	case rpc.CarOperation_OPERATION_CHANGE:
		newlicense := request.GetNewLicense()
		ret = data.UserChangeCarAuth(uid, license)
		if ret == rpc.CarOperationResult_OPERATION_CHANGE_SUCCESS {
			ss.SqlChangeCar(uid, license, newlicense, nowTime)
		}
	}
	return &rpc.CarOperationResponse{Result: ret}, nil
}
