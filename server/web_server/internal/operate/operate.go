package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

func checkRegister(username, password string) rpc.RegistrationResult {
	request := &rpc.UserRegistrationRequest{
		Username: username,
		Password: password,
	}
	resp, err := RPCClient.UserRegistration(context.Background(), request)
	if err != nil {
		return rpc.RegistrationResult_REGISTRATION_FAIL_ALREADY_EXIST
	}
	return resp.GetResult()
}

func checkLogin(username, password string) (int32, rpc.LoginResult) {
	request := &rpc.UserLoginRequest{
		Username: username,
		Password: password,
	}
	resp, err := RPCClient.UserLogin(context.Background(), request)
	if err != nil {
		return -1, rpc.LoginResult_LOGIN_FAIL_NOT_EXIST
	}
	return resp.GetUId(), resp.GetResult()
}

func getInfo(uid int32) ([]*rpc.CarInfo, error) {
	request := &rpc.GetUserDataRequest{
		UId: uid,
	}
	resp, err := RPCClient.GetUserData(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return resp.GetCarInfoArr(), nil
}

func recharge(uid, amount int32) int32 {
	return 0
}

func carOperator(uid int32, operation rpc.CarOperation, license, newLicense string) rpc.CarOperationResult {
	return rpc.CarOperationResult_OPERATION_ADD_SUCCESS
}
