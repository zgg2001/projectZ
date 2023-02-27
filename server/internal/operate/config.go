package operate

import (
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
	ss.funcChan = make(chan func(), 10)
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
		//注册用户
	}
}
