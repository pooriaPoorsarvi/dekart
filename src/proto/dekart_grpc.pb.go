// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: proto/dekart.proto

package proto

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

// DekartClient is the client API for Dekart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DekartClient interface {
	//reports
	CreateReport(ctx context.Context, in *CreateReportRequest, opts ...grpc.CallOption) (*CreateReportResponse, error)
	ForkReport(ctx context.Context, in *ForkReportRequest, opts ...grpc.CallOption) (*ForkReportResponse, error)
	UpdateReport(ctx context.Context, in *UpdateReportRequest, opts ...grpc.CallOption) (*UpdateReportResponse, error)
	ArchiveReport(ctx context.Context, in *ArchiveReportRequest, opts ...grpc.CallOption) (*ArchiveReportResponse, error)
	SetDiscoverable(ctx context.Context, in *SetDiscoverableRequest, opts ...grpc.CallOption) (*SetDiscoverableResponse, error)
	// datasets
	CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*CreateDatasetResponse, error)
	RemoveDataset(ctx context.Context, in *RemoveDatasetRequest, opts ...grpc.CallOption) (*RemoveDatasetResponse, error)
	// files
	CreateFile(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*CreateFileResponse, error)
	// queries
	CreateQuery(ctx context.Context, in *CreateQueryRequest, opts ...grpc.CallOption) (*CreateQueryResponse, error)
	RunQuery(ctx context.Context, in *RunQueryRequest, opts ...grpc.CallOption) (*RunQueryResponse, error)
	CancelQuery(ctx context.Context, in *CancelQueryRequest, opts ...grpc.CallOption) (*CancelQueryResponse, error)
	GetEnv(ctx context.Context, in *GetEnvRequest, opts ...grpc.CallOption) (*GetEnvResponse, error)
	// streams
	GetReportStream(ctx context.Context, in *ReportStreamRequest, opts ...grpc.CallOption) (Dekart_GetReportStreamClient, error)
	GetReportListStream(ctx context.Context, in *ReportListRequest, opts ...grpc.CallOption) (Dekart_GetReportListStreamClient, error)
	//statistics
	GetUsage(ctx context.Context, in *GetUsageRequest, opts ...grpc.CallOption) (*GetUsageResponse, error)
}

type dekartClient struct {
	cc grpc.ClientConnInterface
}

func NewDekartClient(cc grpc.ClientConnInterface) DekartClient {
	return &dekartClient{cc}
}

