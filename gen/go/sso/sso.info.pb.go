// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: sso/sso.info.proto

package ssov1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    int64                  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Email     string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	VisitedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=visited_at,json=visitedAt,proto3" json:"visited_at,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_info_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserResponse) GetVisitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.VisitedAt
	}
	return nil
}

type UserRolesByAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *UserRequest `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AppID int64        `protobuf:"varint,2,opt,name=appID,proto3" json:"appID,omitempty"`
}

func (x *UserRolesByAppRequest) Reset() {
	*x = UserRolesByAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRolesByAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRolesByAppRequest) ProtoMessage() {}

func (x *UserRolesByAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRolesByAppRequest.ProtoReflect.Descriptor instead.
func (*UserRolesByAppRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_info_proto_rawDescGZIP(), []int{2}
}

func (x *UserRolesByAppRequest) GetUser() *UserRequest {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserRolesByAppRequest) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

type UserRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role []*Role `protobuf:"bytes,1,rep,name=role,proto3" json:"role,omitempty"`
}

func (x *UserRolesResponse) Reset() {
	*x = UserRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRolesResponse) ProtoMessage() {}

func (x *UserRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRolesResponse.ProtoReflect.Descriptor instead.
func (*UserRolesResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_info_proto_rawDescGZIP(), []int{3}
}

func (x *UserRolesResponse) GetRole() []*Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type UserRolesByAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role []*Role `protobuf:"bytes,1,rep,name=role,proto3" json:"role,omitempty"`
}

func (x *UserRolesByAppResponse) Reset() {
	*x = UserRolesByAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRolesByAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRolesByAppResponse) ProtoMessage() {}

func (x *UserRolesByAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRolesByAppResponse.ProtoReflect.Descriptor instead.
func (*UserRolesByAppResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_info_proto_rawDescGZIP(), []int{4}
}

func (x *UserRolesByAppResponse) GetRole() []*Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     int64                  `protobuf:"varint,2,opt,name=appID,proto3" json:"appID,omitempty"`
	Name      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UsedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=used_at,json=usedAt,proto3" json:"used_at,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_sso_sso_info_proto_rawDescGZIP(), []int{5}
}

func (x *App) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *App) GetUsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UsedAt
	}
	return nil
}

type AppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App *App `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *AppResponse) Reset() {
	*x = AppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_info_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppResponse) ProtoMessage() {}

