package data

import (
	"database/sql"
	"errors"

	"github.com/go-redis/redis"
)

// redis
const (
	RedisAddr = "localhost:7892"

	ParingInfoPrefix         = "z-parking-info-"
	ParingSpaceDataPrefix    = "z-parking-space-data-"
	ParingSpaceLicensePrefix = "z-parking-space-License-"

	UserInfoPrefix  = "z-user-info-"
	UserLoginMapKey = "z-user-login-map"

	LicenseInfoPrefix     = "z-license-info-"
	LicenseSetByUIDPrefix = "z-license-set-by-uid"
)

// mysql
const (
	DriverName     = "mysql"
	DataSourceName = "root:password@/projectZ"
)

var (
	MySqlClient *sql.DB
	RedisClient *redis.Client

	ErrTableNum error = errors.New("wrong number of tables")
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
		"SET license=?, checkin_time=? " +
		"WHERE license=?;"

	SqlSelectNextPrimaryId = "SELECT AUTO_INCREMENT from INFORMATION_SCHEMA.TABLES where TABLE_NAME=?;"
)
