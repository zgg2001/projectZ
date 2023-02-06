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
	pMgr data.ParkingMgr
	uMgr data.UserMgr
}

func (ss *serverService) Init() error {
	err := ss.pMgr.Init()
	if err != nil {
		return err
	}
	err = ss.uMgr.Init()
	if err != nil {
		return err
	}
	return nil
}
