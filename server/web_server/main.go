package main

import (
	"log"

	"github.com/zgg2001/projectZ/server/web_server/internal/transmission"
)

func main() {
	conn, err := transmission.CreateRPCClientConn()
	if err != nil {
		log.Println("RPC clientConn create error:", err)
		return
	}
	err = transmission.StartHttpServer(conn)
	if err != nil {
		log.Println("Http server start error:", err)
		return
	}
}
