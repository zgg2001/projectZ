package data

type car struct {
	license         string
	parkingPtr      *parking
	parkingSpacePtr *parkingSpace
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
