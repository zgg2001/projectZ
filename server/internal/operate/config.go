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

func (ss *serverService) Init() {
	ss.pMgr.Init()
	ss.uMgr.Init()
}
