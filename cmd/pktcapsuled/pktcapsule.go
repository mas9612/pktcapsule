package main

import (
	"log"
	"net"

	pb "github.com/mas9612/pktcapsule"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPktCapsuleServer(grpcServer, &pktcapsuleServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
