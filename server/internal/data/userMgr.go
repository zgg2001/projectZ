package data

import "log"

type UserMgr struct {
	userArr    []user
	idMap      map[int32]*user
	licenseMap map[string]*user
}

func (um *UserMgr) Init() {
	log.Println("hello init from usermgr")
}
