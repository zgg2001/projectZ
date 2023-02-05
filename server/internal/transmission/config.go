package transmission

import "database/sql"

// grpc
const (
	TLS           = "tcp"
	RPCServerIddr = ":8888"
	ServerPemPath = "./auth/server.pem"
	ServerKeyPath = "./auth/server.key"
	CACrtPath     = "./auth/ca.crt"
)

// mysql
const (
	DriverName     = "mysql"
	DataSourceName = "root:password@/projectZ"
)

var (
	DB *sql.DB
)

const (
	SqlGetTableNum        = "SELECT COUNT(TABLE_NAME) FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'projectZ';"
	SqlCreateUserTbl      = "CREATE TABLE `z_user` (`id` int unsigned NOT NULL AUTO_INCREMENT, `username` varchar(255) NOT NULL DEFAULT 'user', `balance` int unsigned NOT NULL DEFAULT '0', `creation_time` bigint unsigned DEFAULT '0', `last_Modified` bigint unsigned DEFAULT '0', PRIMARY KEY (`id`));"
	SqlCreateLicenseTbl   = "CREATE TABLE `z_license` (`id` int unsigned NOT NULL DEFAULT '0', `license` varchar(255) NOT NULL DEFAULT '豫A88888', `checkin_time` bigint unsigned DEFAULT '0', PRIMARY KEY (`id`));"
	SqlCreateEntryTimeTbl = "CREATE TABLE `z_record` (`license` varchar(255) NOT NULL DEFAULT '豫A88888', `entry_time` bigint unsigned DEFAULT '0', PRIMARY KEY (`license`));"
)
