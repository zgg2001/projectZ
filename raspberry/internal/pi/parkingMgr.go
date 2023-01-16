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
	mgr.Count = count
	mgr.LastSubscript = 0
	for id := 1; id <= count; id++ {
		mgr.Spaces = append(mgr.Spaces, Parking{Id: id, IsUsing: false, License: ""})
	}
	mgr.Client = cli
}

func (mgr *ParkingMgr) DriveIntoCar(license string) error {

	_, ok := mgr.SpaceMap[license]
	if ok {
		return ErrLicenseAlreadyExists
	}

	tempSubscript := mgr.LastSubscript
	for count := 0; count < mgr.Count; count++ {
		tempSubscript = (tempSubscript + 1) % mgr.Count
		parking := &mgr.Spaces[tempSubscript]
		if parking.GetStatus() == EmptyParkingSpace {
			mgr.SpaceMap[license] = parking
			parking.DriveInto(license, mgr.Client)
			mgr.LastSubscript = tempSubscript
			return NoErr
		}
	}

	mgr.LastSubscript = 0
	return ErrNoParkingSpace
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
