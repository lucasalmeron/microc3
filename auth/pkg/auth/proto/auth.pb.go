// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: pkg/auth/proto/auth.proto

package go_micro_service_auth

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RequestAuthID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestAuthID) Reset() {
	*x = RequestAuthID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_auth_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAuthID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAuthID) ProtoMessage() {}

func (x *RequestAuthID) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAuthID.ProtoReflect.Descriptor instead.
func (*RequestAuthID) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RequestAuthID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResponseAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone      string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address    string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	District   string   `protobuf:"bytes,5,opt,name=district,proto3" json:"district,omitempty"`
	Department string   `protobuf:"bytes,6,opt,name=department,proto3" json:"department,omitempty"`
	Actions    []string `protobuf:"bytes,7,rep,name=actions,proto3" json:"actions,omitempty"`
	CreatedAt  int64    `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ModifiedAt int64    `protobuf:"varint,9,opt,name=modifiedAt,proto3" json:"modifiedAt,omitempty"`
	DeletedAt  int64    `protobuf:"varint,10,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *ResponseAuth) Reset() {
	*x = ResponseAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_auth_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAuth) ProtoMessage() {}

func (x *ResponseAuth) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAuth.ProtoReflect.Descriptor instead.
func (*ResponseAuth) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseAuth) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseAuth) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseAuth) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ResponseAuth) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ResponseAuth) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *ResponseAuth) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *ResponseAuth) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *ResponseAuth) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ResponseAuth) GetModifiedAt() int64 {
	if x != nil {
		return x.ModifiedAt
	}
	return 0
}

func (x *ResponseAuth) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type RequestAuthLogIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RequestAuthLogIn) Reset() {
	*x = RequestAuthLogIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_auth_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAuthLogIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAuthLogIn) ProtoMessage() {}

func (x *RequestAuthLogIn) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAuthLogIn.ProtoReflect.Descriptor instead.
func (*RequestAuthLogIn) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RequestAuthLogIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RequestAuthLogIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ResponseLogIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ResponseLogIn) Reset() {
	*x = ResponseLogIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_auth_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLogIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLogIn) ProtoMessage() {}

func (x *ResponseLogIn) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLogIn.ProtoReflect.Descriptor instead.
func (*ResponseLogIn) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseLogIn) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RequestAuthPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Method string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *RequestAuthPath) Reset() {
	*x = RequestAuthPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_auth_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAuthPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAuthPath) ProtoMessage() {}

func (x *RequestAuthPath) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAuthPath.ProtoReflect.Descriptor instead.
func (*RequestAuthPath) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *RequestAuthPath) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RequestAuthPath) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RequestAuthPath) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type ResponseAuthPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
}

func (x *ResponseAuthPath) Reset() {
	*x = ResponseAuthPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_auth_proto_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAuthPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAuthPath) ProtoMessage() {}

func (x *ResponseAuthPath) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAuthPath.ProtoReflect.Descriptor instead.
func (*ResponseAuthPath) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseAuthPath) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

type RequestCreateAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone      string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Address    string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	District   string   `protobuf:"bytes,4,opt,name=district,proto3" json:"district,omitempty"`
	Department string   `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
	Actions    []string `protobuf:"bytes,6,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *RequestCreateAuth) Reset() {
	*x = RequestCreateAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_auth_proto_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreateAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreateAuth) ProtoMessage() {}

func (x *RequestCreateAuth) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreateAuth.ProtoReflect.Descriptor instead.
func (*RequestCreateAuth) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_auth_proto_rawDescGZIP(), []int{6}
}

