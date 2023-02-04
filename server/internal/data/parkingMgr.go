package data

import "log"

type ParkingMgr struct {
	parkingArr []parking
	idMap      map[int32]*parking
}

func (pm *ParkingMgr) Init() {
	log.Println("hello init from parkingmgr")
}
