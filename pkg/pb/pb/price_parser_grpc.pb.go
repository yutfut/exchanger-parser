// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: price_parser.proto

package exchanger_parser_pb

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

const (
	PriceParserService_GetRate_FullMethodName = "/price_parser.PriceParserService/GetRate"
)

// PriceParserServiceClient is the client API for PriceParserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PriceParserServiceClient interface {
	GetRate(ctx context.Context, in *GetRateRequest, opts ...grpc.CallOption) (*GetRateResponse, error)
}

type priceParserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceParserServiceClient(cc grpc.ClientConnInterface) PriceParserServiceClient {
	return &priceParserServiceClient{cc}
}

func (c *priceParserServiceClient) GetRate(ctx context.Context, in *GetRateRequest, opts ...grpc.CallOption) (*GetRateResponse, error) {
	out := new(GetRateResponse)
	err := c.cc.Invoke(ctx, PriceParserService_GetRate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceParserServiceServer is the server API for PriceParserService service.
// All implementations must embed UnimplementedPriceParserServiceServer
// for forward compatibility
type PriceParserServiceServer interface {
	GetRate(context.Context, *GetRateRequest) (*GetRateResponse, error)
	mustEmbedUnimplementedPriceParserServiceServer()
}

// UnimplementedPriceParserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPriceParserServiceServer struct {
}

func (UnimplementedPriceParserServiceServer) GetRate(context.Context, *GetRateRequest) (*GetRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}
func (UnimplementedPriceParserServiceServer) mustEmbedUnimplementedPriceParserServiceServer() {}

// UnsafePriceParserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PriceParserServiceServer will
// result in compilation errors.
type UnsafePriceParserServiceServer interface {
	mustEmbedUnimplementedPriceParserServiceServer()
}

func RegisterPriceParserServiceServer(s grpc.ServiceRegistrar, srv PriceParserServiceServer) {
	s.RegisterService(&PriceParserService_ServiceDesc, srv)
}

func _PriceParserService_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceParserServiceServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceParserService_GetRate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceParserServiceServer).GetRate(ctx, req.(*GetRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PriceParserService_ServiceDesc is the grpc.ServiceDesc for PriceParserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PriceParserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "price_parser.PriceParserService",
	HandlerType: (*PriceParserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _PriceParserService_GetRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "price_parser.proto",
}
