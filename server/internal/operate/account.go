package operate

import (
	"context"

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
	ret := ss.uMgr.RegistrationAuth(username, password)
	if ret == rpc.RegistrationResult_REGISTRATION_SUCCESS {
		ss.RegisterUser(username, password)
	}
	return &rpc.UserRegistrationResponse{Result: ret}, nil
}
