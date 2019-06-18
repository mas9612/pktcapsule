package main

import (
	"context"

	pb "github.com/mas9612/pktcapsule"
	"github.com/mas9612/pktcapsule/ipinip"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type pktcapsuleServer struct {
	logger *zap.Logger
}

// Encapsulate adds newly IP header to the given packet data.
func (s *pktcapsuleServer) Encapsulate(ctx context.Context, req *pb.EncapsulateRequest) (*pb.Packet, error) {
	b, err := ipinip.Encapsulate(req.Data, ipinip.IntToIP(req.SrcIp), ipinip.IntToIP(req.DstIp))
	if err != nil {
		s.logger.Error("Encapsulate error", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &pb.Packet{
		Data: b,
	}, nil
}

// Decapsulate removes outer IP header from the given packet data.
func (s *pktcapsuleServer) Decapsulate(ctx context.Context, req *pb.DecapsulateRequest) (*pb.Packet, error) {
	b, err := ipinip.Decapsulate(req.Data)
	if err != nil {
		s.logger.Error("Decapsulate error", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &pb.Packet{
		Data: b,
	}, nil
}
