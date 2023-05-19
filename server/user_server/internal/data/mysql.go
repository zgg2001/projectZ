package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type UserRow struct {
	Id           int32
	Username     string
	Password     string
	Balance      int32
	CreationTime int64
	LastModified int64
}

type LicenseRow struct {
	License     string
	Id          int32
	CheckInTime int64
}

type RecordRow struct {
	License   string
	PId       int32
	SId       int32
	EntryTime int64
}

func InitMySql() error {

	err := connectMysql()
	if err != nil {
		return err
	}

	err = checkTable()
	if err != nil {
		return err
	}

	return nil
}

func InsertUserTbl(username, password string, balance int32, nowTime int64) (int32, error) {

	var uid int32 = -1

	ret, err := MySqlClient.Query(SqlSelectNextPrimaryId, "z_user")
	if err != nil {
		return uid, err
	}
	defer ret.Close()
	for ret.Next() {
		err := ret.Scan(&uid)
		if err != nil {
			return uid, err
		}
	}

	_, err = MySqlClient.Query(SqlInsertUserTbl, username, password, balance, nowTime, nowTime)
	if err != nil {
		log.Println(err)
		return uid, err
	}
	return uid, nil
}

func InsertLicenseTbl(uid int32, license string, nowTime int64) error {
	_, err := MySqlClient.Query(SqlInsertLicenseTbl, license, uid, nowTime)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLicenseTbl(license string) error {
	_, err := MySqlClient.Exec(SqlDeleteLicenseTbl, license)
	if err != nil {
		return err
	}
	return nil
}

func ChangeLicenseTbl(license, newlicense string, checkInTime int64) error {
	_, err := MySqlClient.Exec(SqlUpdateLicenseTbl, newlicense, checkInTime, license)
	if err != nil {
		return err
	}
	return nil
}

func ChangeUserBalanceTbl(uid string, balance int32) error {
	_, err := MySqlClient.Exec(SqlUpdateUserBalanceTbl, balance, uid)
	if err != nil {
		return err
	}
	return nil
}

func connectMysql() error {

	var err error
	log.Println("Connect DB ...")

	MySqlClient, err = sql.Open(DriverName, DataSourceName)
	if err != nil {
		return err
	}
	err = MySqlClient.Ping()
	if err != nil {
		return err
	}
	MySqlClient.SetMaxOpenConns(50)
	MySqlClient.SetMaxIdleConns(20)

	return nil
}

func checkTable() error {

	log.Println("Check DB tables ...")

	ret, err := MySqlClient.Query(SqlGetTableNum)
	if err != nil {
		return err
	}
	defer ret.Close()

	tableNum := 0
	for ret.Next() {
		err := ret.Scan(&tableNum)
		if err != nil {
			return err
		}
	}
	if tableNum != 5 {
		return ErrTableNum
	}

	return nil
}

func MySqlGetParkingPasswordById(id int32) (bool, int32, string) {
	ret, err := MySqlClient.Query(SqlSelectParkingPasswordByPid, id)
	if err != nil {
		return false, 0, ""
	}
	if ret.Next() {
		var password string
		var count int32
		err := ret.Scan(&password, &count)
		if err != nil {
			return false, 0, ""
		}
		return true, count, password
	}
	return false, 0, ""
}

func MySqlParkingGetSpaceInfo(pid, sid int32) (bool, string, int64) {
	ret, err := MySqlClient.Query(SqlSelectParkingSpaceInfo, pid, sid)
	if err != nil {
		return false, "", 0
	}
	if ret.Next() {
		var license string
		var entryTime int64
		err := ret.Scan(&license, &entryTime)
		if err != nil || entryTime == 0 {
			return false, "", 0
		}
		return true, license, entryTime
	}
	return false, "", 0
}

func MySqlCheckCarIsEntered(license string) bool {
	ret, err := MySqlClient.Query(SqlSelectRecordByLicense, license)
	if err != nil {
		return false
	}
	if ret.Next() {
		var r RecordRow
		err := ret.Scan(&r.License, &r.PId, &r.SId, &r.EntryTime)
		if err != nil {
			return false
		}
		if r.PId >= 0 && r.SId >= 0 && r.EntryTime > 0 {
			return true
		}
		return false
	}
	return false
}

func MySqlGetPasswordByUsername(username string) (bool, int32, string) {
	ret, err := MySqlClient.Query(SqlSelectUserPasswordByUsername, username)
	if err != nil {
		return false, 0, ""
	}
	if ret.Next() {
		var uid int32
		var password string
		err := ret.Scan(&uid, &password)
		if err != nil {
			return false, 0, ""
		}
		return true, uid, password
	}
	return false, 0, ""
}
