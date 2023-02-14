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

	pm.idMap = make(map[int32]*parking)

	// read and load parking
	parkingRet, err := ReadParkingTbl()
	if err != nil {
		return err
	}
	for _, tempParking := range parkingRet {
		fmt.Println(tempParking)
		p := parking{
			id:              tempParking.Id,
			parkingSpaceArr: nil,
			count:           tempParking.Count,
			temperature:     0,
			humidity:        0,
			weather:         0,
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

func (pm *ParkingMgr) MgrGetParkingPtrPair(pid, sid int32) (*parking, *parkingSpace, error) {

	var pptr *parking
	var ok bool

	if pptr, ok = pm.idMap[pid]; !ok {
		return nil, nil, ErrPIdNotFound
	}
	sptr, err := pptr.GetParkingPtr(sid)
	if err != nil {
		return nil, nil, err
	}

	return pptr, sptr, nil
}

func (pm *ParkingMgr) MgrGetParkingPtr(pid int32) (*parking, error) {
	var pptr *parking
	var ok bool
	if pptr, ok = pm.idMap[pid]; !ok {
		return nil, ErrPIdNotFound
	}
	return pptr, nil
}
