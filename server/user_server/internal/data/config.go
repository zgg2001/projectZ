package data

import (
	"database/sql"
	"errors"
)

// mysql
const (
	DriverName     = "mysql"
	DataSourceName = "root:password@/projectZ"
)

var (
	DB *sql.DB

	ErrTableNum error = errors.New("wrong number of tables")
)

// data
const (
	startId int32 = 1
)

var (
	ErrPIdNotFound                  error = errors.New("parking id not found")
	ErrSIdNotFound                  error = errors.New("parkingSpace id not found")
	ErrUserLicenseNotFound          error = errors.New("user license not found")
	ErrParkingRecordNotFound        error = errors.New("parking record not found")
	ErrParkingRecordDuplicateRecord error = errors.New("parking record duplicate record")
	ErrUserNotExist                 error = errors.New("user not exist")
)

// sql语句
const (
	SqlGetTableNum = "SELECT COUNT(TABLE_NAME) " +
		"FROM information_schema.TABLES " +
		"WHERE TABLE_SCHEMA = 'projectZ';"

	SqlInsertUserTbl = "INSERT INTO " +
		"z_user(username, password, balance, creation_time, last_Modified) " +
		"VALUES (?, ?, ?, ?, ?);"
	SqlInsertLicenseTbl = "INSERT INTO " +
		"z_license(license, id, checkin_time) " +
		"VALUES (?, ?, ?);"

	SqlDeleteLicenseTbl = "DELETE FROM z_license " +
		"WHERE license = ?;"

	SqlUpdateLicenseTbl = "UPDATE z_license " +
		"SET license=? " +
		"WHERE license=?;"

	SqlSelectNextPrimaryId    = "SELECT AUTO_INCREMENT from INFORMATION_SCHEMA.TABLES where TABLE_NAME=?;"
	SqlSelectParkingTbl       = "SELECT * FROM z_parking;"
	SqlSelectUserTbl          = "SELECT * FROM z_user;"
	SqlSelectLicenseTbl       = "SELECT * FROM z_license;"
	SqlSelectRecordTbl        = "SELECT * FROM z_record;"
	SqlSelectParkingRecordTbl = "SELECT * FROM z_parking_record;"
)
