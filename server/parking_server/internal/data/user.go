package data

import (
	"log"
)

func UserInit() error {

	log.Println("User data load ...")

	// read and load user
	userRet, err := ReadUserTbl()
	if err != nil {
		return err
	}
	for _, tempUser := range userRet {
		err = RedisAddUser(tempUser)
		if err != nil {
			return err
		}
	}

	// read and load user license
	licenseRet, err := ReadLicenseTbl()
	if err != nil {
		return err
	}
	for _, tempLicense := range licenseRet {
		err = RedisAddLicense(tempLicense)
		if err != nil {
			return err
		}
	}

	// read and load record
	recordRet, err := ReadRecordTbl()
	if err != nil {
		return err
	}
	for _, tempRecord := range recordRet {
		err = RedisAddRecord(tempRecord)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetUserByLicense(license string) (ok bool, uid string) {
	ok, uid = RedisGetUidByLicense(license)
	if ok {
		return
	}
	return
}

func GetBalanceByUid(uid string) int32 {
	ok, balance := RedisGetBalanceByUid(uid)
	if ok {
		return balance
	}
	return 0
}

func CheckCarIsEntered(license string) error {
	ok := RedisCheckCarIsEntered(license)
	if ok {
		return nil
	}
	return ErrParkingRecordDuplicateRecord
}
