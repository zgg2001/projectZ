package data

import (
	"fmt"
	"log"
)

type ParkingMgr struct {
	parkingArr []parking
	idMap      map[int32]*parking
}

func (pm *ParkingMgr) Init() error {

	log.Println("ParkingMgr init ...")

	// read and load parking
	userRet, err := ReadParkingTbl()
	if err != nil {
		return err
	}
	for _, tempParking := range userRet {
		fmt.Println(tempParking)
		p := parking{
			id:              tempParking.Id,
			parkingSpaceArr: nil,
			count:           tempParking.Count,
			temperature:     0,
			humidity:        0,
			weather:         "",
			info:            tempParking.Info,
		}
		pm.parkingArr = append(pm.parkingArr, p)
		pm.idMap[p.id] = &p
		for sid := startId; sid <= p.count; sid++ {
			p.parkingSpaceArr = append(p.parkingSpaceArr, parkingSpace{
				id:          sid,
				temperature: 0,
				humidity:    0,
				alarm:       NoAlarm,
			})
		}
	}

	return nil
}
