package data

import (
	"log"

	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

func UserIsExists(username string) bool {
	exist := RedisCheckUserIsExists(username)
	// Todo add mysql
	return exist
}

func UserIsExistsByUid(uid int32) bool {
	exist := RedisCheckUserIsExistsByUid(uid)
	// Todo add mysql
	return exist
}

func LicenseIsExists(license string) bool {
	exist := RedisCheckLicenseIsExists(license)
	// Todo add mysql
	return exist
}

func CheckCarIsEntered(license string) (bool, error) {
	ok := RedisCheckCarIsEntered(license)
	if ok {
		return true, nil
	}
	// Todo add mysql
	return false, ErrParkingRecordDuplicateRecord
}

func LoginAuth(username, password string) (int32, rpc.LoginResult) {
	ok, uid, getPassword := getPasswordByUsername(username)
	if !ok {
		return -1, rpc.LoginResult_LOGIN_FAIL_NOT_EXIST
	}
	changedPasswd := GetMD5Hash(password)
	if getPassword == changedPasswd {
		return uid, rpc.LoginResult_LOGIN_SUCCESS
	}
	return -1, rpc.LoginResult_LOGIN_FAIL_WRONG_PASSWORD
}

func getPasswordByUsername(username string) (bool, int32, string) {
	ok, uid, password := RedisGetPasswordByUsername(username)
	if ok {
		return true, uid, password
	}
	// Todo add mysql
	return false, -1, ""
}

func UserRegistrationAuth(username string) rpc.RegistrationResult {
	if exist := UserIsExists(username); exist {
		return rpc.RegistrationResult_REGISTRATION_FAIL_ALREADY_EXIST
	}
	return rpc.RegistrationResult_REGISTRATION_SUCCESS
}

func UserAddCarAuth(uid int32, license string) rpc.CarOperationResult {
	if exist := LicenseIsExists(license); exist {
		return rpc.CarOperationResult_OPERATION_ADD_FAIL_ALREADY_EXIST
	}
	if exist := UserIsExistsByUid(uid); !exist {
		return rpc.CarOperationResult_OPERATION_ADD_FAIL_USER_NOT_EXIST
	}
	return rpc.CarOperationResult_OPERATION_ADD_SUCCESS
}

func UserDeleteCarAuth(uid int32, license string) rpc.CarOperationResult {
	if exist := LicenseIsExists(license); !exist {
		return rpc.CarOperationResult_OPERATION_DELETE_FAIL_NOT_EXIST
	}
	if exist := UserIsExistsByUid(uid); !exist {
		return rpc.CarOperationResult_OPERATION_DELETE_FAIL_USER_NOT_EXIST
	}
	if isEntered, _ := CheckCarIsEntered(license); isEntered {
		return rpc.CarOperationResult_OPERATION_DELETE_FAIL_ENTERED
	}
	return rpc.CarOperationResult_OPERATION_DELETE_SUCCESS
}

func UserChangeCarAuth(uid int32, license string) rpc.CarOperationResult {
	if exist := LicenseIsExists(license); !exist {
		return rpc.CarOperationResult_OPERATION_CHANGE_FAIL_NOT_EXIST
	}
	if exist := UserIsExistsByUid(uid); !exist {
		return rpc.CarOperationResult_OPERATION_CHANGE_FAIL_USER_NOT_EXIST
	}
	if isEntered, _ := CheckCarIsEntered(license); isEntered {
		return rpc.CarOperationResult_OPERATION_CHANGE_FAIL_ENTERED
	}
	return rpc.CarOperationResult_OPERATION_CHANGE_SUCCESS
}

// 获取用户全部车牌
func UserGetLicenseArr(uid int32) []string {
	licenseArr, err := RedisGetLicensesByUID(uid)
	if err != nil {
		log.Print(err)
	}
	// Todo add mysql
	return licenseArr
}
