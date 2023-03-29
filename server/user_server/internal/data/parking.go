package data

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

func ParkingLoginAuth(id int32, password string) (int32, rpc.LoginResult) {
	ok, count, getPassword := getParkingPasswordById(id)
	if !ok {
		return -1, rpc.LoginResult_LOGIN_FAIL_NOT_EXIST
	}
	changedPasswd := GetMD5Hash(password)
	if getPassword == changedPasswd {
		return count, rpc.LoginResult_LOGIN_SUCCESS
	}
	return -1, rpc.LoginResult_LOGIN_FAIL_WRONG_PASSWORD
}

func getParkingPasswordById(id int32) (bool, int32, string) {
	ok, count, password := RedisGetParkingPasswordById(id)
	if ok {
		return true, count, password
	}
	// Todo add mysql
	return false, -1, ""
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
