package data

type userMgr struct {
	userArr    []user
	idMap      map[int32]*user
	licenseMap map[string]*user
}
