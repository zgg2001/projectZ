package data

import (
	"fmt"
	"log"
)

type UserMgr struct {
	userArr    []user
	idMap      map[int32]*user
	licenseMap map[string]*user
}

func (um *UserMgr) Init() error {

	log.Println("UserMgr init ...")

	// read and load
	userRet, err := ReadUserTbl()
	if err != nil {
		return err
	}

	for user := range userRet {
		fmt.Println(user)
	}

	return nil
}
