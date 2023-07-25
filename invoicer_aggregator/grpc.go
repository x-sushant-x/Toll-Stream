package main

import (
	"context"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

type GRPCAggregatorServer struct {
	types.UnimplementedDistanceAggregatorServer
	svc InvoiceAggregator
}

func NewGRPCServer(svc InvoiceAggregator) *GRPCAggregatorServer {
	return &GRPCAggregatorServer{
		svc: svc,
	}
}

func (s *GRPCAggregatorServer) AggregateDistance(ctx context.Context, req *types.DistanceAggregateRequest) (*types.None, error) {
	distance := types.CalculatedDistance{
		OBUID:     int(req.ObuID),
		Distance:  req.Value,
		Timestamp: req.Unix,
	}

	s.svc.AggregateDistance(distance)
	return &types.None{}, nil
}
