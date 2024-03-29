// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/dog.proto

package interact

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type DogAction int32

const (
	DogAction_UNKNOWN_DOG_ACTION DogAction = 0
	DogAction_POPIS              DogAction = 1
	DogAction_WALK               DogAction = 2
)

// Enum value maps for DogAction.
var (
	DogAction_name = map[int32]string{
		0: "UNKNOWN_DOG_ACTION",
		1: "POPIS",
		2: "WALK",
	}
	DogAction_value = map[string]int32{
		"UNKNOWN_DOG_ACTION": 0,
		"POPIS":              1,
		"WALK":               2,
	}
)

func (x DogAction) Enum() *DogAction {
	p := new(DogAction)
	*p = x
	return p
}

func (x DogAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DogAction) Descriptor() protoreflect.EnumDescriptor {
	return file_api_dog_proto_enumTypes[0].Descriptor()
}

func (DogAction) Type() protoreflect.EnumType {
	return &file_api_dog_proto_enumTypes[0]
}

func (x DogAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DogAction.Descriptor instead.
func (DogAction) EnumDescriptor() ([]byte, []int) {
	return file_api_dog_proto_rawDescGZIP(), []int{0}
}

type SetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action DogAction `protobuf:"varint,1,opt,name=action,proto3,enum=interact.DogAction" json:"action,omitempty"`
}

func (x *SetStatusRequest) Reset() {
	*x = SetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusRequest) ProtoMessage() {}

func (x *SetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusRequest.ProtoReflect.Descriptor instead.
func (*SetStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_dog_proto_rawDescGZIP(), []int{0}
}

func (x *SetStatusRequest) GetAction() DogAction {
	if x != nil {
		return x.Action
	}
	return DogAction_UNKNOWN_DOG_ACTION
}

type SetStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *SetStatusReply_Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SetStatusReply) Reset() {
	*x = SetStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusReply) ProtoMessage() {}

func (x *SetStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_dog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusReply.ProtoReflect.Descriptor instead.
func (*SetStatusReply) Descriptor() ([]byte, []int) {
	return file_api_dog_proto_rawDescGZIP(), []int{1}
}

func (x *SetStatusReply) GetResult() *SetStatusReply_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type StatusListrequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *StatusListrequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *StatusListrequest) Reset() {
	*x = StatusListrequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusListrequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusListrequest) ProtoMessage() {}

func (x *StatusListrequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusListrequest.ProtoReflect.Descriptor instead.
func (*StatusListrequest) Descriptor() ([]byte, []int) {
	return file_api_dog_proto_rawDescGZIP(), []int{2}
}

func (x *StatusListrequest) GetFilter() *StatusListrequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type StatusListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*StatusListReply_Action `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *StatusListReply) Reset() {
	*x = StatusListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusListReply) ProtoMessage() {}

func (x *StatusListReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_dog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusListReply.ProtoReflect.Descriptor instead.
func (*StatusListReply) Descriptor() ([]byte, []int) {
	return file_api_dog_proto_rawDescGZIP(), []int{3}
}

func (x *StatusListReply) GetResult() []*StatusListReply_Action {
	if x != nil {
		return x.Result
	}
	return nil
}

type SetStatusReply_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created bool `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *SetStatusReply_Result) Reset() {
	*x = SetStatusReply_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStatusReply_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusReply_Result) ProtoMessage() {}

