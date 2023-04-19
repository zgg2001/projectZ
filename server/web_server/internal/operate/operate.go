package operate

import (
	"context"

	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

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
