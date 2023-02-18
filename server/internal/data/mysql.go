package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type parkingRow struct {
	Id      int32
	Count   int32
	address string
}

type userRow struct {
	Id           int32
	Username     string
	Balance      int32
	CreationTime int64
	LastModified int64
}

type licenseRow struct {
	Id          int32
	License     string
	CheckInTime int64
}

type recordRow struct {
	License   string
	PId       int32
	SId       int32
	EntryTime int64
}

func InitDB() error {

	err := connectDB()
	if err != nil {
		return err
	}

	err = checkTable()
	if err != nil {
		return err
	}

	return nil
}

func ReadParkingTbl() ([]*parkingRow, error) {

	var data []*parkingRow

	ret, err := DB.Query(SqlSelectParkingTbl)
	if err != nil {
		return nil, err
	}
	defer ret.Close()

	for ret.Next() {
		var d parkingRow
		err := ret.Scan(&d.Id, &d.Count, &d.address)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}

	return data, nil
}

func ReadUserTbl() ([]*userRow, error) {

	var data []*userRow

	ret, err := DB.Query(SqlSelectUserTbl)
	if err != nil {
		return nil, err
	}
	defer ret.Close()

	for ret.Next() {
		var d userRow
		err := ret.Scan(&d.Id, &d.Username, &d.Balance, &d.CreationTime, &d.LastModified)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}

	return data, nil
}

func ReadLicenseTbl() ([]*licenseRow, error) {

	var data []*licenseRow

	ret, err := DB.Query(SqlSelectLicenseTbl)
	if err != nil {
		return nil, err
	}
	defer ret.Close()

	for ret.Next() {
		var d licenseRow
		err := ret.Scan(&d.Id, &d.License, &d.CheckInTime)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}

	return data, nil
}

func ReadRecordTbl() ([]*recordRow, error) {

	var data []*recordRow

	ret, err := DB.Query(SqlSelectRecordTbl)
	if err != nil {
		return nil, err
	}
	defer ret.Close()

	for ret.Next() {
		var d recordRow
		err := ret.Scan(&d.License, &d.PId, &d.SId, &d.EntryTime)
		if err != nil {
			return nil, err
		}
		data = append(data, &d)
	}

	return data, nil
}

func InsertRecordTbl(license string, pid, sid int32, etime int64) {
	_, err := DB.Query(SqlInsertRecordTbl, license, pid, sid, etime)
	if err != nil {
		log.Println(err)
	}
}

func InsertParkingRecordTbl(license string, pid, sid, state int32, time int64) {
	_, err := DB.Query(SqlInsertParkingRecordTbl, license, pid, sid, state, time)
	if err != nil {
		log.Println(err)
	}
}

func DeleteRecordTbl(license string) {
	_, err := DB.Exec(SqlDeleteRecordTbl, license)
	if err != nil {
		log.Println(err)
	}
}

func SelectRecordTbl(license string) (int64, error) {
	ret, err := DB.Query(SqlSelectRecordUsingLicenseTbl, license)
	if err != nil {
		return 0, err
	}
	if ret.Next() {
		var d recordRow
		err := ret.Scan(&d.License, &d.PId, &d.SId, &d.EntryTime)
		if err != nil {
			return 0, err
		}
		return d.EntryTime, nil
	}
	return 0, ErrParkingRecordNotFound
}

func connectDB() error {

	var err error
	log.Println("Connect DB ...")

	DB, err = sql.Open(DriverName, DataSourceName)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(20)

	return nil
}

func checkTable() error {

	log.Println("Check DB tables ...")

	ret, err := DB.Query(SqlGetTableNum)
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
		_, err = DB.Exec(SqlCreateParkingTbl)
		if err != nil {
			return err
		}
		_, err := DB.Exec(SqlCreateUserTbl)
		if err != nil {
			return err
		}
		_, err = DB.Exec(SqlCreateLicenseTbl)
		if err != nil {
			return err
		}
		_, err = DB.Exec(SqlCreateRecordTbl)
		if err != nil {
			return err
		}
		_, err = DB.Exec(SqlCreateParkingRecordTbl)
		if err != nil {
			return err
		}
	} else if tableNum != 5 {
		return ErrTableNum
	}

	return nil
}