func (x *SetStatusReply_Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_dog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusReply_Result.ProtoReflect.Descriptor instead.
func (*SetStatusReply_Result) Descriptor() ([]byte, []int) {
	return file_api_dog_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SetStatusReply_Result) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

type StatusListrequest_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date    *StatusListrequest_Filter_Date `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Actions []DogAction                    `protobuf:"varint,2,rep,packed,name=actions,proto3,enum=interact.DogAction" json:"actions,omitempty"`
}

func (x *StatusListrequest_Filter) Reset() {
	*x = StatusListrequest_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusListrequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusListrequest_Filter) ProtoMessage() {}

func (x *StatusListrequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_api_dog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusListrequest_Filter.ProtoReflect.Descriptor instead.
func (*StatusListrequest_Filter) Descriptor() ([]byte, []int) {
	return file_api_dog_proto_rawDescGZIP(), []int{2, 0}
}

func (x *StatusListrequest_Filter) GetDate() *StatusListrequest_Filter_Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *StatusListrequest_Filter) GetActions() []DogAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

type StatusListrequest_Filter_Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *StatusListrequest_Filter_Date) Reset() {
	*x = StatusListrequest_Filter_Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusListrequest_Filter_Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusListrequest_Filter_Date) ProtoMessage() {}

func (x *StatusListrequest_Filter_Date) ProtoReflect() protoreflect.Message {
	mi := &file_api_dog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusListrequest_Filter_Date.ProtoReflect.Descriptor instead.
func (*StatusListrequest_Filter_Date) Descriptor() ([]byte, []int) {
	return file_api_dog_proto_rawDescGZIP(), []int{2, 0, 0}
}

func (x *StatusListrequest_Filter_Date) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *StatusListrequest_Filter_Date) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

type StatusListReply_Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Action DogAction              `protobuf:"varint,2,opt,name=action,proto3,enum=interact.DogAction" json:"action,omitempty"`
}

func (x *StatusListReply_Action) Reset() {
	*x = StatusListReply_Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusListReply_Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusListReply_Action) ProtoMessage() {}

func (x *StatusListReply_Action) ProtoReflect() protoreflect.Message {
	mi := &file_api_dog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusListReply_Action.ProtoReflect.Descriptor instead.
func (*StatusListReply_Action) Descriptor() ([]byte, []int) {
	return file_api_dog_proto_rawDescGZIP(), []int{3, 0}
}

func (x *StatusListReply_Action) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *StatusListReply_Action) GetAction() DogAction {
	if x != nil {
		return x.Action
	}
	return DogAction_UNKNOWN_DOG_ACTION
}

var File_api_dog_proto protoreflect.FileDescriptor

var file_api_dog_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x10, 0x53, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x44, 0x6f, 0x67, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x22, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0xaa, 0x02, 0x0a, 0x11, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0xd8, 0x01, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x44, 0x6f, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x62, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x65, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x44, 0x6f, 0x67, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x38, 0x0a, 0x09,
	0x44, 0x6f, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x44, 0x4f, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x4f, 0x50, 0x49, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x57, 0x41, 0x4c, 0x4b, 0x10, 0x02, 0x32, 0x8e, 0x01, 0x0a, 0x03, 0x44, 0x6f, 0x67, 0x12, 0x41,
	0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x44, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x2d, 0x73, 0x6b, 0x79, 0x2f, 0x64, 0x6f,
	0x67, 0x5f, 0x62, 0x6f, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_dog_proto_rawDescOnce sync.Once
	file_api_dog_proto_rawDescData = file_api_dog_proto_rawDesc
)

func file_api_dog_proto_rawDescGZIP() []byte {
	file_api_dog_proto_rawDescOnce.Do(func() {
		file_api_dog_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_dog_proto_rawDescData)
	})
	return file_api_dog_proto_rawDescData
}

var file_api_dog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_dog_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_dog_proto_goTypes = []interface{}{
	(DogAction)(0),                        // 0: interact.DogAction
	(*SetStatusRequest)(nil),              // 1: interact.SetStatusRequest
	(*SetStatusReply)(nil),                // 2: interact.SetStatusReply
	(*StatusListrequest)(nil),             // 3: interact.StatusListrequest
	(*StatusListReply)(nil),               // 4: interact.StatusListReply
	(*SetStatusReply_Result)(nil),         // 5: interact.SetStatusReply.Result
	(*StatusListrequest_Filter)(nil),      // 6: interact.StatusListrequest.Filter
	(*StatusListrequest_Filter_Date)(nil), // 7: interact.StatusListrequest.Filter.Date
	(*StatusListReply_Action)(nil),        // 8: interact.StatusListReply.Action
	(*timestamppb.Timestamp)(nil),         // 9: google.protobuf.Timestamp
}
var file_api_dog_proto_depIdxs = []int32{
	0,  // 0: interact.SetStatusRequest.action:type_name -> interact.DogAction
	5,  // 1: interact.SetStatusReply.result:type_name -> interact.SetStatusReply.Result
	6,  // 2: interact.StatusListrequest.filter:type_name -> interact.StatusListrequest.Filter
	8,  // 3: interact.StatusListReply.result:type_name -> interact.StatusListReply.Action
	7,  // 4: interact.StatusListrequest.Filter.date:type_name -> interact.StatusListrequest.Filter.Date
	0,  // 5: interact.StatusListrequest.Filter.actions:type_name -> interact.DogAction
	9,  // 6: interact.StatusListrequest.Filter.Date.from:type_name -> google.protobuf.Timestamp
	9,  // 7: interact.StatusListrequest.Filter.Date.to:type_name -> google.protobuf.Timestamp
	9,  // 8: interact.StatusListReply.Action.date:type_name -> google.protobuf.Timestamp
	0,  // 9: interact.StatusListReply.Action.action:type_name -> interact.DogAction
	1,  // 10: interact.Dog.SetStatus:input_type -> interact.SetStatusRequest
	3,  // 11: interact.Dog.StatusList:input_type -> interact.StatusListrequest
	2,  // 12: interact.Dog.SetStatus:output_type -> interact.SetStatusReply
	4,  // 13: interact.Dog.StatusList:output_type -> interact.StatusListReply
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_dog_proto_init() }
func file_api_dog_proto_init() {
	if File_api_dog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_dog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStatusRequest); i {
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
		file_api_dog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStatusReply); i {
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
		file_api_dog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusListrequest); i {
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
		file_api_dog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusListReply); i {
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
		file_api_dog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStatusReply_Result); i {
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
		file_api_dog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusListrequest_Filter); i {
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
		file_api_dog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusListrequest_Filter_Date); i {
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
		file_api_dog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusListReply_Action); i {
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
			RawDescriptor: file_api_dog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_dog_proto_goTypes,
		DependencyIndexes: file_api_dog_proto_depIdxs,
		EnumInfos:         file_api_dog_proto_enumTypes,
		MessageInfos:      file_api_dog_proto_msgTypes,
	}.Build()
	File_api_dog_proto = out.File
	file_api_dog_proto_rawDesc = nil
	file_api_dog_proto_goTypes = nil
	file_api_dog_proto_depIdxs = nil
}
