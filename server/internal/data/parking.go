package data

type parking struct {
	id              int32
	parkingSpaceArr []parkingSpace
	count           int32
	temperature     int32
	humidity        int32
	weather         string
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
