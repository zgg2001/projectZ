package operate

import (
	"log"
	"strconv"

	"github.com/zgg2001/projectZ/server/user_server/internal/data"
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

var ServerService = &serverService{}

type ServiceOperation interface {
	Init()
}

type serverService struct {
	rpc.UnimplementedProjectServiceServer
	funcChan chan func()
}

func (ss *serverService) Init() error {
	ss.funcChan = make(chan func(), 22)
	return nil
}

func (ss *serverService) DBMgrTaskQueueRunning() {
	for {
		f, ok := <-ss.funcChan
		if !ok {
			break
		}
		f()
	}
}

func (ss *serverService) SqlRegisterUser(username, paasword string, nowTime int64) {
	ss.funcChan <- func() {
		// mysql
		var balance int32 = 0
		changedPasswd := data.GetMD5Hash(paasword)
		uid, err := data.InsertUserTbl(username, changedPasswd, balance, nowTime)
		if err != nil {
			log.Println(err)
		}
		// redis
		data.RedisAddUser(&data.UserRow{
			Id:           uid,
			Username:     username,
			Password:     changedPasswd,
			Balance:      balance,
			CreationTime: nowTime,
			LastModified: nowTime,
		})
	}
}

func (ss *serverService) SqlUserRecharge(uid, amount int32) int32 {
	// redis
	uidStr := strconv.Itoa(int(uid))
	ok, oldBalance := data.RedisGetBalanceByUid(uidStr)
	if ok {
		data.RedisSetBalance(uidStr, oldBalance+amount)
	}
	// mysql
	err := data.ChangeUserBalanceTbl(uidStr, oldBalance+amount)
	if err != nil {
		log.Println(err)
	}
	return 0
}

func (ss *serverService) SqlAddCar(uid int32, license string, nowTime int64) {
	// redis
	data.RedisAddLicense(&data.LicenseRow{
		License:     license,
		Id:          uid,
		CheckInTime: nowTime,
	})
	// mysql
	ss.funcChan <- func() {
		err := data.InsertLicenseTbl(uid, license, nowTime)
		if err != nil {
			log.Println(err)
		}
	}
}

func (ss *serverService) SqlDeleteCar(uid int32, license string) {
	// redis
	data.RedisDelLicense(uid, license)
	// mysql
	ss.funcChan <- func() {
		err := data.DeleteLicenseTbl(license)
		if err != nil {
			log.Println(err)
		}
	}
}

func (ss *serverService) SqlChangeCar(uid int32, license, newlicense string, nowTime int64) {
	// redis
	data.RedisUpdLicense(uid, license, newlicense, nowTime)
	// mysql
	ss.funcChan <- func() {
		err := data.ChangeLicenseTbl(license, newlicense, nowTime)
		if err != nil {
			log.Println(err)
		}
	}
}
