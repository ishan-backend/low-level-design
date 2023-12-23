package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	servicepb "../proto"
)

func main() {
	list, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	serviceServer := ServiceServer{}
	servicepb.RegisterStreamingGrpcServer(grpcServer, &serviceServer)

	grpcServer.Serve(list)
}
