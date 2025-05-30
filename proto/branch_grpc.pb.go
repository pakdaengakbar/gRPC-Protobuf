// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/branch.proto

package branch

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BranchService_CreateBranch_FullMethodName = "/grpcserver.proto.BranchService/CreateBranch"
	BranchService_GetBranch_FullMethodName    = "/grpcserver.proto.BranchService/GetBranch"
)

// BranchServiceClient is the client API for BranchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BranchServiceClient interface {
	CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error)
	GetBranch(ctx context.Context, in *GetBranchRequest, opts ...grpc.CallOption) (*GetBranchResponse, error)
}

type branchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBranchServiceClient(cc grpc.ClientConnInterface) BranchServiceClient {
	return &branchServiceClient{cc}
}

func (c *branchServiceClient) CreateBranch(ctx context.Context, in *CreateBranchRequest, opts ...grpc.CallOption) (*CreateBranchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBranchResponse)
	err := c.cc.Invoke(ctx, BranchService_CreateBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) GetBranch(ctx context.Context, in *GetBranchRequest, opts ...grpc.CallOption) (*GetBranchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBranchResponse)
	err := c.cc.Invoke(ctx, BranchService_GetBranch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BranchServiceServer is the server API for BranchService service.
// All implementations must embed UnimplementedBranchServiceServer
// for forward compatibility.
type BranchServiceServer interface {
	CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error)
	GetBranch(context.Context, *GetBranchRequest) (*GetBranchResponse, error)
	mustEmbedUnimplementedBranchServiceServer()
}

// UnimplementedBranchServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBranchServiceServer struct{}

func (UnimplementedBranchServiceServer) CreateBranch(context.Context, *CreateBranchRequest) (*CreateBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBranch not implemented")
}
func (UnimplementedBranchServiceServer) GetBranch(context.Context, *GetBranchRequest) (*GetBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranch not implemented")
}
func (UnimplementedBranchServiceServer) mustEmbedUnimplementedBranchServiceServer() {}
func (UnimplementedBranchServiceServer) testEmbeddedByValue()                       {}

// UnsafeBranchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BranchServiceServer will
// result in compilation errors.
type UnsafeBranchServiceServer interface {
	mustEmbedUnimplementedBranchServiceServer()
}

func RegisterBranchServiceServer(s grpc.ServiceRegistrar, srv BranchServiceServer) {
	// If the following call pancis, it indicates UnimplementedBranchServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BranchService_ServiceDesc, srv)
}

func _BranchService_CreateBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).CreateBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_CreateBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).CreateBranch(ctx, req.(*CreateBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_GetBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).GetBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BranchService_GetBranch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).GetBranch(ctx, req.(*GetBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BranchService_ServiceDesc is the grpc.ServiceDesc for BranchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BranchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcserver.proto.BranchService",
	HandlerType: (*BranchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBranch",
			Handler:    _BranchService_CreateBranch_Handler,
		},
		{
			MethodName: "GetBranch",
			Handler:    _BranchService_GetBranch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/branch.proto",
}
