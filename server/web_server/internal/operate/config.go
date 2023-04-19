package operate

import (
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
	"google.golang.org/grpc"
)

var (
	RPCConn   *grpc.ClientConn
	RPCClient rpc.ProjectServiceClient
)
