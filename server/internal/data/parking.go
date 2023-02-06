package data

type parking struct {
	parkingSpaceArr []parkingSpace
	spaceMap        map[int32]*parkingSpace
	count           int32
	temperature     int32
	humidity        int32
	weather         string
	info            string
}
