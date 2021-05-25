package main

import (
	"log"
	"net"
	"os"

	"github.com/sideroff/protocol-buffer-101/protocgen/currency"
	"github.com/sideroff/protocol-buffer-101/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {

	l := log.New(os.Stdout, "test ", log.LstdFlags)

	currencyService := server.NewCurrency(l)

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	currency.RegisterCurrencyServer(grpcServer, currencyService)

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen to port, %s", err)
		os.Exit(1)
	}


	grpcServer.Serve(listener)
}