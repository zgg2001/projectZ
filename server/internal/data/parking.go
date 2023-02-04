package data

type parking struct {
	parkingSpaceArr []parkingSpace
	spaceMap        map[string]*parkingSpace
	count           int32
	temperature     int32
	humidity        int32
	weather         string
	info            string
}
