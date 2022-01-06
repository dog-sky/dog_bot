// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package interact

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

// DogClient is the client API for Dog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DogClient interface {
	SetStatus(ctx context.Context, in *SetStatusRequest, opts ...grpc.CallOption) (*SetStatusReply, error)
	StatusList(ctx context.Context, in *StatusListrequest, opts ...grpc.CallOption) (*StatusListReply, error)
}

type dogClient struct {
	cc grpc.ClientConnInterface
}

func NewDogClient(cc grpc.ClientConnInterface) DogClient {
	return &dogClient{cc}
}

func (c *dogClient) SetStatus(ctx context.Context, in *SetStatusRequest, opts ...grpc.CallOption) (*SetStatusReply, error) {
	out := new(SetStatusReply)
	err := c.cc.Invoke(ctx, "/interact.Dog/SetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dogClient) StatusList(ctx context.Context, in *StatusListrequest, opts ...grpc.CallOption) (*StatusListReply, error) {
	out := new(StatusListReply)
	err := c.cc.Invoke(ctx, "/interact.Dog/StatusList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DogServer is the server API for Dog service.
// All implementations must embed UnimplementedDogServer
// for forward compatibility
type DogServer interface {
	SetStatus(context.Context, *SetStatusRequest) (*SetStatusReply, error)
	StatusList(context.Context, *StatusListrequest) (*StatusListReply, error)
	mustEmbedUnimplementedDogServer()
}

// UnimplementedDogServer must be embedded to have forward compatible implementations.
type UnimplementedDogServer struct {
}

func (UnimplementedDogServer) SetStatus(context.Context, *SetStatusRequest) (*SetStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatus not implemented")
}
func (UnimplementedDogServer) StatusList(context.Context, *StatusListrequest) (*StatusListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusList not implemented")
}
func (UnimplementedDogServer) mustEmbedUnimplementedDogServer() {}

// UnsafeDogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DogServer will
// result in compilation errors.
type UnsafeDogServer interface {
	mustEmbedUnimplementedDogServer()
}

func RegisterDogServer(s grpc.ServiceRegistrar, srv DogServer) {
	s.RegisterService(&Dog_ServiceDesc, srv)
}

func _Dog_SetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DogServer).SetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.Dog/SetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DogServer).SetStatus(ctx, req.(*SetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dog_StatusList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusListrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DogServer).StatusList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.Dog/StatusList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DogServer).StatusList(ctx, req.(*StatusListrequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dog_ServiceDesc is the grpc.ServiceDesc for Dog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interact.Dog",
	HandlerType: (*DogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetStatus",
			Handler:    _Dog_SetStatus_Handler,
		},
		{
			MethodName: "StatusList",
			Handler:    _Dog_StatusList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/dog.proto",
}
