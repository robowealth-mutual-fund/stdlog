// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: masterData/country/v1/service.proto

package v1

import (
	context "context"
	message "git.robodev.co/roa-v2/proto-open-account/pkg/masterData/country/v1/message"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_masterData_country_v1_service_proto protoreflect.FileDescriptor

var file_masterData_country_v1_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x2b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xba, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x06,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2a, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x3a,
	0x01, 0x2a, 0x28, 0x01, 0x12, 0x7c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a,
	0x01, 0x2a, 0x1a, 0x21, 0x92, 0x41, 0x1e, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x13, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x20, 0x44, 0x61, 0x74, 0x61, 0x20, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0xdd, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x6f, 0x2f, 0x72, 0x6f, 0x61, 0x2d, 0x76, 0x32, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x92, 0x41, 0x95, 0x01,
	0x12, 0x5e, 0x0a, 0x13, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x20, 0x44, 0x61, 0x74, 0x61, 0x20,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x40, 0x0a, 0x25, 0x4f, 0x70, 0x65, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x20, 0x41, 0x63, 0x63, 0x6f, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x20, 0x44, 0x61, 0x74, 0x61, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x17, 0x79, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x40, 0x72, 0x6f, 0x62, 0x6f, 0x77, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x63, 0x6f, 0x2e, 0x74, 0x68, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30,
	0x22, 0x0b, 0x7b, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x7d, 0x2a, 0x02, 0x01,
	0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_masterData_country_v1_service_proto_goTypes = []interface{}{
	(*message.ImportRequest)(nil),  // 0: master.data.country.message.ImportRequest
	(*message.ListRequest)(nil),    // 1: master.data.country.message.ListRequest
	(*message.ImportResponse)(nil), // 2: master.data.country.message.ImportResponse
	(*message.ListResponse)(nil),   // 3: master.data.country.message.ListResponse
}
var file_masterData_country_v1_service_proto_depIdxs = []int32{
	0, // 0: master.data.country.service.CountryService.Import:input_type -> master.data.country.message.ImportRequest
	1, // 1: master.data.country.service.CountryService.List:input_type -> master.data.country.message.ListRequest
	2, // 2: master.data.country.service.CountryService.Import:output_type -> master.data.country.message.ImportResponse
	3, // 3: master.data.country.service.CountryService.List:output_type -> master.data.country.message.ListResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_masterData_country_v1_service_proto_init() }
func file_masterData_country_v1_service_proto_init() {
	if File_masterData_country_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_masterData_country_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_masterData_country_v1_service_proto_goTypes,
		DependencyIndexes: file_masterData_country_v1_service_proto_depIdxs,
	}.Build()
	File_masterData_country_v1_service_proto = out.File
	file_masterData_country_v1_service_proto_rawDesc = nil
	file_masterData_country_v1_service_proto_goTypes = nil
	file_masterData_country_v1_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CountryServiceClient is the client API for CountryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountryServiceClient interface {
	Import(ctx context.Context, opts ...grpc.CallOption) (CountryService_ImportClient, error)
	List(ctx context.Context, in *message.ListRequest, opts ...grpc.CallOption) (*message.ListResponse, error)
}

type countryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCountryServiceClient(cc grpc.ClientConnInterface) CountryServiceClient {
	return &countryServiceClient{cc}
}

func (c *countryServiceClient) Import(ctx context.Context, opts ...grpc.CallOption) (CountryService_ImportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CountryService_serviceDesc.Streams[0], "/master.data.country.service.CountryService/Import", opts...)
	if err != nil {
		return nil, err
	}
	x := &countryServiceImportClient{stream}
	return x, nil
}

type CountryService_ImportClient interface {
	Send(*message.ImportRequest) error
	CloseAndRecv() (*message.ImportResponse, error)
	grpc.ClientStream
}

type countryServiceImportClient struct {
	grpc.ClientStream
}

func (x *countryServiceImportClient) Send(m *message.ImportRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *countryServiceImportClient) CloseAndRecv() (*message.ImportResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(message.ImportResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *countryServiceClient) List(ctx context.Context, in *message.ListRequest, opts ...grpc.CallOption) (*message.ListResponse, error) {
	out := new(message.ListResponse)
	err := c.cc.Invoke(ctx, "/master.data.country.service.CountryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountryServiceServer is the server API for CountryService service.
type CountryServiceServer interface {
	Import(CountryService_ImportServer) error
	List(context.Context, *message.ListRequest) (*message.ListResponse, error)
}

// UnimplementedCountryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCountryServiceServer struct {
}

func (*UnimplementedCountryServiceServer) Import(CountryService_ImportServer) error {
	return status.Errorf(codes.Unimplemented, "method Import not implemented")
}
func (*UnimplementedCountryServiceServer) List(context.Context, *message.ListRequest) (*message.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterCountryServiceServer(s *grpc.Server, srv CountryServiceServer) {
	s.RegisterService(&_CountryService_serviceDesc, srv)
}

func _CountryService_Import_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CountryServiceServer).Import(&countryServiceImportServer{stream})
}

type CountryService_ImportServer interface {
	SendAndClose(*message.ImportResponse) error
	Recv() (*message.ImportRequest, error)
	grpc.ServerStream
}

type countryServiceImportServer struct {
	grpc.ServerStream
}

func (x *countryServiceImportServer) SendAndClose(m *message.ImportResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *countryServiceImportServer) Recv() (*message.ImportRequest, error) {
	m := new(message.ImportRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CountryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.data.country.service.CountryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountryServiceServer).List(ctx, req.(*message.ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "master.data.country.service.CountryService",
	HandlerType: (*CountryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CountryService_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Import",
			Handler:       _CountryService_Import_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "masterData/country/v1/service.proto",
}