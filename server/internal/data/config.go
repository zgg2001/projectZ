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

	NoAlarm         int32 = 0
	FireAlarm       int32 = 1
	GasAlarm        int32 = 2
	FireAndGasAlarm int32 = 3
)

var (
	ErrPIdNotFound                  error = errors.New("parking id not found")
	ErrSIdNotFound                  error = errors.New("parkingSpace id not found")
	ErrUserLicenseNotFound          error = errors.New("user license not found")
	ErrParkingRecordNotFound        error = errors.New("parking record not found")
	ErrParkingRecordDuplicateRecord error = errors.New("parking record duplicate record")
)

// sql语句
const (
	SqlGetTableNum = "SELECT COUNT(TABLE_NAME) " +
		"FROM information_schema.TABLES " +
		"WHERE TABLE_SCHEMA = 'projectZ';"

	SqlCreateParkingTbl = "CREATE TABLE `z_parking` (" +
		"`id` int unsigned NOT NULL DEFAULT '0'," +
		"`count` int unsigned NOT NULL DEFAULT '0'," +
		"`info` varchar(255) NOT NULL DEFAULT ''," +
		"PRIMARY KEY (`id`));"
	SqlCreateUserTbl = "CREATE TABLE `z_user` (" +
		"`id` int unsigned NOT NULL AUTO_INCREMENT, " +
		"`username` varchar(255) NOT NULL DEFAULT 'user'," +
		"`balance` int unsigned NOT NULL DEFAULT '0', " +
		"`creation_time` bigint unsigned DEFAULT '0', " +
		"`last_Modified` bigint unsigned DEFAULT '0', " +
		"PRIMARY KEY (`id`));"
	SqlCreateLicenseTbl = "CREATE TABLE `z_license` (" +
		"`id` int unsigned NOT NULL DEFAULT '0'," +
		"`license` varchar(255) NOT NULL DEFAULT '豫A88888'," +
		"`checkin_time` bigint unsigned DEFAULT '0', " +
		"PRIMARY KEY (`id`));"
	SqlCreateRecordTbl = "CREATE TABLE `z_record` (" +
		"`license` varchar(255) NOT NULL DEFAULT '豫A88888'," +
		"`pid` int unsigned NOT NULL DEFAULT '0'," +
		"`sid` int unsigned NOT NULL DEFAULT '0'," +
		"`entry_time` bigint unsigned DEFAULT '0'," +
		"PRIMARY KEY (`license`));"
	SqlCreateParkingRecordTbl = "CREATE TABLE `z_parking_record` (" +
		"`license` varchar(255) NOT NULL DEFAULT '豫A88888'," +
		"`pid` int unsigned NOT NULL DEFAULT '0'," +
		"`sid` int unsigned NOT NULL DEFAULT '0'," +
		"`state` TINYINT(1) NOT NULL DEFAULT '0'," +
		"`time` bigint unsigned DEFAULT '0');"

	SqlInsertRecordTbl = "INSERT INTO " +
		"z_record(license, pid, sid, entry_time)" +
		"VALUES (?, ?, ?, ?);"
	SqlInsertParkingRecordTbl = "INSERT INTO " +
		"z_parking_record(license, pid, sid, state, time)" +
		"VALUES (?, ?, ?, ?, ?);"

	SqlDeleteRecordTbl = "DELETE FROM mark_plan" +
		"WHERE license = ?;"

	SqlSelectRecordUsingLicenseTbl = "SELECT * FROM z_record WHERE license = ?"

	SqlSelectParkingTbl       = "SELECT * FROM z_parking;"
	SqlSelectUserTbl          = "SELECT * FROM z_user;"
	SqlSelectLicenseTbl       = "SELECT * FROM z_license;"
	SqlSelectRecordTbl        = "SELECT * FROM z_record;"
	SqlSelectParkingRecordTbl = "SELECT * FROM z_parking_record;"
)
