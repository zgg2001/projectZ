package data

import "sync/atomic"

type parkingSpace struct {
	id          int32
	temperature int32
	humidity    int32
	alarm       int32
}

func (p *parkingSpace) UpdateParkingSpaceData(t, h, a int32) {
	atomic.StoreInt32(&p.temperature, int32(t))
	atomic.StoreInt32(&p.humidity, int32(h))
	atomic.StoreInt32(&p.alarm, int32(a))
}
