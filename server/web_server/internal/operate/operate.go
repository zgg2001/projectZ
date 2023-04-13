package operate

import (
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

func checkLogin(username, password string) (int32, rpc.LoginResult) {
	return 1, rpc.LoginResult_LOGIN_SUCCESS
}
