package server

import (
	"context"
	"google.golang.org/grpc"
)

type GRPCServer struct {
}

func (s *GRPCServer) CalcPrime(ctx context.Context, in *CalcPrimeRequest, opts ...grpc.CallOption) (CalcServ_CalcPrimeClient, error) {
	return &in.CalcPrimeResponse{}
}
func (s *GRPCServer) CalcAvg(ctx context.Context, opts ...grpc.CallOption) (CalcServ_CalcAvgClient, error)
