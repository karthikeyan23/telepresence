// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package connector

import (
	context "context"
	rpc "github.com/datawire/telepresence2/pkg/rpc"
	version "github.com/datawire/telepresence2/pkg/rpc/version"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ConnectorClient is the client API for Connector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorClient interface {
	// Returns version information from the Connector
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*version.VersionInfo, error)
	// Returns the current connectivity status
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConnectorStatus, error)
	// Connects the daemon to a cluster
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectInfo, error)
	// Adds a deployment intercept
	CreateIntercept(ctx context.Context, in *rpc.CreateInterceptRequest, opts ...grpc.CallOption) (*InterceptResult, error)
	// Deactivates and removes an existent deployment intercept.
	RemoveIntercept(ctx context.Context, in *rpc.RemoveInterceptRequest2, opts ...grpc.CallOption) (*InterceptResult, error)
	// Returns a list of deployments available for intercept.
	AvailableIntercepts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.AgentInfoSnapshot, error)
	// Returns a list of currently active intercepts.
	ListIntercepts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.InterceptInfoSnapshot, error)
	// Quits (terminates) the service.
	Quit(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type connectorClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorClient(cc grpc.ClientConnInterface) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*version.VersionInfo, error) {
	out := new(version.VersionInfo)
	err := c.cc.Invoke(ctx, "/telepresence.Connector/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ConnectorStatus, error) {
	out := new(ConnectorStatus)
	err := c.cc.Invoke(ctx, "/telepresence.Connector/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectInfo, error) {
	out := new(ConnectInfo)
	err := c.cc.Invoke(ctx, "/telepresence.Connector/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) CreateIntercept(ctx context.Context, in *rpc.CreateInterceptRequest, opts ...grpc.CallOption) (*InterceptResult, error) {
	out := new(InterceptResult)
	err := c.cc.Invoke(ctx, "/telepresence.Connector/CreateIntercept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) RemoveIntercept(ctx context.Context, in *rpc.RemoveInterceptRequest2, opts ...grpc.CallOption) (*InterceptResult, error) {
	out := new(InterceptResult)
	err := c.cc.Invoke(ctx, "/telepresence.Connector/RemoveIntercept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) AvailableIntercepts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.AgentInfoSnapshot, error) {
	out := new(rpc.AgentInfoSnapshot)
	err := c.cc.Invoke(ctx, "/telepresence.Connector/AvailableIntercepts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) ListIntercepts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*rpc.InterceptInfoSnapshot, error) {
	out := new(rpc.InterceptInfoSnapshot)
	err := c.cc.Invoke(ctx, "/telepresence.Connector/ListIntercepts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) Quit(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telepresence.Connector/Quit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorServer is the server API for Connector service.
// All implementations must embed UnimplementedConnectorServer
// for forward compatibility
type ConnectorServer interface {
	// Returns version information from the Connector
	Version(context.Context, *empty.Empty) (*version.VersionInfo, error)
	// Returns the current connectivity status
	Status(context.Context, *empty.Empty) (*ConnectorStatus, error)
	// Connects the daemon to a cluster
	Connect(context.Context, *ConnectRequest) (*ConnectInfo, error)
	// Adds a deployment intercept
	CreateIntercept(context.Context, *rpc.CreateInterceptRequest) (*InterceptResult, error)
	// Deactivates and removes an existent deployment intercept.
	RemoveIntercept(context.Context, *rpc.RemoveInterceptRequest2) (*InterceptResult, error)
	// Returns a list of deployments available for intercept.
	AvailableIntercepts(context.Context, *empty.Empty) (*rpc.AgentInfoSnapshot, error)
	// Returns a list of currently active intercepts.
	ListIntercepts(context.Context, *empty.Empty) (*rpc.InterceptInfoSnapshot, error)
	// Quits (terminates) the service.
	Quit(context.Context, *empty.Empty) (*empty.Empty, error)
	mustEmbedUnimplementedConnectorServer()
}

// UnimplementedConnectorServer must be embedded to have forward compatible implementations.
type UnimplementedConnectorServer struct {
}

func (UnimplementedConnectorServer) Version(context.Context, *empty.Empty) (*version.VersionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedConnectorServer) Status(context.Context, *empty.Empty) (*ConnectorStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedConnectorServer) Connect(context.Context, *ConnectRequest) (*ConnectInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedConnectorServer) CreateIntercept(context.Context, *rpc.CreateInterceptRequest) (*InterceptResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIntercept not implemented")
}
func (UnimplementedConnectorServer) RemoveIntercept(context.Context, *rpc.RemoveInterceptRequest2) (*InterceptResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveIntercept not implemented")
}
func (UnimplementedConnectorServer) AvailableIntercepts(context.Context, *empty.Empty) (*rpc.AgentInfoSnapshot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AvailableIntercepts not implemented")
}
func (UnimplementedConnectorServer) ListIntercepts(context.Context, *empty.Empty) (*rpc.InterceptInfoSnapshot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIntercepts not implemented")
}
func (UnimplementedConnectorServer) Quit(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Quit not implemented")
}
func (UnimplementedConnectorServer) mustEmbedUnimplementedConnectorServer() {}

// UnsafeConnectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorServer will
// result in compilation errors.
type UnsafeConnectorServer interface {
	mustEmbedUnimplementedConnectorServer()
}

func RegisterConnectorServer(s grpc.ServiceRegistrar, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Connector/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Connector/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Connector/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_CreateIntercept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.CreateInterceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).CreateIntercept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Connector/CreateIntercept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).CreateIntercept(ctx, req.(*rpc.CreateInterceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_RemoveIntercept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rpc.RemoveInterceptRequest2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).RemoveIntercept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Connector/RemoveIntercept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).RemoveIntercept(ctx, req.(*rpc.RemoveInterceptRequest2))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_AvailableIntercepts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).AvailableIntercepts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Connector/AvailableIntercepts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).AvailableIntercepts(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_ListIntercepts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).ListIntercepts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Connector/ListIntercepts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).ListIntercepts(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_Quit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Quit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telepresence.Connector/Quit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Quit(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telepresence.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Connector_Version_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Connector_Status_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Connector_Connect_Handler,
		},
		{
			MethodName: "CreateIntercept",
			Handler:    _Connector_CreateIntercept_Handler,
		},
		{
			MethodName: "RemoveIntercept",
			Handler:    _Connector_RemoveIntercept_Handler,
		},
		{
			MethodName: "AvailableIntercepts",
			Handler:    _Connector_AvailableIntercepts_Handler,
		},
		{
			MethodName: "ListIntercepts",
			Handler:    _Connector_ListIntercepts_Handler,
		},
		{
			MethodName: "Quit",
			Handler:    _Connector_Quit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/connector/connector.proto",
}