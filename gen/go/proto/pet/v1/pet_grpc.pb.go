// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/pet/v1/pet.proto

package petv1

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

// PetStoreClient is the client API for PetStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetStoreClient interface {
	GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetResponse, error)
}

type petStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewPetStoreClient(cc grpc.ClientConnInterface) PetStoreClient {
	return &petStoreClient{cc}
}

func (c *petStoreClient) GetPet(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*GetPetResponse, error) {
	out := new(GetPetResponse)
	err := c.cc.Invoke(ctx, "/pet.v1.PetStore/GetPet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetStoreServer is the server API for PetStore service.
// All implementations should embed UnimplementedPetStoreServer
// for forward compatibility
type PetStoreServer interface {
	GetPet(context.Context, *GetPetRequest) (*GetPetResponse, error)
}

// UnimplementedPetStoreServer should be embedded to have forward compatible implementations.
type UnimplementedPetStoreServer struct {
}

func (UnimplementedPetStoreServer) GetPet(context.Context, *GetPetRequest) (*GetPetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPet not implemented")
}

// UnsafePetStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetStoreServer will
// result in compilation errors.
type UnsafePetStoreServer interface {
	mustEmbedUnimplementedPetStoreServer()
}

func RegisterPetStoreServer(s grpc.ServiceRegistrar, srv PetStoreServer) {
	s.RegisterService(&PetStore_ServiceDesc, srv)
}

func _PetStore_GetPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetStoreServer).GetPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pet.v1.PetStore/GetPet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetStoreServer).GetPet(ctx, req.(*GetPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetStore_ServiceDesc is the grpc.ServiceDesc for PetStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pet.v1.PetStore",
	HandlerType: (*PetStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPet",
			Handler:    _PetStore_GetPet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pet/v1/pet.proto",
}
