package data

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"sync"

	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

type UserLoginInfo struct {
	uid      int32
	password string
}

type UserMgr struct {
	userArr        []user
	idMap          map[int32]*user
	idMapLock      *sync.RWMutex
	licenseMap     map[string]*user
	licenseMapLock *sync.RWMutex
	loginMap       map[string]UserLoginInfo
	loginMapLock   *sync.RWMutex
}

func (um *UserMgr) Init(pm *ParkingMgr) error {

	log.Println("UserMgr init ...")

	um.idMap = make(map[int32]*user)
	um.idMapLock = new(sync.RWMutex)
	um.licenseMap = make(map[string]*user)
	um.licenseMapLock = new(sync.RWMutex)
	um.loginMap = make(map[string]UserLoginInfo)
	um.loginMapLock = new(sync.RWMutex)

	// read and load user
	userRet, err := ReadUserTbl()
	if err != nil {
		return err
	}
	for _, tempUser := range userRet {
		fmt.Println(tempUser)
		u := user{
			id:           tempUser.Id,
			balance:      tempUser.Balance,
			username:     tempUser.Username,
			creationTime: tempUser.CreationTime,
			lastModified: tempUser.LastModified,
			cars:         nil,
			carMap:       make(map[string]*car),
			carMapLock:   new(sync.RWMutex),
		}
		um.userArr = append(um.userArr, u)
		um.idMap[u.id] = &u
		um.loginMap[tempUser.Username] = UserLoginInfo{uid: tempUser.Id, password: tempUser.Password}
	}

	// read and load user license
	licenseRet, err := ReadLicenseTbl()
	if err != nil {
		return err
	}
	for _, tempLicense := range licenseRet {
		fmt.Println(tempLicense)
		i := tempLicense.Id
		c := car{
			license:         tempLicense.License,
			parkingPtr:      nil,
			parkingSpacePtr: nil,
			checkInTime:     tempLicense.CheckInTime,
			entryTime:       0,
		}
		u := um.idMap[i]
		u.cars = append(u.cars, c)
		u.carMap[c.license] = &c
		um.licenseMap[c.license] = u
	}

	// read and load record
	recordRet, err := ReadRecordTbl()
	if err != nil {
		return err
	}
	for _, tempRecord := range recordRet {
		fmt.Println(tempRecord)
		l := tempRecord.License
		pptr, sptr, err := pm.MgrGetParkingPtrPair(tempRecord.PId, tempRecord.SId)
		if err != nil {
			return err
		}
		tCar := um.licenseMap[l].carMap[l]
		tCar.SetParkingSpace(pptr, sptr, tempRecord.EntryTime)
	}

	return nil
}

func (um *UserMgr) GetUserByLicense(license string) (bool, *user) {
	um.licenseMapLock.RLock()
	defer um.licenseMapLock.RUnlock()
	if uptr, ok := um.licenseMap[license]; ok {
		return true, uptr
	}
	return false, nil
}

func (um *UserMgr) GetUserById(uid int32) (*user, error) {
	um.idMapLock.RLock()
	defer um.idMapLock.RUnlock()
	if uptr, ok := um.idMap[uid]; ok {
		return uptr, nil
	}
	return nil, ErrUserNotExist
}

func (um *UserMgr) LoginAuth(username, password string) (int32, rpc.LoginResult) {
	um.loginMapLock.RLock()
	defer um.loginMapLock.RUnlock()
	if userInfo, ok := um.loginMap[username]; ok {
		changedPasswd := GetMD5Hash(password)
		if userInfo.password == changedPasswd {
			return userInfo.uid, rpc.LoginResult_LOGIN_SUCCESS
		}
		return -1, rpc.LoginResult_LOGIN_FAIL_WRONG_PASSWORD
	}
	return -1, rpc.LoginResult_LOGIN_FAIL_NOT_EXIST
}

func (um *UserMgr) RegistrationAuth(username, password string) rpc.RegistrationResult {
	um.loginMapLock.RLock()
	defer um.loginMapLock.RUnlock()
	if _, ok := um.loginMap[username]; ok {
		return rpc.RegistrationResult_REGISTRATION_FAIL_ALREADY_EXIST
	}
	return rpc.RegistrationResult_REGISTRATION_SUCCESS
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
