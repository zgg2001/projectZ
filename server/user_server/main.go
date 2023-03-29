package main

import (
	"log"

	"github.com/zgg2001/projectZ/server/user_server/internal/data"
	"github.com/zgg2001/projectZ/server/user_server/internal/operate"
	"github.com/zgg2001/projectZ/server/user_server/internal/transmission"
)

func main() {

	err := data.InitRedis()
	if err != nil {
		log.Println("Redis init error:", err)
		return
	}
	err = data.InitMySql()
	if err != nil {
		log.Println("Mysql init error:", err)
		return
	}

	// service init
	err = operate.ServerService.Init()
	if err != nil {
		log.Println("Server service init error:", err)
		return
	}

	// db task
	go operate.ServerService.DBMgrTaskQueueRunning()

	transmission.StartRPCService()
}
