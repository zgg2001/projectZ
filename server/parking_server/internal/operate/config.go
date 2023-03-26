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
	funcChan chan func()
}

func (ss *serverService) DataInit() error {
	ss.funcChan = make(chan func(), 22)
	err := data.UserInit()
	if err != nil {
		return err
	}
	err = data.ParkingInit()
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