func (x *RequestCreateAuth) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestCreateAuth) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RequestCreateAuth) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RequestCreateAuth) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *RequestCreateAuth) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *RequestCreateAuth) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type RequestUpdateAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone      string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address    string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	District   string   `protobuf:"bytes,5,opt,name=district,proto3" json:"district,omitempty"`
	Department string   `protobuf:"bytes,6,opt,name=department,proto3" json:"department,omitempty"`
	Actions    []string `protobuf:"bytes,7,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *RequestUpdateAuth) Reset() {
	*x = RequestUpdateAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_auth_proto_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUpdateAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUpdateAuth) ProtoMessage() {}

func (x *RequestUpdateAuth) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUpdateAuth.ProtoReflect.Descriptor instead.
func (*RequestUpdateAuth) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_auth_proto_rawDescGZIP(), []int{7}
}

func (x *RequestUpdateAuth) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RequestUpdateAuth) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestUpdateAuth) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RequestUpdateAuth) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RequestUpdateAuth) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *RequestUpdateAuth) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *RequestUpdateAuth) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_pkg_auth_proto_auth_proto protoreflect.FileDescriptor

var file_pkg_auth_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x1f, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x94, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x44, 0x0a, 0x10, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x53, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x32, 0x0a, 0x10,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x22, 0xad, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xbd, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0xa4, 0x04, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x58, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x12, 0x27, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x1a, 0x24, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x26, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x00, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x24, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x49, 0x44, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x23,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x49, 0x44,
	0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_auth_proto_auth_proto_rawDescOnce sync.Once
	file_pkg_auth_proto_auth_proto_rawDescData = file_pkg_auth_proto_auth_proto_rawDesc
)

func file_pkg_auth_proto_auth_proto_rawDescGZIP() []byte {
	file_pkg_auth_proto_auth_proto_rawDescOnce.Do(func() {
		file_pkg_auth_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_auth_proto_auth_proto_rawDescData)
	})
	return file_pkg_auth_proto_auth_proto_rawDescData
}

var file_pkg_auth_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_auth_proto_auth_proto_goTypes = []interface{}{
	(*RequestAuthID)(nil),     // 0: go.micro.service.auth.RequestAuthID
	(*ResponseAuth)(nil),      // 1: go.micro.service.auth.ResponseAuth
	(*RequestAuthLogIn)(nil),  // 2: go.micro.service.auth.RequestAuthLogIn
	(*ResponseLogIn)(nil),     // 3: go.micro.service.auth.ResponseLogIn
	(*RequestAuthPath)(nil),   // 4: go.micro.service.auth.RequestAuthPath
	(*ResponseAuthPath)(nil),  // 5: go.micro.service.auth.ResponseAuthPath
	(*RequestCreateAuth)(nil), // 6: go.micro.service.auth.RequestCreateAuth
	(*RequestUpdateAuth)(nil), // 7: go.micro.service.auth.RequestUpdateAuth
}
var file_pkg_auth_proto_auth_proto_depIdxs = []int32{
	2, // 0: go.micro.service.auth.Auth.LogIn:input_type -> go.micro.service.auth.RequestAuthLogIn
	4, // 1: go.micro.service.auth.Auth.AuthPath:input_type -> go.micro.service.auth.RequestAuthPath
	0, // 2: go.micro.service.auth.Auth.GetByID:input_type -> go.micro.service.auth.RequestAuthID
	6, // 3: go.micro.service.auth.Auth.Create:input_type -> go.micro.service.auth.RequestCreateAuth
	7, // 4: go.micro.service.auth.Auth.Update:input_type -> go.micro.service.auth.RequestUpdateAuth
	0, // 5: go.micro.service.auth.Auth.Delete:input_type -> go.micro.service.auth.RequestAuthID
	3, // 6: go.micro.service.auth.Auth.LogIn:output_type -> go.micro.service.auth.ResponseLogIn
	5, // 7: go.micro.service.auth.Auth.AuthPath:output_type -> go.micro.service.auth.ResponseAuthPath
	1, // 8: go.micro.service.auth.Auth.GetByID:output_type -> go.micro.service.auth.ResponseAuth
	1, // 9: go.micro.service.auth.Auth.Create:output_type -> go.micro.service.auth.ResponseAuth
	1, // 10: go.micro.service.auth.Auth.Update:output_type -> go.micro.service.auth.ResponseAuth
	1, // 11: go.micro.service.auth.Auth.Delete:output_type -> go.micro.service.auth.ResponseAuth
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_auth_proto_init() }
func file_pkg_auth_proto_auth_proto_init() {
	if File_pkg_auth_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_auth_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAuthID); i {
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
		file_pkg_auth_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAuth); i {
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
		file_pkg_auth_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAuthLogIn); i {
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
		file_pkg_auth_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLogIn); i {
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
		file_pkg_auth_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAuthPath); i {
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
		file_pkg_auth_proto_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAuthPath); i {
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
		file_pkg_auth_proto_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreateAuth); i {
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
		file_pkg_auth_proto_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUpdateAuth); i {
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
			RawDescriptor: file_pkg_auth_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_auth_proto_auth_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_auth_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_auth_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_auth_proto = out.File
	file_pkg_auth_proto_auth_proto_rawDesc = nil
	file_pkg_auth_proto_auth_proto_goTypes = nil
	file_pkg_auth_proto_auth_proto_depIdxs = nil
}
