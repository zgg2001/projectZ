package transmission

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"github.com/zgg2001/projectZ/server/parking_server/internal/operate"
	"github.com/zgg2001/projectZ/server/parking_server/pkg/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func StartRPCService() {

	log.Println("Start the rpc service ...")

	cert, err := tls.LoadX509KeyPair(ServerPemPath, ServerKeyPath)
	if err != nil {
		log.Fatal("TLS error:", err)
		return
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(CACrtPath)
	if err != nil {
		log.Fatal("NewCertPool error:", err)
		return
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	listen, err := net.Listen(TLS, RPCServerIddr)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	rpc.RegisterProjectServiceServer(grpcServer, operate.ServerService)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
