// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: hotelpb/hotel.proto

package hotelpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HotelServiceClient is the client API for HotelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelServiceClient interface {
	RegisterHotel(ctx context.Context, in *RegisterHotelRequest, opts ...grpc.CallOption) (*Hotel, error)
	GetHotel(ctx context.Context, in *GetHotelRequest, opts ...grpc.CallOption) (*Hotel, error)
	UpdateHotel(ctx context.Context, in *Hotel, opts ...grpc.CallOption) (*Hotel, error)
	DeleteHotel(ctx context.Context, in *DeleteHotelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetHotels(ctx context.Context, in *GetHotelsRequest, opts ...grpc.CallOption) (*Hotels, error)
}

type hotelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelServiceClient(cc grpc.ClientConnInterface) HotelServiceClient {
	return &hotelServiceClient{cc}
}

func (c *hotelServiceClient) RegisterHotel(ctx context.Context, in *RegisterHotelRequest, opts ...grpc.CallOption) (*Hotel, error) {
	out := new(Hotel)
	err := c.cc.Invoke(ctx, "/hotelpb.HotelService/RegisterHotel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetHotel(ctx context.Context, in *GetHotelRequest, opts ...grpc.CallOption) (*Hotel, error) {
	out := new(Hotel)
	err := c.cc.Invoke(ctx, "/hotelpb.HotelService/GetHotel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) UpdateHotel(ctx context.Context, in *Hotel, opts ...grpc.CallOption) (*Hotel, error) {
	out := new(Hotel)
	err := c.cc.Invoke(ctx, "/hotelpb.HotelService/UpdateHotel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) DeleteHotel(ctx context.Context, in *DeleteHotelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hotelpb.HotelService/DeleteHotel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetHotels(ctx context.Context, in *GetHotelsRequest, opts ...grpc.CallOption) (*Hotels, error) {
	out := new(Hotels)
	err := c.cc.Invoke(ctx, "/hotelpb.HotelService/GetHotels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelServiceServer is the server API for HotelService service.
// All implementations must embed UnimplementedHotelServiceServer
// for forward compatibility
type HotelServiceServer interface {
	RegisterHotel(context.Context, *RegisterHotelRequest) (*Hotel, error)
	GetHotel(context.Context, *GetHotelRequest) (*Hotel, error)
	UpdateHotel(context.Context, *Hotel) (*Hotel, error)
	DeleteHotel(context.Context, *DeleteHotelRequest) (*emptypb.Empty, error)
	GetHotels(context.Context, *GetHotelsRequest) (*Hotels, error)
	mustEmbedUnimplementedHotelServiceServer()
}

// UnimplementedHotelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHotelServiceServer struct {
}

func (UnimplementedHotelServiceServer) RegisterHotel(context.Context, *RegisterHotelRequest) (*Hotel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterHotel not implemented")
}
func (UnimplementedHotelServiceServer) GetHotel(context.Context, *GetHotelRequest) (*Hotel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotel not implemented")
}
func (UnimplementedHotelServiceServer) UpdateHotel(context.Context, *Hotel) (*Hotel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHotel not implemented")
}
func (UnimplementedHotelServiceServer) DeleteHotel(context.Context, *DeleteHotelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHotel not implemented")
}
func (UnimplementedHotelServiceServer) GetHotels(context.Context, *GetHotelsRequest) (*Hotels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotels not implemented")
}
func (UnimplementedHotelServiceServer) mustEmbedUnimplementedHotelServiceServer() {}

// UnsafeHotelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelServiceServer will
// result in compilation errors.
type UnsafeHotelServiceServer interface {
	mustEmbedUnimplementedHotelServiceServer()
}

func RegisterHotelServiceServer(s grpc.ServiceRegistrar, srv HotelServiceServer) {
	s.RegisterService(&HotelService_ServiceDesc, srv)
}

func _HotelService_RegisterHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterHotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).RegisterHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotelpb.HotelService/RegisterHotel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).RegisterHotel(ctx, req.(*RegisterHotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotelpb.HotelService/GetHotel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetHotel(ctx, req.(*GetHotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_UpdateHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hotel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).UpdateHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotelpb.HotelService/UpdateHotel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).UpdateHotel(ctx, req.(*Hotel))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_DeleteHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).DeleteHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotelpb.HotelService/DeleteHotel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).DeleteHotel(ctx, req.(*DeleteHotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetHotels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetHotels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotelpb.HotelService/GetHotels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetHotels(ctx, req.(*GetHotelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HotelService_ServiceDesc is the grpc.ServiceDesc for HotelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HotelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hotelpb.HotelService",
	HandlerType: (*HotelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterHotel",
			Handler:    _HotelService_RegisterHotel_Handler,
		},
		{
			MethodName: "GetHotel",
			Handler:    _HotelService_GetHotel_Handler,
		},
		{
			MethodName: "UpdateHotel",
			Handler:    _HotelService_UpdateHotel_Handler,
		},
		{
			MethodName: "DeleteHotel",
			Handler:    _HotelService_DeleteHotel_Handler,
		},
		{
			MethodName: "GetHotels",
			Handler:    _HotelService_GetHotels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hotelpb/hotel.proto",
}
