package data

import "github.com/zgg2001/projectZ/server/user_server/pkg/rpc"

// Todo fix logic

func ParkingLoginAuth(pid int32, password string) rpc.LoginResult {
	ok, getPassword := parkingGetPasswordByUsername(pid)
	if !ok {
		return rpc.LoginResult_LOGIN_FAIL_NOT_EXIST
	}
	changedPasswd := GetMD5Hash(password)
	if getPassword == changedPasswd {
		return rpc.LoginResult_LOGIN_SUCCESS
	}
	return rpc.LoginResult_LOGIN_FAIL_WRONG_PASSWORD
}

func parkingGetPasswordByUsername(uid int32) (bool, string) {
	/*ok, uid, password := RedisGetPasswordByUsername(uid)
	if ok {
		return true, password
	}*/
	return false, ""
}
