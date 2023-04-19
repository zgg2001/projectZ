package transmission

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateRPCClientConn() (*grpc.ClientConn, error) {
	log.Println("Create grpc client conn ...")
	var err error
	RPCConn, err := RPCNewClient()
	if err != nil {
		return nil, err
	}
	return RPCConn, nil
}

func RPCNewClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(RPCServerIddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
