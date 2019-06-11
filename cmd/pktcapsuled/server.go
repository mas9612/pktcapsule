package main

import (
	"context"

	pb "github.com/mas9612/pktcapsule"
)

type pktcapsuleServer struct{}

// Encapsulate adds newly IP header to the given packet data.
func (s *pktcapsuleServer) Encapsulate(ctx context.Context, req *pb.EncapsulateRequest) (*pb.Packet, error) {
	return &pb.Packet{}, nil
}

// Decapsulate removes outer IP header from the given packet data.
func (s *pktcapsuleServer) Decapsulate(ctx context.Context, req *pb.DecapsulateRequest) (*pb.Packet, error) {
	return &pb.Packet{}, nil
}
