package data

type car struct {
	license         string
	parkingPtr      *parking
	parkingSpacePtr *parkingSpace
	entryTime       int64
}

type user struct {
	id       int32
	balance  int32
	username string
	cars     []car
	carMap   map[string]*car
}
