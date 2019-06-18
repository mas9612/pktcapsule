package main

import (
	"log"
	"net"

	pb "github.com/mas9612/pktcapsule"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if err := run(logger); err != nil {
		logger.Error("gRPC server error", zap.Error(err))
	}
}

func run(logger *zap.Logger) error {
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPktCapsuleServer(grpcServer, &pktcapsuleServer{logger: logger})
	logger.Info("listening on :10000")
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
