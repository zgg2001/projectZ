package transmission

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/zgg2001/projectZ/server/pkg/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RPCNewClient() rpc.ProjectServiceClient {

	cert, err := tls.LoadX509KeyPair(ClientPemPath, ClientKeyPath)
	if err != nil {
		log.Fatal("TLS error", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(CACrtPath)
	if err != nil {
		log.Fatal("NewCertPool error", err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "*.zgg2001.com",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(RPCServerIddr, grpc.WithTransportCredentials(creds))
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
