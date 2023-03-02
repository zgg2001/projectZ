package operate

import (
	"context"
	"time"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func (ss *serverService) UserLogin(ctx context.Context, request *rpc.UserLoginRequest) (*rpc.UserLoginResponse, error) {
	username := request.GetUsername()
	password := request.GetPassword()
	uid, ret := ss.uMgr.LoginAuth(username, password)
	return &rpc.UserLoginResponse{Result: ret, UId: uid}, nil
}

func (ss *serverService) UserRegistration(ctx context.Context, request *rpc.UserRegistrationRequest) (*rpc.UserRegistrationResponse, error) {
	username := request.GetUsername()
	password := request.GetPassword()
	ret := ss.uMgr.UserRegistrationAuth(username, password)
	nowTime := time.Now().Unix()
	if ret == rpc.RegistrationResult_REGISTRATION_SUCCESS {
		ss.SqlRegisterUser(username, password, nowTime)
	}
	return &rpc.UserRegistrationResponse{Result: ret}, nil
}

func (ss *serverService) CarOperation(ctx context.Context, request *rpc.CarOperationRequest) (*rpc.CarOperationResponse, error) {
	var ret rpc.CarOperationResult
	uid := request.GetUId()
	license := request.GetLicense()
	nowTime := time.Now().Unix()
	switch request.GetOperation() {
	case rpc.CarOperation_OPERATION_ADD:
		ret = ss.uMgr.UserAddCarAuth(uid, license)
		if ret == rpc.CarOperationResult_OPERATION_ADD_SUCCESS {
			ss.SqlAddCar(uid, license, nowTime)
		}
	case rpc.CarOperation_OPERATION_DELETE:
		ret = ss.uMgr.UserDeleteCarAuth(uid, license)
		if ret == rpc.CarOperationResult_OPERATION_DELETE_SUCCESS {
			ss.SqlDeleteCar(uid, license)
		}
	case rpc.CarOperation_OPERATION_CHANGE:
		newlicense := request.GetNewLicense()
		ret = ss.uMgr.UserChangeCarAuth(uid, license)
		if ret == rpc.CarOperationResult_OPERATION_CHANGE_SUCCESS {
			ss.SqlChangeCar(uid, license, newlicense)
		}
	}
	return &rpc.CarOperationResponse{Result: ret}, nil
}
