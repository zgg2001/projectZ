package transmission

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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

	return nil
}
