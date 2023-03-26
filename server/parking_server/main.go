package main

import (
	"log"

	"github.com/zgg2001/projectZ/server/parking_server/internal/data"
	"github.com/zgg2001/projectZ/server/parking_server/internal/operate"
	"github.com/zgg2001/projectZ/server/parking_server/internal/transmission"
)

func main() {

	err := data.InitDB()
	if err != nil {
		log.Println("Database init error:", err)
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
