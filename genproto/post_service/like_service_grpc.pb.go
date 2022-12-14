// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: like_service.proto

package post_service

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

// LikeServiceClient is the client API for LikeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeServiceClient interface {
	CreateOrUpdateLike(ctx context.Context, in *CreateOrUpdateLikeRequest, opts ...grpc.CallOption) (*Pustoy, error)
	GetLike(ctx context.Context, in *Get, opts ...grpc.CallOption) (*CreateOrUpdateLikeRequest, error)
	GetAllLik(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type likeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeServiceClient(cc grpc.ClientConnInterface) LikeServiceClient {
	return &likeServiceClient{cc}
}

func (c *likeServiceClient) CreateOrUpdateLike(ctx context.Context, in *CreateOrUpdateLikeRequest, opts ...grpc.CallOption) (*Pustoy, error) {
	out := new(Pustoy)
	err := c.cc.Invoke(ctx, "/genproto.LikeService/CreateOrUpdateLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) GetLike(ctx context.Context, in *Get, opts ...grpc.CallOption) (*CreateOrUpdateLikeRequest, error) {
	out := new(CreateOrUpdateLikeRequest)
	err := c.cc.Invoke(ctx, "/genproto.LikeService/GetLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) GetAllLik(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/genproto.LikeService/GetAllLik", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServiceServer is the server API for LikeService service.
// All implementations must embed UnimplementedLikeServiceServer
// for forward compatibility
type LikeServiceServer interface {
	CreateOrUpdateLike(context.Context, *CreateOrUpdateLikeRequest) (*Pustoy, error)
	GetLike(context.Context, *Get) (*CreateOrUpdateLikeRequest, error)
	GetAllLik(context.Context, *GetAllRequest) (*GetAllResponse, error)
	mustEmbedUnimplementedLikeServiceServer()
}

// UnimplementedLikeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServiceServer struct {
}

func (UnimplementedLikeServiceServer) CreateOrUpdateLike(context.Context, *CreateOrUpdateLikeRequest) (*Pustoy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateLike not implemented")
}
func (UnimplementedLikeServiceServer) GetLike(context.Context, *Get) (*CreateOrUpdateLikeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLike not implemented")
}
func (UnimplementedLikeServiceServer) GetAllLik(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLik not implemented")
}
func (UnimplementedLikeServiceServer) mustEmbedUnimplementedLikeServiceServer() {}

// UnsafeLikeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServiceServer will
// result in compilation errors.
type UnsafeLikeServiceServer interface {
	mustEmbedUnimplementedLikeServiceServer()
}

func RegisterLikeServiceServer(s grpc.ServiceRegistrar, srv LikeServiceServer) {
	s.RegisterService(&LikeService_ServiceDesc, srv)
}

func _LikeService_CreateOrUpdateLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).CreateOrUpdateLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.LikeService/CreateOrUpdateLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).CreateOrUpdateLike(ctx, req.(*CreateOrUpdateLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_GetLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Get)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).GetLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.LikeService/GetLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).GetLike(ctx, req.(*Get))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_GetAllLik_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).GetAllLik(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.LikeService/GetAllLik",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).GetAllLik(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LikeService_ServiceDesc is the grpc.ServiceDesc for LikeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LikeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.LikeService",
	HandlerType: (*LikeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdateLike",
			Handler:    _LikeService_CreateOrUpdateLike_Handler,
		},
		{
			MethodName: "GetLike",
			Handler:    _LikeService_GetLike_Handler,
		},
		{
			MethodName: "GetAllLik",
			Handler:    _LikeService_GetAllLik_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "like_service.proto",
}
