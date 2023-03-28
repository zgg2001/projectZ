package pi

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ParkingMgrOperation interface {
	Init(count int)
	DriveIntoCar(license string) error
	DriveOutCar(license string) error
}

type ParkingMgr struct {
	Spaces        []Parking
	SpaceMap      map[string]*Parking
	Count         int
	LastSubscript int
	Client        *mqtt.Client
}

func (mgr *ParkingMgr) Init(count int, cli *mqtt.Client) {
	mgr.SpaceMap = make(map[string]*Parking)
	mgr.Count = count
	mgr.LastSubscript = 0
	for id := 1; id <= count; id++ {
		mgr.Spaces = append(mgr.Spaces, Parking{Id: id, IsUsing: false, License: ""})
	}
	mgr.Client = cli
}

func (mgr *ParkingMgr) FindEmptySpace() (int, error) {

	tempSubscript := mgr.LastSubscript
	for count := 0; count < mgr.Count; count++ {
		tempSubscript = (tempSubscript + 1) % mgr.Count
		parking := &mgr.Spaces[tempSubscript]
		if parking.GetStatus() == EmptyParkingSpace {
			return tempSubscript, nil
		}
	}

	mgr.LastSubscript = 0
	return 0, ErrNoParkingSpace
}

func (mgr *ParkingMgr) DriveIntoCar(license string, sub int) error {

	_, ok := mgr.SpaceMap[license]
	if ok {
		return ErrLicenseAlreadyExists
	}

	parking := &mgr.Spaces[sub]
	mgr.SpaceMap[license] = parking
	parking.DriveInto(license, mgr.Client)
	mgr.LastSubscript = sub

	return NoErr
}

func (mgr *ParkingMgr) DriveOutCar(license string) error {

	parking, ok := mgr.SpaceMap[license]
	if !ok {
		return ErrLicenseNotExists
	}

	delete(mgr.SpaceMap, license)
	parking.DriveOut(mgr.Client)

	return NoErr
}
