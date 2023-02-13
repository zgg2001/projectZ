package data

import (
	"sync/atomic"
	"unsafe"
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
}

func (c *car) SetParkingSpace(pptr *parking, sptr *parkingSpace, etime int64) {
	atomic.StorePointer(&c.parkingPtr, unsafe.Pointer(pptr))
	atomic.StorePointer(&c.parkingSpacePtr, unsafe.Pointer(sptr))
	atomic.StoreInt64(&c.entryTime, etime)
}

func (u *user) GetBalance() int32 {
	return atomic.LoadInt32(&u.balance)
}

func (u *user) SetBalance(balance int32) {
	atomic.StoreInt32(&u.balance, balance)
}

func (u *user) GetCarPtr(license string) (*car, error) {
	if cptr, ok := u.carMap[license]; ok {
		return cptr, nil
	}
	return nil, ErrUserLicenseNotFound
}
