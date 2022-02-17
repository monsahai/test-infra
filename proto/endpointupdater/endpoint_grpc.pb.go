// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package endpointupdater

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

// TestUpdaterClient is the client API for TestUpdater service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestUpdaterClient interface {
	// Sends an update
	UpdateTest(ctx context.Context, in *TestUpdateRequest, opts ...grpc.CallOption) (*TestUpdateReply, error)
	QuitTestUpdateServer(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
}

type testUpdaterClient struct {
	cc grpc.ClientConnInterface
}

func NewTestUpdaterClient(cc grpc.ClientConnInterface) TestUpdaterClient {
	return &testUpdaterClient{cc}
}

func (c *testUpdaterClient) UpdateTest(ctx context.Context, in *TestUpdateRequest, opts ...grpc.CallOption) (*TestUpdateReply, error) {
	out := new(TestUpdateReply)
	err := c.cc.Invoke(ctx, "/endpointupdater.TestUpdater/UpdateTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testUpdaterClient) QuitTestUpdateServer(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/endpointupdater.TestUpdater/QuitTestUpdateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestUpdaterServer is the server API for TestUpdater service.
// All implementations must embed UnimplementedTestUpdaterServer
// for forward compatibility
type TestUpdaterServer interface {
	// Sends an update
	UpdateTest(context.Context, *TestUpdateRequest) (*TestUpdateReply, error)
	QuitTestUpdateServer(context.Context, *Void) (*Void, error)
	mustEmbedUnimplementedTestUpdaterServer()
}

// UnimplementedTestUpdaterServer must be embedded to have forward compatible implementations.
type UnimplementedTestUpdaterServer struct {
}

func (UnimplementedTestUpdaterServer) UpdateTest(context.Context, *TestUpdateRequest) (*TestUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTest not implemented")
}
func (UnimplementedTestUpdaterServer) QuitTestUpdateServer(context.Context, *Void) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuitTestUpdateServer not implemented")
}
func (UnimplementedTestUpdaterServer) mustEmbedUnimplementedTestUpdaterServer() {}

// UnsafeTestUpdaterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestUpdaterServer will
// result in compilation errors.
type UnsafeTestUpdaterServer interface {
	mustEmbedUnimplementedTestUpdaterServer()
}

func RegisterTestUpdaterServer(s grpc.ServiceRegistrar, srv TestUpdaterServer) {
	s.RegisterService(&TestUpdater_ServiceDesc, srv)
}

func _TestUpdater_UpdateTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUpdaterServer).UpdateTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpointupdater.TestUpdater/UpdateTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUpdaterServer).UpdateTest(ctx, req.(*TestUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestUpdater_QuitTestUpdateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestUpdaterServer).QuitTestUpdateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpointupdater.TestUpdater/QuitTestUpdateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestUpdaterServer).QuitTestUpdateServer(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// TestUpdater_ServiceDesc is the grpc.ServiceDesc for TestUpdater service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestUpdater_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpointupdater.TestUpdater",
	HandlerType: (*TestUpdaterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateTest",
			Handler:    _TestUpdater_UpdateTest_Handler,
		},
		{
			MethodName: "QuitTestUpdateServer",
			Handler:    _TestUpdater_QuitTestUpdateServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "endpoint.proto",
}
