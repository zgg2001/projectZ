package main

import (
	"log"

	"github.com/zgg2001/projectZ/server/internal/operate"
	"github.com/zgg2001/projectZ/server/internal/transmission"
)

func main() {

	err := transmission.InitDB()
	if err != nil {
		log.Println("Database init error:", err)
		return
	}

	// service init
	operate.ServerService.Init()

	transmission.StartRPCService()
}
