package transmission

import (
	"log"
	"net"

	"github.com/zgg2001/projectZ/server/user_server/internal/operate"
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"

	"google.golang.org/grpc"
)

func StartRPCService() {

	log.Println("Start the rpc service ...")

	listen, err := net.Listen(TLS, RPCServerIddr)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	rpc.RegisterProjectServiceServer(grpcServer, operate.ServerService)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