func (x *AppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_info_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppResponse.ProtoReflect.Descriptor instead.
func (*AppResponse) Descriptor() ([]byte, []int) {
	return file_sso_sso_info_proto_rawDescGZIP(), []int{6}
}

func (x *AppResponse) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

type AppByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID int64 `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
}

func (x *AppByIdRequest) Reset() {
	*x = AppByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_sso_info_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppByIdRequest) ProtoMessage() {}

func (x *AppByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_sso_info_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppByIdRequest.ProtoReflect.Descriptor instead.
func (*AppByIdRequest) Descriptor() ([]byte, []int) {
	return file_sso_sso_info_proto_rawDescGZIP(), []int{7}
}

func (x *AppByIdRequest) GetAppID() int64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

var File_sso_sso_info_proto protoreflect.FileDescriptor

var file_sso_sso_info_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x73, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x73, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0xb2, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x76,
	0x69, 0x73, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x76, 0x69, 0x73,
	0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x54, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x42, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x22, 0x33, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0x38, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x79,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x03,
	0x41, 0x70, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x75, 0x73, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2a, 0x0a,
	0x0b, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x03,
	0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0x26, 0x0a, 0x0e, 0x41, 0x70, 0x70,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x44, 0x32, 0xcc, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12,
	0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x43, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x11, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x4f, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x41, 0x70, 0x70, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x3a, 0x01, 0x2a,
	0x32, 0x8e, 0x01, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x0a, 0x03,
	0x41, 0x70, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x06, 0x12, 0x04, 0x2f, 0x61, 0x70, 0x70, 0x12, 0x43, 0x0a, 0x07,
	0x41, 0x70, 0x70, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x41,
	0x70, 0x70, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22, 0x04, 0x2f, 0x61, 0x70, 0x70, 0x3a, 0x01,
	0x2a, 0x42, 0x17, 0x5a, 0x15, 0x6b, 0x75, 0x72, 0x62, 0x61, 0x6e, 0x6f, 0x76, 0x2e, 0x73, 0x73,
	0x6f, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x73, 0x6f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sso_sso_info_proto_rawDescOnce sync.Once
	file_sso_sso_info_proto_rawDescData = file_sso_sso_info_proto_rawDesc
)

func file_sso_sso_info_proto_rawDescGZIP() []byte {
	file_sso_sso_info_proto_rawDescOnce.Do(func() {
		file_sso_sso_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_sso_info_proto_rawDescData)
	})
	return file_sso_sso_info_proto_rawDescData
}

var file_sso_sso_info_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sso_sso_info_proto_goTypes = []interface{}{
	(*UserRequest)(nil),            // 0: info.UserRequest
	(*UserResponse)(nil),           // 1: info.UserResponse
	(*UserRolesByAppRequest)(nil),  // 2: info.UserRolesByAppRequest
	(*UserRolesResponse)(nil),      // 3: info.UserRolesResponse
	(*UserRolesByAppResponse)(nil), // 4: info.UserRolesByAppResponse
	(*App)(nil),                    // 5: info.App
	(*AppResponse)(nil),            // 6: info.AppResponse
	(*AppByIdRequest)(nil),         // 7: info.AppByIdRequest
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(*Role)(nil),                   // 9: role.Role
	(*emptypb.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_sso_sso_info_proto_depIdxs = []int32{
	8,  // 0: info.UserResponse.created_at:type_name -> google.protobuf.Timestamp
	8,  // 1: info.UserResponse.visited_at:type_name -> google.protobuf.Timestamp
	0,  // 2: info.UserRolesByAppRequest.user:type_name -> info.UserRequest
	9,  // 3: info.UserRolesResponse.role:type_name -> role.Role
	9,  // 4: info.UserRolesByAppResponse.role:type_name -> role.Role
	8,  // 5: info.App.created_at:type_name -> google.protobuf.Timestamp
	8,  // 6: info.App.used_at:type_name -> google.protobuf.Timestamp
	5,  // 7: info.AppResponse.app:type_name -> info.App
	10, // 8: info.UserInfo.User:input_type -> google.protobuf.Empty
	0,  // 9: info.UserInfo.UserByID:input_type -> info.UserRequest
	0,  // 10: info.UserInfo.UserRoles:input_type -> info.UserRequest
	2,  // 11: info.UserInfo.UserRolesByApp:input_type -> info.UserRolesByAppRequest
	10, // 12: info.AppInfo.App:input_type -> google.protobuf.Empty
	7,  // 13: info.AppInfo.AppByID:input_type -> info.AppByIdRequest
	1,  // 14: info.UserInfo.User:output_type -> info.UserResponse
	1,  // 15: info.UserInfo.UserByID:output_type -> info.UserResponse
	3,  // 16: info.UserInfo.UserRoles:output_type -> info.UserRolesResponse
	4,  // 17: info.UserInfo.UserRolesByApp:output_type -> info.UserRolesByAppResponse
	6,  // 18: info.AppInfo.App:output_type -> info.AppResponse
	6,  // 19: info.AppInfo.AppByID:output_type -> info.AppResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_sso_sso_info_proto_init() }
func file_sso_sso_info_proto_init() {
	if File_sso_sso_info_proto != nil {
		return
	}
	file_sso_sso_role_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sso_sso_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sso_sso_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sso_sso_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRolesByAppRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sso_sso_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRolesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sso_sso_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRolesByAppResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sso_sso_info_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sso_sso_info_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sso_sso_info_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sso_sso_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_sso_sso_info_proto_goTypes,
		DependencyIndexes: file_sso_sso_info_proto_depIdxs,
		MessageInfos:      file_sso_sso_info_proto_msgTypes,
	}.Build()
	File_sso_sso_info_proto = out.File
	file_sso_sso_info_proto_rawDesc = nil
	file_sso_sso_info_proto_goTypes = nil
	file_sso_sso_info_proto_depIdxs = nil
}