package main

import (
	"log"
	"net"

	"github.com/zgg2001/projectZ/server/internal/operate"
	"github.com/zgg2001/projectZ/server/pkg/rpc"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println(err)
	}
	grpcServer := grpc.NewServer()
	rpc.RegisterProjectServiceServer(grpcServer, operate.CheckService)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
