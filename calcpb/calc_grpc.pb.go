// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalcServClient is the client API for CalcServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalcServClient interface {
	CalcPrime(ctx context.Context, in *CalcPrimeRequest, opts ...grpc.CallOption) (CalcServ_CalcPrimeClient, error)
	CalcAvg(ctx context.Context, opts ...grpc.CallOption) (CalcServ_CalcAvgClient, error)
}

type calcServClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServClient(cc grpc.ClientConnInterface) CalcServClient {
	return &calcServClient{cc}
}

func (c *calcServClient) CalcPrime(ctx context.Context, in *CalcPrimeRequest, opts ...grpc.CallOption) (CalcServ_CalcPrimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcServ_ServiceDesc.Streams[0], "/calcpb.CalcServ/CalcPrime", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServCalcPrimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcServ_CalcPrimeClient interface {
	Recv() (*CalcPrimeResponse, error)
	grpc.ClientStream
}

type calcServCalcPrimeClient struct {
	grpc.ClientStream
}

func (x *calcServCalcPrimeClient) Recv() (*CalcPrimeResponse, error) {
	m := new(CalcPrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServClient) CalcAvg(ctx context.Context, opts ...grpc.CallOption) (CalcServ_CalcAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalcServ_ServiceDesc.Streams[1], "/calcpb.CalcServ/CalcAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServCalcAvgClient{stream}
	return x, nil
}

type CalcServ_CalcAvgClient interface {
	Send(*CalcAvgRequest) error
	CloseAndRecv() (*CalcAvgResponse, error)
	grpc.ClientStream
}

type calcServCalcAvgClient struct {
	grpc.ClientStream
}

func (x *calcServCalcAvgClient) Send(m *CalcAvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServCalcAvgClient) CloseAndRecv() (*CalcAvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalcAvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServServer is the server API for CalcServ service.
// All implementations must embed UnimplementedCalcServServer
// for forward compatibility
type CalcServServer interface {
	CalcPrime(*CalcPrimeRequest, CalcServ_CalcPrimeServer) error
	CalcAvg(CalcServ_CalcAvgServer) error
	mustEmbedUnimplementedCalcServServer()
}

// UnimplementedCalcServServer must be embedded to have forward compatible implementations.
type UnimplementedCalcServServer struct {
}

func (UnimplementedCalcServServer) CalcPrime(*CalcPrimeRequest, CalcServ_CalcPrimeServer) error {
	return status.Errorf(codes.Unimplemented, "method CalcPrime not implemented")
}
func (UnimplementedCalcServServer) CalcAvg(CalcServ_CalcAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method CalcAvg not implemented")
}
func (UnimplementedCalcServServer) mustEmbedUnimplementedCalcServServer() {}

// UnsafeCalcServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalcServServer will
// result in compilation errors.
type UnsafeCalcServServer interface {
	mustEmbedUnimplementedCalcServServer()
}

func RegisterCalcServServer(s grpc.ServiceRegistrar, srv CalcServServer) {
	s.RegisterService(&CalcServ_ServiceDesc, srv)
}

func _CalcServ_CalcPrime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CalcPrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServServer).CalcPrime(m, &calcServCalcPrimeServer{stream})
}

type CalcServ_CalcPrimeServer interface {
	Send(*CalcPrimeResponse) error
	grpc.ServerStream
}

type calcServCalcPrimeServer struct {
	grpc.ServerStream
}

func (x *calcServCalcPrimeServer) Send(m *CalcPrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalcServ_CalcAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServServer).CalcAvg(&calcServCalcAvgServer{stream})
}

type CalcServ_CalcAvgServer interface {
	SendAndClose(*CalcAvgResponse) error
	Recv() (*CalcAvgRequest, error)
	grpc.ServerStream
}

type calcServCalcAvgServer struct {
	grpc.ServerStream
}

func (x *calcServCalcAvgServer) SendAndClose(m *CalcAvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServCalcAvgServer) Recv() (*CalcAvgRequest, error) {
	m := new(CalcAvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServ_ServiceDesc is the grpc.ServiceDesc for CalcServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalcServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calcpb.CalcServ",
	HandlerType: (*CalcServServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalcPrime",
			Handler:       _CalcServ_CalcPrime_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CalcAvg",
			Handler:       _CalcServ_CalcAvg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calc.proto",
}
