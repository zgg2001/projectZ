package operate

import (
	"log"
	"time"

	"github.com/zgg2001/projectZ/server/internal/data"
	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

var ServerService = &serverService{}

type ServiceOperation interface {
	Init()
}

type serverService struct {
	rpc.UnimplementedProjectServiceServer
	pMgr     data.ParkingMgr
	uMgr     data.UserMgr
	funcChan chan func()
}

func (ss *serverService) Init() error {
	ss.funcChan = make(chan func(), 22)
	err := ss.pMgr.Init()
	if err != nil {
		return err
	}
	err = ss.uMgr.Init(&ss.pMgr)
	if err != nil {
		return err
	}
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

func (ss *serverService) RegisterUser(username, paasword string) {
	ss.funcChan <- func() {
		var balance int32 = 0
		nowTime := time.Now().Unix()
		changedPasswd := data.GetMD5Hash(paasword)
		uid, err := data.InsertUserTbl(username, changedPasswd, balance, nowTime)
		if err != nil {
			log.Println(err)
		}
		ss.uMgr.AddUser(username, paasword, uid, balance, nowTime)
	}
}
