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

func (um *UserMgr) Init(pm *ParkingMgr) error {

	log.Println("UserMgr init ...")

	um.idMap = make(map[int32]*user)
	um.licenseMap = make(map[string]*user)

	// read and load user
	userRet, err := ReadUserTbl()
	if err != nil {
		return err
	}
	for _, tempUser := range userRet {
		fmt.Println(tempUser)
		u := user{
			id:           tempUser.Id,
			balance:      tempUser.Balance,
			username:     tempUser.Username,
			creationTime: tempUser.CreationTime,
			lastModified: tempUser.LastModified,
			cars:         nil,
			carMap:       make(map[string]*car),
		}
		um.userArr = append(um.userArr, u)
		um.idMap[u.id] = &u
	}

	// read and load user license
	licenseRet, err := ReadLicenseTbl()
	if err != nil {
		return err
	}
	for _, tempLicense := range licenseRet {
		fmt.Println(tempLicense)
		i := tempLicense.Id
		c := car{
			license:         tempLicense.License,
			parkingPtr:      nil,
			parkingSpacePtr: nil,
			checkInTime:     tempLicense.CheckInTime,
			entryTime:       0,
		}
		u := um.idMap[i]
		u.cars = append(u.cars, c)
		u.carMap[c.license] = &c
		um.licenseMap[c.license] = u
	}

	// read and load record
	recordRet, err := ReadRecordTbl()
	if err != nil {
		return err
	}
	for _, tempRecord := range recordRet {
		fmt.Println(tempRecord)
		l := tempRecord.License
		pptr, sptr, err := pm.MgrGetParkingPtrPair(tempRecord.PId, tempRecord.SId)
		if err != nil {
			return err
		}
		tCar := um.licenseMap[l].carMap[l]
		tCar.SetParkingSpace(pptr, sptr, tempRecord.EntryTime)
	}

	return nil
}

func (um *UserMgr) GetUser(license string) (bool, *user) {
	if uptr, ok := um.licenseMap[license]; ok {
		return true, uptr
	}
	return false, nil
}
