package data

import "sync/atomic"

type parking struct {
	id              int32
	parkingSpaceArr []parkingSpace
	count           int32
	temperature     int32
	humidity        int32
	weather         int32
	info            string
}

func (p *parking) GetParkingPtr(sid int32) (*parkingSpace, error) {

	var sptr *parkingSpace

	if sid > p.count || sid <= 0 {
		return nil, ErrSIdNotFound
	}

	sptr = &p.parkingSpaceArr[sid-1]
	return sptr, nil
}

func (p *parking) UpdateParkingData(t, h, w int32) {
	atomic.StoreInt32(&p.temperature, int32(t))
	atomic.StoreInt32(&p.humidity, int32(h))
	atomic.StoreInt32(&p.weather, int32(w))
}
