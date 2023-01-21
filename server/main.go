package main

import (
	"github.com/zgg2001/projectZ/server/internal/transmission"
)

func main() {
	// 启动rpc服务
	transmission.StartRPCService()
}
