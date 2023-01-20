package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"github.com/zgg2001/projectZ/server/internal/operate"
	"github.com/zgg2001/projectZ/server/pkg/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	cert, err := tls.LoadX509KeyPair("./auth/server.pem", "./auth/server.key")
	if err != nil {
		log.Fatal("TLS error", err)
		return
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./auth/ca.crt")
	if err != nil {
		log.Fatal("NewCertPool error", err)
		return
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	rpc.RegisterProjectServiceServer(grpcServer, operate.CheckService)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}
