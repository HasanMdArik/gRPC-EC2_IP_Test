package main

import (
	"ec2-grpc-ip-test/server/IPTestService"
	"log"
	"net"

	"google.golang.org/grpc"
)

func panicOnErr(err error) {
	if err != nil {
		log.Println("Error:", err.Error())
		panic(err)
	}
}

func main() {
	log.Println("IP test server")
	// Listener for server
	listener, err := net.Listen("tcp", ":2500")
	panicOnErr(err)
	// Initializing the handler
	serverHandler := &IPTestServiceHandler{}

	// Initializing the gRPC server
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	IPTestService.RegisterIP_TestServer(grpcServer, serverHandler)
	// Serving the server
	err = grpcServer.Serve(listener)
	panicOnErr(err)
}
