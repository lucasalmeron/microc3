// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: pkg/querypoints/proto/querypoints.proto

package go_micro_service_querypoints

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Say string `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSay() string {
	if x != nil {
		return x.Say
	}
	return ""
}

type RequestQueryPointID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestQueryPointID) Reset() {
	*x = RequestQueryPointID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestQueryPointID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestQueryPointID) ProtoMessage() {}

func (x *RequestQueryPointID) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestQueryPointID.ProtoReflect.Descriptor instead.
func (*RequestQueryPointID) Descriptor() ([]byte, []int) {
	return file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP(), []int{1}
}

func (x *RequestQueryPointID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResponseQueryPointsArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryPoints []*ResponseQueryPoint `protobuf:"bytes,1,rep,name=queryPoints,proto3" json:"queryPoints,omitempty"`
}

func (x *ResponseQueryPointsArray) Reset() {
	*x = ResponseQueryPointsArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseQueryPointsArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseQueryPointsArray) ProtoMessage() {}

func (x *ResponseQueryPointsArray) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseQueryPointsArray.ProtoReflect.Descriptor instead.
func (*ResponseQueryPointsArray) Descriptor() ([]byte, []int) {
	return file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseQueryPointsArray) GetQueryPoints() []*ResponseQueryPoint {
	if x != nil {
		return x.QueryPoints
	}
	return nil
}

type RequestPageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber      int64                        `protobuf:"varint,1,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	RegistersNumber int64                        `protobuf:"varint,2,opt,name=registersNumber,proto3" json:"registersNumber,omitempty"`
	OrderBy         *RequestPageOptions_Filter   `protobuf:"bytes,3,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
	Filters         []*RequestPageOptions_Filter `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *RequestPageOptions) Reset() {
	*x = RequestPageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPageOptions) ProtoMessage() {}

func (x *RequestPageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPageOptions.ProtoReflect.Descriptor instead.
func (*RequestPageOptions) Descriptor() ([]byte, []int) {
	return file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP(), []int{3}
}

func (x *RequestPageOptions) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *RequestPageOptions) GetRegistersNumber() int64 {
	if x != nil {
		return x.RegistersNumber
	}
	return 0
}

func (x *RequestPageOptions) GetOrderBy() *RequestPageOptions_Filter {
	if x != nil {
		return x.OrderBy
	}
	return nil
}

func (x *RequestPageOptions) GetFilters() []*RequestPageOptions_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

type ResponsePage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length        int64                 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	PageNumber    int64                 `protobuf:"varint,2,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	NumberOfPages int64                 `protobuf:"varint,3,opt,name=numberOfPages,proto3" json:"numberOfPages,omitempty"`
	Data          []*ResponseQueryPoint `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponsePage) Reset() {
	*x = ResponsePage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePage) ProtoMessage() {}

func (x *ResponsePage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePage.ProtoReflect.Descriptor instead.
func (*ResponsePage) Descriptor() ([]byte, []int) {
	return file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP(), []int{4}
}

func (x *ResponsePage) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *ResponsePage) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ResponsePage) GetNumberOfPages() int64 {
	if x != nil {
		return x.NumberOfPages
	}
	return 0
}

func (x *ResponsePage) GetData() []*ResponseQueryPoint {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseQueryPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone        string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address      string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	District     string   `protobuf:"bytes,5,opt,name=district,proto3" json:"district,omitempty"`
	Department   string   `protobuf:"bytes,6,opt,name=department,proto3" json:"department,omitempty"`
	Responsibles []string `protobuf:"bytes,7,rep,name=responsibles,proto3" json:"responsibles,omitempty"`
	Actions      []string `protobuf:"bytes,8,rep,name=actions,proto3" json:"actions,omitempty"`
	CreatedAt    int64    `protobuf:"varint,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ModifiedAt   int64    `protobuf:"varint,10,opt,name=modifiedAt,proto3" json:"modifiedAt,omitempty"`
	DeletedAt    int64    `protobuf:"varint,11,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *ResponseQueryPoint) Reset() {
	*x = ResponseQueryPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseQueryPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseQueryPoint) ProtoMessage() {}

func (x *ResponseQueryPoint) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseQueryPoint.ProtoReflect.Descriptor instead.
func (*ResponseQueryPoint) Descriptor() ([]byte, []int) {
	return file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseQueryPoint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseQueryPoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResponseQueryPoint) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ResponseQueryPoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ResponseQueryPoint) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *ResponseQueryPoint) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *ResponseQueryPoint) GetResponsibles() []string {
	if x != nil {
		return x.Responsibles
	}
	return nil
}

