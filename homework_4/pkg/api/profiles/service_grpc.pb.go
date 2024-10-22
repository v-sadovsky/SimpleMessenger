// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: api/profiles/service.proto

package profiles

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
	ProfilesService_SaveProfile_FullMethodName  = "/github.com.v_sadovsky.simple_messenger.homework_4.ProfilesService/SaveProfile"
	ProfilesService_GetProfile_FullMethodName   = "/github.com.v_sadovsky.simple_messenger.homework_4.ProfilesService/GetProfile"
	ProfilesService_ListProfiles_FullMethodName = "/github.com.v_sadovsky.simple_messenger.homework_4.ProfilesService/ListProfiles"
)

// ProfilesServiceClient is the client API for ProfilesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ProfilesService - user profiles service
type ProfilesServiceClient interface {
	// SaveProfileRequest - save a user profile
	SaveProfile(ctx context.Context, in *SaveProfileRequest, opts ...grpc.CallOption) (*SaveProfileResponse, error)
	// GetProfile - get user profile by id
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	// ListProfiles - list all users
	ListProfiles(ctx context.Context, in *ListProfilesRequest, opts ...grpc.CallOption) (*ListProfilesResponse, error)
}

type profilesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfilesServiceClient(cc grpc.ClientConnInterface) ProfilesServiceClient {
	return &profilesServiceClient{cc}
}

func (c *profilesServiceClient) SaveProfile(ctx context.Context, in *SaveProfileRequest, opts ...grpc.CallOption) (*SaveProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveProfileResponse)
	err := c.cc.Invoke(ctx, ProfilesService_SaveProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, ProfilesService_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesServiceClient) ListProfiles(ctx context.Context, in *ListProfilesRequest, opts ...grpc.CallOption) (*ListProfilesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProfilesResponse)
	err := c.cc.Invoke(ctx, ProfilesService_ListProfiles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilesServiceServer is the server API for ProfilesService service.
// All implementations must embed UnimplementedProfilesServiceServer
// for forward compatibility.
//
// ProfilesService - user profiles service
type ProfilesServiceServer interface {
	// SaveProfileRequest - save a user profile
	SaveProfile(context.Context, *SaveProfileRequest) (*SaveProfileResponse, error)
	// GetProfile - get user profile by id
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	// ListProfiles - list all users
	ListProfiles(context.Context, *ListProfilesRequest) (*ListProfilesResponse, error)
	mustEmbedUnimplementedProfilesServiceServer()
}

// UnimplementedProfilesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfilesServiceServer struct{}

func (UnimplementedProfilesServiceServer) SaveProfile(context.Context, *SaveProfileRequest) (*SaveProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProfile not implemented")
}
func (UnimplementedProfilesServiceServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedProfilesServiceServer) ListProfiles(context.Context, *ListProfilesRequest) (*ListProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProfiles not implemented")
}
func (UnimplementedProfilesServiceServer) mustEmbedUnimplementedProfilesServiceServer() {}
func (UnimplementedProfilesServiceServer) testEmbeddedByValue()                         {}

// UnsafeProfilesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfilesServiceServer will
// result in compilation errors.
type UnsafeProfilesServiceServer interface {
	mustEmbedUnimplementedProfilesServiceServer()
}

func RegisterProfilesServiceServer(s grpc.ServiceRegistrar, srv ProfilesServiceServer) {
	// If the following call pancis, it indicates UnimplementedProfilesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProfilesService_ServiceDesc, srv)
}

func _ProfilesService_SaveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).SaveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfilesService_SaveProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).SaveProfile(ctx, req.(*SaveProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilesService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfilesService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilesService_ListProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServiceServer).ListProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfilesService_ListProfiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServiceServer).ListProfiles(ctx, req.(*ListProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfilesService_ServiceDesc is the grpc.ServiceDesc for ProfilesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfilesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.v_sadovsky.simple_messenger.homework_4.ProfilesService",
	HandlerType: (*ProfilesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveProfile",
			Handler:    _ProfilesService_SaveProfile_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _ProfilesService_GetProfile_Handler,
		},
		{
			MethodName: "ListProfiles",
			Handler:    _ProfilesService_ListProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/profiles/service.proto",
}