func (c *dekartClient) CreateReport(ctx context.Context, in *CreateReportRequest, opts ...grpc.CallOption) (*CreateReportResponse, error) {
	out := new(CreateReportResponse)
	err := c.cc.Invoke(ctx, "/Dekart/CreateReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) ForkReport(ctx context.Context, in *ForkReportRequest, opts ...grpc.CallOption) (*ForkReportResponse, error) {
	out := new(ForkReportResponse)
	err := c.cc.Invoke(ctx, "/Dekart/ForkReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) UpdateReport(ctx context.Context, in *UpdateReportRequest, opts ...grpc.CallOption) (*UpdateReportResponse, error) {
	out := new(UpdateReportResponse)
	err := c.cc.Invoke(ctx, "/Dekart/UpdateReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) ArchiveReport(ctx context.Context, in *ArchiveReportRequest, opts ...grpc.CallOption) (*ArchiveReportResponse, error) {
	out := new(ArchiveReportResponse)
	err := c.cc.Invoke(ctx, "/Dekart/ArchiveReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) SetDiscoverable(ctx context.Context, in *SetDiscoverableRequest, opts ...grpc.CallOption) (*SetDiscoverableResponse, error) {
	out := new(SetDiscoverableResponse)
	err := c.cc.Invoke(ctx, "/Dekart/SetDiscoverable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*CreateDatasetResponse, error) {
	out := new(CreateDatasetResponse)
	err := c.cc.Invoke(ctx, "/Dekart/CreateDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) RemoveDataset(ctx context.Context, in *RemoveDatasetRequest, opts ...grpc.CallOption) (*RemoveDatasetResponse, error) {
	out := new(RemoveDatasetResponse)
	err := c.cc.Invoke(ctx, "/Dekart/RemoveDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) CreateFile(ctx context.Context, in *CreateFileRequest, opts ...grpc.CallOption) (*CreateFileResponse, error) {
	out := new(CreateFileResponse)
	err := c.cc.Invoke(ctx, "/Dekart/CreateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) CreateQuery(ctx context.Context, in *CreateQueryRequest, opts ...grpc.CallOption) (*CreateQueryResponse, error) {
	out := new(CreateQueryResponse)
	err := c.cc.Invoke(ctx, "/Dekart/CreateQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) RunQuery(ctx context.Context, in *RunQueryRequest, opts ...grpc.CallOption) (*RunQueryResponse, error) {
	out := new(RunQueryResponse)
	err := c.cc.Invoke(ctx, "/Dekart/RunQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) CancelQuery(ctx context.Context, in *CancelQueryRequest, opts ...grpc.CallOption) (*CancelQueryResponse, error) {
	out := new(CancelQueryResponse)
	err := c.cc.Invoke(ctx, "/Dekart/CancelQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) GetEnv(ctx context.Context, in *GetEnvRequest, opts ...grpc.CallOption) (*GetEnvResponse, error) {
	out := new(GetEnvResponse)
	err := c.cc.Invoke(ctx, "/Dekart/GetEnv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dekartClient) GetReportStream(ctx context.Context, in *ReportStreamRequest, opts ...grpc.CallOption) (Dekart_GetReportStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dekart_ServiceDesc.Streams[0], "/Dekart/GetReportStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dekartGetReportStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dekart_GetReportStreamClient interface {
	Recv() (*ReportStreamResponse, error)
	grpc.ClientStream
}

type dekartGetReportStreamClient struct {
	grpc.ClientStream
}

func (x *dekartGetReportStreamClient) Recv() (*ReportStreamResponse, error) {
	m := new(ReportStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dekartClient) GetReportListStream(ctx context.Context, in *ReportListRequest, opts ...grpc.CallOption) (Dekart_GetReportListStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dekart_ServiceDesc.Streams[1], "/Dekart/GetReportListStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dekartGetReportListStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dekart_GetReportListStreamClient interface {
	Recv() (*ReportListResponse, error)
	grpc.ClientStream
}

type dekartGetReportListStreamClient struct {
	grpc.ClientStream
}

func (x *dekartGetReportListStreamClient) Recv() (*ReportListResponse, error) {
	m := new(ReportListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dekartClient) GetUsage(ctx context.Context, in *GetUsageRequest, opts ...grpc.CallOption) (*GetUsageResponse, error) {
	out := new(GetUsageResponse)
	err := c.cc.Invoke(ctx, "/Dekart/GetUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DekartServer is the server API for Dekart service.
// All implementations must embed UnimplementedDekartServer
// for forward compatibility
type DekartServer interface {
	//reports
	CreateReport(context.Context, *CreateReportRequest) (*CreateReportResponse, error)
	ForkReport(context.Context, *ForkReportRequest) (*ForkReportResponse, error)
	UpdateReport(context.Context, *UpdateReportRequest) (*UpdateReportResponse, error)
	ArchiveReport(context.Context, *ArchiveReportRequest) (*ArchiveReportResponse, error)
	SetDiscoverable(context.Context, *SetDiscoverableRequest) (*SetDiscoverableResponse, error)
	// datasets
	CreateDataset(context.Context, *CreateDatasetRequest) (*CreateDatasetResponse, error)
	RemoveDataset(context.Context, *RemoveDatasetRequest) (*RemoveDatasetResponse, error)
	// files
	CreateFile(context.Context, *CreateFileRequest) (*CreateFileResponse, error)
	// queries
	CreateQuery(context.Context, *CreateQueryRequest) (*CreateQueryResponse, error)
	RunQuery(context.Context, *RunQueryRequest) (*RunQueryResponse, error)
	CancelQuery(context.Context, *CancelQueryRequest) (*CancelQueryResponse, error)
	GetEnv(context.Context, *GetEnvRequest) (*GetEnvResponse, error)
	// streams
	GetReportStream(*ReportStreamRequest, Dekart_GetReportStreamServer) error
	GetReportListStream(*ReportListRequest, Dekart_GetReportListStreamServer) error
	//statistics
	GetUsage(context.Context, *GetUsageRequest) (*GetUsageResponse, error)
	mustEmbedUnimplementedDekartServer()
}

// UnimplementedDekartServer must be embedded to have forward compatible implementations.
type UnimplementedDekartServer struct {
}

func (UnimplementedDekartServer) CreateReport(context.Context, *CreateReportRequest) (*CreateReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReport not implemented")
}
func (UnimplementedDekartServer) ForkReport(context.Context, *ForkReportRequest) (*ForkReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForkReport not implemented")
}
func (UnimplementedDekartServer) UpdateReport(context.Context, *UpdateReportRequest) (*UpdateReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReport not implemented")
}
func (UnimplementedDekartServer) ArchiveReport(context.Context, *ArchiveReportRequest) (*ArchiveReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveReport not implemented")
}
func (UnimplementedDekartServer) SetDiscoverable(context.Context, *SetDiscoverableRequest) (*SetDiscoverableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDiscoverable not implemented")
}
func (UnimplementedDekartServer) CreateDataset(context.Context, *CreateDatasetRequest) (*CreateDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDataset not implemented")
}
func (UnimplementedDekartServer) RemoveDataset(context.Context, *RemoveDatasetRequest) (*RemoveDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDataset not implemented")
}
func (UnimplementedDekartServer) CreateFile(context.Context, *CreateFileRequest) (*CreateFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedDekartServer) CreateQuery(context.Context, *CreateQueryRequest) (*CreateQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuery not implemented")
}
func (UnimplementedDekartServer) RunQuery(context.Context, *RunQueryRequest) (*RunQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunQuery not implemented")
}
func (UnimplementedDekartServer) CancelQuery(context.Context, *CancelQueryRequest) (*CancelQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelQuery not implemented")
}
func (UnimplementedDekartServer) GetEnv(context.Context, *GetEnvRequest) (*GetEnvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnv not implemented")
}
func (UnimplementedDekartServer) GetReportStream(*ReportStreamRequest, Dekart_GetReportStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetReportStream not implemented")
}
func (UnimplementedDekartServer) GetReportListStream(*ReportListRequest, Dekart_GetReportListStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetReportListStream not implemented")
}
func (UnimplementedDekartServer) GetUsage(context.Context, *GetUsageRequest) (*GetUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsage not implemented")
}
func (UnimplementedDekartServer) mustEmbedUnimplementedDekartServer() {}

// UnsafeDekartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DekartServer will
// result in compilation errors.
type UnsafeDekartServer interface {
	mustEmbedUnimplementedDekartServer()
}

func RegisterDekartServer(s grpc.ServiceRegistrar, srv DekartServer) {
	s.RegisterService(&Dekart_ServiceDesc, srv)
}

func _Dekart_CreateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).CreateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/CreateReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).CreateReport(ctx, req.(*CreateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_ForkReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForkReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).ForkReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/ForkReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).ForkReport(ctx, req.(*ForkReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_UpdateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).UpdateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/UpdateReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).UpdateReport(ctx, req.(*UpdateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_ArchiveReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).ArchiveReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/ArchiveReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).ArchiveReport(ctx, req.(*ArchiveReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_SetDiscoverable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDiscoverableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).SetDiscoverable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/SetDiscoverable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).SetDiscoverable(ctx, req.(*SetDiscoverableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_CreateDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).CreateDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/CreateDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).CreateDataset(ctx, req.(*CreateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_RemoveDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).RemoveDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/RemoveDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).RemoveDataset(ctx, req.(*RemoveDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).CreateFile(ctx, req.(*CreateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_CreateQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).CreateQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/CreateQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).CreateQuery(ctx, req.(*CreateQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_RunQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).RunQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/RunQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).RunQuery(ctx, req.(*RunQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_CancelQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).CancelQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/CancelQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).CancelQuery(ctx, req.(*CancelQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_GetEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).GetEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/GetEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).GetEnv(ctx, req.(*GetEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dekart_GetReportStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReportStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DekartServer).GetReportStream(m, &dekartGetReportStreamServer{stream})
}

type Dekart_GetReportStreamServer interface {
	Send(*ReportStreamResponse) error
	grpc.ServerStream
}

type dekartGetReportStreamServer struct {
	grpc.ServerStream
}

func (x *dekartGetReportStreamServer) Send(m *ReportStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Dekart_GetReportListStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReportListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DekartServer).GetReportListStream(m, &dekartGetReportListStreamServer{stream})
}

type Dekart_GetReportListStreamServer interface {
	Send(*ReportListResponse) error
	grpc.ServerStream
}

type dekartGetReportListStreamServer struct {
	grpc.ServerStream
}

func (x *dekartGetReportListStreamServer) Send(m *ReportListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Dekart_GetUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DekartServer).GetUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dekart/GetUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DekartServer).GetUsage(ctx, req.(*GetUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dekart_ServiceDesc is the grpc.ServiceDesc for Dekart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dekart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Dekart",
	HandlerType: (*DekartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReport",
			Handler:    _Dekart_CreateReport_Handler,
		},
		{
			MethodName: "ForkReport",
			Handler:    _Dekart_ForkReport_Handler,
		},
		{
			MethodName: "UpdateReport",
			Handler:    _Dekart_UpdateReport_Handler,
		},
		{
			MethodName: "ArchiveReport",
			Handler:    _Dekart_ArchiveReport_Handler,
		},
		{
			MethodName: "SetDiscoverable",
			Handler:    _Dekart_SetDiscoverable_Handler,
		},
		{
			MethodName: "CreateDataset",
			Handler:    _Dekart_CreateDataset_Handler,
		},
		{
			MethodName: "RemoveDataset",
			Handler:    _Dekart_RemoveDataset_Handler,
		},
		{
			MethodName: "CreateFile",
			Handler:    _Dekart_CreateFile_Handler,
		},
		{
			MethodName: "CreateQuery",
			Handler:    _Dekart_CreateQuery_Handler,
		},
		{
			MethodName: "RunQuery",
			Handler:    _Dekart_RunQuery_Handler,
		},
		{
			MethodName: "CancelQuery",
			Handler:    _Dekart_CancelQuery_Handler,
		},
		{
			MethodName: "GetEnv",
			Handler:    _Dekart_GetEnv_Handler,
		},
		{
			MethodName: "GetUsage",
			Handler:    _Dekart_GetUsage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetReportStream",
			Handler:       _Dekart_GetReportStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetReportListStream",
			Handler:       _Dekart_GetReportListStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/dekart.proto",
}
