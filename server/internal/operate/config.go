package operate

import (
	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

var ServerService = &serverService{}

type serverService struct {
	rpc.UnimplementedProjectServiceServer
}
