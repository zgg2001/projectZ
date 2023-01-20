package transmission

import (
	"context"
	"log"

	"github.com/zgg2001/projectZ/server/pkg/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RPCNewClient() rpc.ProjectServiceClient {

	conn, err := grpc.Dial(RPCServerIddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Connect error", err)
	}

	defer conn.Close()

	client := rpc.NewProjectServiceClient(conn)

	request := &rpc.LPCheckRequest{
		License: "123",
	}
	resp, err := client.LicencePlateCheck(context.Background(), request)
	if err != nil {
		log.Fatal("Get error", err)
	}
	log.Println("Get success", resp.Result)

	return client
}