func (x *ResponseQueryPoint) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *ResponseQueryPoint) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ResponseQueryPoint) GetModifiedAt() int64 {
	if x != nil {
		return x.ModifiedAt
	}
	return 0
}

func (x *ResponseQueryPoint) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type RequestCreateQueryPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone        string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Address      string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	District     string   `protobuf:"bytes,4,opt,name=district,proto3" json:"district,omitempty"`
	Department   string   `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
	Responsibles []string `protobuf:"bytes,6,rep,name=responsibles,proto3" json:"responsibles,omitempty"`
	Actions      []string `protobuf:"bytes,7,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *RequestCreateQueryPoint) Reset() {
	*x = RequestCreateQueryPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreateQueryPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreateQueryPoint) ProtoMessage() {}

func (x *RequestCreateQueryPoint) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreateQueryPoint.ProtoReflect.Descriptor instead.
func (*RequestCreateQueryPoint) Descriptor() ([]byte, []int) {
	return file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP(), []int{6}
}

func (x *RequestCreateQueryPoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestCreateQueryPoint) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RequestCreateQueryPoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RequestCreateQueryPoint) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *RequestCreateQueryPoint) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *RequestCreateQueryPoint) GetResponsibles() []string {
	if x != nil {
		return x.Responsibles
	}
	return nil
}

