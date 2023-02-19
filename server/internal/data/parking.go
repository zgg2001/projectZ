package data

import "sync/atomic"

type parking struct {
	id              int32
	parkingSpaceArr []parkingSpace
	count           int32
	temperature     int32
	humidity        int32
	weather         int32
	address         string
}

func (p *parking) GetParkingPtr(sid int32) (*parkingSpace, error) {

	var sptr *parkingSpace

	if sid > p.count || sid <= 0 {
		return nil, ErrSIdNotFound
	}

	sptr = &p.parkingSpaceArr[sid-1]
	return sptr, nil
}

func (p *parking) UpdateParkingData(temp, hum, weather int32) {
	atomic.StoreInt32(&p.temperature, int32(temp))
	atomic.StoreInt32(&p.humidity, int32(hum))
	atomic.StoreInt32(&p.weather, int32(weather))
}

func (p *parking) GetParkingData() (temp, hum, weather int32, address string) {
	temp = atomic.LoadInt32(&p.temperature)
	hum = atomic.LoadInt32(&p.humidity)
	weather = atomic.LoadInt32(&p.weather)
	address = p.address
	return
}
