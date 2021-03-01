package server

import (
	"com.grpc.tleu/calcpb"
	"context"
	"google.golang.org/grpc"
)

type Server struct {
	calcpb.UnimplementedCalcServServer
}

func (s *Server) CalcPrime(ctx context.Context, in *calcpb.CalcServ_CalcPrimeServer, opts ...grpc.CallOption) (calcpb.CalcServ_CalcPrimeClient, error) {
	return nil
}
func (s *Server) CalcAvg(ctx context.Context, opts ...grpc.CallOption) (calcpb.CalcServ_CalcAvgClient, error) {
	return nil
}
