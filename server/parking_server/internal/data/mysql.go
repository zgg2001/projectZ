package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ParkingRow struct {
	Id      int32
	Count   int32
	address string
}

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

func ReadParkingTbl() ([]*ParkingRow, error) {

	var data []*ParkingRow

	ret, err := MySqlClient.Query(SqlSelectParkingTbl)
	if err != nil {
		return nil, err
	}
	defer ret.Close()

	for ret.Next() {
		var d ParkingRow
		err := ret.Scan(&d.Id, &d.Count, &d.address)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}

	return data, nil
}

func ReadUserTbl() ([]*UserRow, error) {

	var data []*UserRow

	ret, err := MySqlClient.Query(SqlSelectUserTbl)
	if err != nil {
		return nil, err
	}
	defer ret.Close()

	for ret.Next() {
		var d UserRow
		err := ret.Scan(&d.Id, &d.Username, &d.Password, &d.Balance, &d.CreationTime, &d.LastModified)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}

	return data, nil
}

func ReadLicenseTbl() ([]*LicenseRow, error) {

	var data []*LicenseRow

	ret, err := MySqlClient.Query(SqlSelectLicenseTbl)
	if err != nil {
		return nil, err
	}
	defer ret.Close()

	for ret.Next() {
		var d LicenseRow
		err := ret.Scan(&d.License, &d.Id, &d.CheckInTime)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}

	return data, nil
}

func ReadRecordTbl() ([]*RecordRow, error) {

	var data []*RecordRow

	ret, err := MySqlClient.Query(SqlSelectRecordTbl)
	if err != nil {
		return nil, err
	}
	defer ret.Close()

	for ret.Next() {
		var d RecordRow
		err := ret.Scan(&d.License, &d.PId, &d.SId, &d.EntryTime)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}

	return data, nil
}

func SelectRecordTbl(license string) (int64, error) {
	ret, err := MySqlClient.Query(SqlSelectRecordUsingLicenseTbl, license)
	if err != nil {
		return 0, err
	}
	if ret.Next() {
		var d RecordRow
		err := ret.Scan(&d.License, &d.PId, &d.SId, &d.EntryTime)
		if err != nil {
			return 0, err
		}
		return d.EntryTime, nil
	}
	return 0, ErrParkingRecordNotFound
}

func InsertRecordTbl(license string, pid, sid int32, etime int64) {
	_, err := MySqlClient.Query(SqlInsertRecordTbl, license, pid, sid, etime)
	if err != nil {
		log.Println(err)
	}
}

func InsertParkingRecordTbl(license string, pid, sid, state int32, time int64) {
	_, err := MySqlClient.Query(SqlInsertParkingRecordTbl, license, pid, sid, state, time)
	if err != nil {
		log.Println(err)
	}
}

func ChangeUserBalanceTbl(uid string, balance int32) error {
	_, err := MySqlClient.Exec(SqlUpdateUserBalanceTbl, balance, uid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRecordTbl(license string) {
	_, err := MySqlClient.Exec(SqlDeleteRecordTbl, license)
	if err != nil {
		log.Println(err)
	}
}

func connectMysql() error {

	var err error
	log.Println("Connect MySql ...")

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

	if tableNum == 0 {
		log.Println("Create DB tables ...")
		_, err = MySqlClient.Exec(SqlCreateParkingTbl)
		if err != nil {
			return err
		}
		_, err := MySqlClient.Exec(SqlCreateUserTbl)
		if err != nil {
			return err
		}
		_, err = MySqlClient.Exec(SqlCreateLicenseTbl)
		if err != nil {
			return err
		}
		_, err = MySqlClient.Exec(SqlCreateRecordTbl)
		if err != nil {
			return err
		}
		_, err = MySqlClient.Exec(SqlCreateParkingRecordTbl)
		if err != nil {
			return err
		}
	} else if tableNum != 5 {
		return ErrTableNum
	}

	return nil
}
