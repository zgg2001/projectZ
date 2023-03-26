package data

import "sync/atomic"

type parkingSpace struct {
	id          int32
	temperature int32
	humidity    int32
	alarm       int32
}

func (p *parkingSpace) UpdateParkingSpaceData(temp, hum, alarm int32) {
	atomic.StoreInt32(&p.temperature, int32(temp))
	atomic.StoreInt32(&p.humidity, int32(hum))
	atomic.StoreInt32(&p.alarm, int32(alarm))
}

func (p *parkingSpace) GetParkingSpaceData() (id, temp, hum, alarm int32) {
	id = p.id
	temp = atomic.LoadInt32(&p.temperature)
	hum = atomic.LoadInt32(&p.humidity)
	alarm = atomic.LoadInt32(&p.alarm)
	return
}
