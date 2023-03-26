package operate

import (
	"github.com/zgg2001/projectZ/server/parking_server/internal/data"
	"github.com/zgg2001/projectZ/server/parking_server/pkg/rpc"
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