func (x *RequestCreateQueryPoint) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type RequestUpdateQueryPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone        string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address      string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	District     string   `protobuf:"bytes,5,opt,name=district,proto3" json:"district,omitempty"`
	Department   string   `protobuf:"bytes,6,opt,name=department,proto3" json:"department,omitempty"`
	Responsibles []string `protobuf:"bytes,7,rep,name=responsibles,proto3" json:"responsibles,omitempty"`
	Actions      []string `protobuf:"bytes,8,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *RequestUpdateQueryPoint) Reset() {
	*x = RequestUpdateQueryPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUpdateQueryPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUpdateQueryPoint) ProtoMessage() {}

func (x *RequestUpdateQueryPoint) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUpdateQueryPoint.ProtoReflect.Descriptor instead.
func (*RequestUpdateQueryPoint) Descriptor() ([]byte, []int) {
	return file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP(), []int{7}
}

func (x *RequestUpdateQueryPoint) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RequestUpdateQueryPoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestUpdateQueryPoint) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RequestUpdateQueryPoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RequestUpdateQueryPoint) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *RequestUpdateQueryPoint) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *RequestUpdateQueryPoint) GetResponsibles() []string {
	if x != nil {
		return x.Responsibles
	}
	return nil
}

func (x *RequestUpdateQueryPoint) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type RequestPageOptions_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *RequestPageOptions_Filter) Reset() {
	*x = RequestPageOptions_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPageOptions_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPageOptions_Filter) ProtoMessage() {}

func (x *RequestPageOptions_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_querypoints_proto_querypoints_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPageOptions_Filter.ProtoReflect.Descriptor instead.
func (*RequestPageOptions_Filter) Descriptor() ([]byte, []int) {
	return file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP(), []int{3, 0}
}

func (x *RequestPageOptions_Filter) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *RequestPageOptions_Filter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pkg_querypoints_proto_querypoints_proto protoreflect.FileDescriptor

var file_pkg_querypoints_proto_querypoints_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x61,
	0x79, 0x22, 0x25, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x12, 0x52, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0xba, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x28, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x73, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x51, 0x0a, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a,
	0x34, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbe, 0x02, 0x0a, 0x12, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x17,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe7, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32,
	0xa7, 0x05, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x6e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x5b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x73,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a,
	0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_querypoints_proto_querypoints_proto_rawDescOnce sync.Once
	file_pkg_querypoints_proto_querypoints_proto_rawDescData = file_pkg_querypoints_proto_querypoints_proto_rawDesc
)

func file_pkg_querypoints_proto_querypoints_proto_rawDescGZIP() []byte {
	file_pkg_querypoints_proto_querypoints_proto_rawDescOnce.Do(func() {
		file_pkg_querypoints_proto_querypoints_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_querypoints_proto_querypoints_proto_rawDescData)
	})
	return file_pkg_querypoints_proto_querypoints_proto_rawDescData
}

var file_pkg_querypoints_proto_querypoints_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_querypoints_proto_querypoints_proto_goTypes = []interface{}{
	(*Message)(nil),                   // 0: go.micro.service.querypoints.Message
	(*RequestQueryPointID)(nil),       // 1: go.micro.service.querypoints.RequestQueryPointID
	(*ResponseQueryPointsArray)(nil),  // 2: go.micro.service.querypoints.ResponseQueryPointsArray
	(*RequestPageOptions)(nil),        // 3: go.micro.service.querypoints.RequestPageOptions
	(*ResponsePage)(nil),              // 4: go.micro.service.querypoints.ResponsePage
	(*ResponseQueryPoint)(nil),        // 5: go.micro.service.querypoints.ResponseQueryPoint
	(*RequestCreateQueryPoint)(nil),   // 6: go.micro.service.querypoints.RequestCreateQueryPoint
	(*RequestUpdateQueryPoint)(nil),   // 7: go.micro.service.querypoints.RequestUpdateQueryPoint
	(*RequestPageOptions_Filter)(nil), // 8: go.micro.service.querypoints.RequestPageOptions.Filter
	(*empty.Empty)(nil),               // 9: google.protobuf.Empty
}
var file_pkg_querypoints_proto_querypoints_proto_depIdxs = []int32{
	5,  // 0: go.micro.service.querypoints.ResponseQueryPointsArray.queryPoints:type_name -> go.micro.service.querypoints.ResponseQueryPoint
	8,  // 1: go.micro.service.querypoints.RequestPageOptions.orderBy:type_name -> go.micro.service.querypoints.RequestPageOptions.Filter
	8,  // 2: go.micro.service.querypoints.RequestPageOptions.filters:type_name -> go.micro.service.querypoints.RequestPageOptions.Filter
	5,  // 3: go.micro.service.querypoints.ResponsePage.data:type_name -> go.micro.service.querypoints.ResponseQueryPoint
	3,  // 4: go.micro.service.querypoints.QueryPoints.GetPaginated:input_type -> go.micro.service.querypoints.RequestPageOptions
	9,  // 5: go.micro.service.querypoints.QueryPoints.GetList:input_type -> google.protobuf.Empty
	1,  // 6: go.micro.service.querypoints.QueryPoints.GetByID:input_type -> go.micro.service.querypoints.RequestQueryPointID
	6,  // 7: go.micro.service.querypoints.QueryPoints.Create:input_type -> go.micro.service.querypoints.RequestCreateQueryPoint
	7,  // 8: go.micro.service.querypoints.QueryPoints.Update:input_type -> go.micro.service.querypoints.RequestUpdateQueryPoint
	1,  // 9: go.micro.service.querypoints.QueryPoints.Delete:input_type -> go.micro.service.querypoints.RequestQueryPointID
	4,  // 10: go.micro.service.querypoints.QueryPoints.GetPaginated:output_type -> go.micro.service.querypoints.ResponsePage
	2,  // 11: go.micro.service.querypoints.QueryPoints.GetList:output_type -> go.micro.service.querypoints.ResponseQueryPointsArray
	5,  // 12: go.micro.service.querypoints.QueryPoints.GetByID:output_type -> go.micro.service.querypoints.ResponseQueryPoint
	5,  // 13: go.micro.service.querypoints.QueryPoints.Create:output_type -> go.micro.service.querypoints.ResponseQueryPoint
	5,  // 14: go.micro.service.querypoints.QueryPoints.Update:output_type -> go.micro.service.querypoints.ResponseQueryPoint
	5,  // 15: go.micro.service.querypoints.QueryPoints.Delete:output_type -> go.micro.service.querypoints.ResponseQueryPoint
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_querypoints_proto_querypoints_proto_init() }
func file_pkg_querypoints_proto_querypoints_proto_init() {
	if File_pkg_querypoints_proto_querypoints_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_querypoints_proto_querypoints_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_pkg_querypoints_proto_querypoints_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestQueryPointID); i {
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
		file_pkg_querypoints_proto_querypoints_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseQueryPointsArray); i {
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
		file_pkg_querypoints_proto_querypoints_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPageOptions); i {
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
		file_pkg_querypoints_proto_querypoints_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePage); i {
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
		file_pkg_querypoints_proto_querypoints_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseQueryPoint); i {
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
		file_pkg_querypoints_proto_querypoints_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreateQueryPoint); i {
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
		file_pkg_querypoints_proto_querypoints_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUpdateQueryPoint); i {
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
		file_pkg_querypoints_proto_querypoints_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPageOptions_Filter); i {
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
			RawDescriptor: file_pkg_querypoints_proto_querypoints_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_querypoints_proto_querypoints_proto_goTypes,
		DependencyIndexes: file_pkg_querypoints_proto_querypoints_proto_depIdxs,
		MessageInfos:      file_pkg_querypoints_proto_querypoints_proto_msgTypes,
	}.Build()
	File_pkg_querypoints_proto_querypoints_proto = out.File
	file_pkg_querypoints_proto_querypoints_proto_rawDesc = nil
	file_pkg_querypoints_proto_querypoints_proto_goTypes = nil
	file_pkg_querypoints_proto_querypoints_proto_depIdxs = nil
}
