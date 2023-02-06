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

const (
	SqlGetTableNum        = "SELECT COUNT(TABLE_NAME) FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'projectZ';"
	SqlCreateUserTbl      = "CREATE TABLE `z_user` (`id` int unsigned NOT NULL AUTO_INCREMENT, `username` varchar(255) NOT NULL DEFAULT 'user', `balance` int unsigned NOT NULL DEFAULT '0', `creation_time` bigint unsigned DEFAULT '0', `last_Modified` bigint unsigned DEFAULT '0', PRIMARY KEY (`id`));"
	SqlCreateLicenseTbl   = "CREATE TABLE `z_license` (`id` int unsigned NOT NULL DEFAULT '0', `license` varchar(255) NOT NULL DEFAULT '豫A88888', `checkin_time` bigint unsigned DEFAULT '0', PRIMARY KEY (`id`));"
	SqlCreateEntryTimeTbl = "CREATE TABLE `z_record` (`license` varchar(255) NOT NULL DEFAULT '豫A88888', `entry_time` bigint unsigned DEFAULT '0', PRIMARY KEY (`license`));"

	SqlSelectUserTbl    = "SELECT * FROM z_user;"
	SqlSelectLicenseTbl = "SELECT * FROM z_license;"
	SqlSelectRecordTbl  = "SELECT * FROM z_record;"
)
