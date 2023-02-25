package data

import (
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

type car struct {
	license         string
	parkingPtr      unsafe.Pointer // *parking
	parkingSpacePtr unsafe.Pointer // *parkingSpace
	checkInTime     int64
	entryTime       int64
}

type user struct {
	id           int32
	balance      int32
	username     string
	creationTime int64
	lastModified int64
	cars         []car
	carMap       map[string]*car
	carMapLock   *sync.RWMutex
}

func (c *car) SetParkingSpace(pptr *parking, sptr *parkingSpace, etime int64) {
	atomic.StorePointer(&c.parkingPtr, unsafe.Pointer(pptr))
	atomic.StorePointer(&c.parkingSpacePtr, unsafe.Pointer(sptr))
	atomic.StoreInt64(&c.entryTime, etime)
}

func (c *car) GetCarPtrArr() *rpc.CarInfo {
	var pptr *parking = (*parking)(atomic.LoadPointer(&c.parkingPtr))
	var sptr *parkingSpace = (*parkingSpace)(atomic.LoadPointer(&c.parkingSpacePtr))
	temp, hum, weather, address := pptr.GetParkingData()
	sid, stemp, shum, alarm := sptr.GetParkingSpaceData()
	ret := &rpc.CarInfo{
		PTemperature: temp,
		PHumidity:    hum,
		PWeather:     weather,
		PAddress:     address,
		SId:          sid,
		STemperature: stemp,
		SHumidity:    shum,
		SAlarm:       rpc.Alarm(alarm),
	}
	return ret
}

func (u *user) GetBalance() int32 {
	return atomic.LoadInt32(&u.balance)
}

func (u *user) SetBalance(balance int32) {
	atomic.StoreInt32(&u.balance, balance)
}

func (u *user) GetCarPtr(license string) (*car, error) {
	u.carMapLock.RLock()
	defer u.carMapLock.RUnlock()
	if cptr, ok := u.carMap[license]; ok {
		if cptr.entryTime == 0 && cptr.parkingPtr == nil && cptr.parkingSpacePtr == nil {
			return cptr, nil
		}
		return nil, ErrParkingRecordDuplicateRecord
	}
	return nil, ErrUserLicenseNotFound
}

func (u *user) GetCarPtrArr() []*car {
	var ret []*car
	for _, car := range u.cars {
		ret = append(ret, &car)
	}
	return ret
}
