package data

type user struct {
	id              int32
	balance         int32
	username        string
	license         string
	parkingPtr      *parking
	parkingSpacePtr *parkingSpace
}
