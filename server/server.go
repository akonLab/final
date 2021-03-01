package server

import (
	"com.grpc.tleu/calcpb"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct {
	calcpb.UnimplementedCalcServServer
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServServer(s, &Server{})

	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
func (s *Server) CalcPrime(ctx context.Context, req *calcpb.CalcPrimeRequest) (*calcpb.CalcPrimeResponse, error) {
	res := &calcpb.CalcPrimeResponse{
		Num: 1,
	}

	return res, nil
}
func (s *Server) CalcAvg(stream calcpb.CalcServ_CalcAvgServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err.Error())
			return err
		}

		err = stream.SendAndClose(&calcpb.CalcAvgResponse{Num: 3})
		if err != nil {
			log.Fatalf("error while sending to client: %v", err.Error())
			return err
		}
	}
}
