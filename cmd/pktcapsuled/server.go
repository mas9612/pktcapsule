package main

import (
	"context"

	pb "github.com/mas9612/pktcapsule"
	"github.com/mas9612/pktcapsule/ipinip"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type pktcapsuleServer struct{}

// Encapsulate adds newly IP header to the given packet data.
func (s *pktcapsuleServer) Encapsulate(ctx context.Context, req *pb.EncapsulateRequest) (*pb.Packet, error) {
	b, err := ipinip.Encapsulate(req.Data, ipinip.IntToIP(req.SrcIp), ipinip.IntToIP(req.DstIp))
	if err != nil {
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
		return nil, status.Error(codes.Internal, "internal server error")
	}
	return &pb.Packet{
		Data: b,
	}, nil
}
