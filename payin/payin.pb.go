// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.17.3
// source: payin.proto

package payin

import (
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

type OptionQueryPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page             int32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize         int32   `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OperationID      int64   `protobuf:"varint,3,opt,name=operationID,proto3" json:"operationID,omitempty"`
	PayinOrderID     int64   `protobuf:"varint,4,opt,name=payinOrderID,proto3" json:"payinOrderID,omitempty"`
	Statuses         []int32 `protobuf:"varint,5,rep,packed,name=Statuses,proto3" json:"Statuses,omitempty"` // [0, 1, 2] 数组格式
	Timestamp        int64   `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	RequestCount     int32   `protobuf:"varint,7,opt,name=requestCount,proto3" json:"requestCount,omitempty"`
	ResponseHttpCode string  `protobuf:"bytes,8,opt,name=responseHttpCode,proto3" json:"responseHttpCode,omitempty"`
	Msg              string  `protobuf:"bytes,9,opt,name=msg,proto3" json:"msg,omitempty"`
	IsPushed         int64   `protobuf:"varint,10,opt,name=isPushed,proto3" json:"isPushed,omitempty"`
}

func (x *OptionQueryPageRequest) Reset() {
	*x = OptionQueryPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionQueryPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionQueryPageRequest) ProtoMessage() {}

func (x *OptionQueryPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionQueryPageRequest.ProtoReflect.Descriptor instead.
func (*OptionQueryPageRequest) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{0}
}

func (x *OptionQueryPageRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OptionQueryPageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OptionQueryPageRequest) GetOperationID() int64 {
	if x != nil {
		return x.OperationID
	}
	return 0
}

func (x *OptionQueryPageRequest) GetPayinOrderID() int64 {
	if x != nil {
		return x.PayinOrderID
	}
	return 0
}

func (x *OptionQueryPageRequest) GetStatuses() []int32 {
	if x != nil {
		return x.Statuses
	}
	return nil
}

func (x *OptionQueryPageRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *OptionQueryPageRequest) GetRequestCount() int32 {
	if x != nil {
		return x.RequestCount
	}
	return 0
}

func (x *OptionQueryPageRequest) GetResponseHttpCode() string {
	if x != nil {
		return x.ResponseHttpCode
	}
	return ""
}

func (x *OptionQueryPageRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OptionQueryPageRequest) GetIsPushed() int64 {
	if x != nil {
		return x.IsPushed
	}
	return 0
}

type OptionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationID      int64   `protobuf:"varint,3,opt,name=operationID,proto3" json:"operationID,omitempty"`
	PayinOrderID     int64   `protobuf:"varint,4,opt,name=payinOrderID,proto3" json:"payinOrderID,omitempty"`
	Statuses         []int32 `protobuf:"varint,5,rep,packed,name=Statuses,proto3" json:"Statuses,omitempty"` // [0, 1, 2] 数组格式
	Timestamp        int64   `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	RequestCount     int32   `protobuf:"varint,7,opt,name=requestCount,proto3" json:"requestCount,omitempty"`
	ResponseHttpCode string  `protobuf:"bytes,8,opt,name=responseHttpCode,proto3" json:"responseHttpCode,omitempty"`
	Msg              string  `protobuf:"bytes,9,opt,name=msg,proto3" json:"msg,omitempty"`
	IsPushed         int64   `protobuf:"varint,10,opt,name=isPushed,proto3" json:"isPushed,omitempty"`
}

func (x *OptionItem) Reset() {
	*x = OptionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionItem) ProtoMessage() {}

func (x *OptionItem) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionItem.ProtoReflect.Descriptor instead.
func (*OptionItem) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{1}
}

func (x *OptionItem) GetOperationID() int64 {
	if x != nil {
		return x.OperationID
	}
	return 0
}

func (x *OptionItem) GetPayinOrderID() int64 {
	if x != nil {
		return x.PayinOrderID
	}
	return 0
}

func (x *OptionItem) GetStatuses() []int32 {
	if x != nil {
		return x.Statuses
	}
	return nil
}

func (x *OptionItem) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *OptionItem) GetRequestCount() int32 {
	if x != nil {
		return x.RequestCount
	}
	return 0
}

func (x *OptionItem) GetResponseHttpCode() string {
	if x != nil {
		return x.ResponseHttpCode
	}
	return ""
}

func (x *OptionItem) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OptionItem) GetIsPushed() int64 {
	if x != nil {
		return x.IsPushed
	}
	return 0
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize    int32 `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Page        int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	CurrentPage int32 `protobuf:"varint,3,opt,name=currentPage,proto3" json:"currentPage,omitempty"`
	Total       int32 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{2}
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Pagination) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type OptionQueryPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` // 返回码
	Msg        string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`    // 错误信息
	Items      []*OptionItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Pagination *Pagination   `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *OptionQueryPageResponse) Reset() {
	*x = OptionQueryPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionQueryPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionQueryPageResponse) ProtoMessage() {}

func (x *OptionQueryPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionQueryPageResponse.ProtoReflect.Descriptor instead.
func (*OptionQueryPageResponse) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{3}
}

func (x *OptionQueryPageResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OptionQueryPageResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OptionQueryPageResponse) GetItems() []*OptionItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OptionQueryPageResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type OptionLogQueryPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionId int64  `protobuf:"varint,1,opt,name=optionId,proto3" json:"optionId,omitempty"`
	OrderId  string `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *OptionLogQueryPageRequest) Reset() {
	*x = OptionLogQueryPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionLogQueryPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionLogQueryPageRequest) ProtoMessage() {}

func (x *OptionLogQueryPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionLogQueryPageRequest.ProtoReflect.Descriptor instead.
func (*OptionLogQueryPageRequest) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{4}
}

func (x *OptionLogQueryPageRequest) GetOptionId() int64 {
	if x != nil {
		return x.OptionId
	}
	return 0
}

func (x *OptionLogQueryPageRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OptionLogQueryPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` // 返回码
	Msg        string           `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`    // 错误信息
	Items      []*OptionLogItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Pagination *Pagination      `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *OptionLogQueryPageResponse) Reset() {
	*x = OptionLogQueryPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionLogQueryPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionLogQueryPageResponse) ProtoMessage() {}

func (x *OptionLogQueryPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionLogQueryPageResponse.ProtoReflect.Descriptor instead.
func (*OptionLogQueryPageResponse) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{5}
}

func (x *OptionLogQueryPageResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OptionLogQueryPageResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OptionLogQueryPageResponse) GetItems() []*OptionLogItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *OptionLogQueryPageResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type OptionLogItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OptionLogItem) Reset() {
	*x = OptionLogItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionLogItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionLogItem) ProtoMessage() {}

func (x *OptionLogItem) ProtoReflect() protoreflect.Message {
	mi := &file_payin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionLogItem.ProtoReflect.Descriptor instead.
func (*OptionLogItem) Descriptor() ([]byte, []int) {
	return file_payin_proto_rawDescGZIP(), []int{6}
}

var File_payin_proto protoreflect.FileDescriptor

var file_payin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x22, 0xc6, 0x02, 0x0a, 0x16, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x69, 0x6e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70,
	0x61, 0x79, 0x69, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x74, 0x74,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x50, 0x75, 0x73,
	0x68, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x73,
	0x68, 0x65, 0x64, 0x22, 0x8a, 0x02, 0x0a, 0x0a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x69,
	0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x74, 0x74, 0x70, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x50, 0x75, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x73, 0x68, 0x65, 0x64,
	0x22, 0x74, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xa7, 0x01, 0x0a, 0x17, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x51, 0x0a, 0x19, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x1a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x0f, 0x0a, 0x0d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67,
	0x49, 0x74, 0x65, 0x6d, 0x32, 0xd0, 0x01, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x12, 0x5e,
	0x0a, 0x0f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x61, 0x79, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67,
	0x0a, 0x12, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x26, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79,
	0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x61,
	0x79, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payin_proto_rawDescOnce sync.Once
	file_payin_proto_rawDescData = file_payin_proto_rawDesc
)

func file_payin_proto_rawDescGZIP() []byte {
	file_payin_proto_rawDescOnce.Do(func() {
		file_payin_proto_rawDescData = protoimpl.X.CompressGZIP(file_payin_proto_rawDescData)
	})
	return file_payin_proto_rawDescData
}

var file_payin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_payin_proto_goTypes = []interface{}{
	(*OptionQueryPageRequest)(nil),     // 0: order.payin.OptionQueryPageRequest
	(*OptionItem)(nil),                 // 1: order.payin.OptionItem
	(*Pagination)(nil),                 // 2: order.payin.Pagination
	(*OptionQueryPageResponse)(nil),    // 3: order.payin.OptionQueryPageResponse
	(*OptionLogQueryPageRequest)(nil),  // 4: order.payin.OptionLogQueryPageRequest
	(*OptionLogQueryPageResponse)(nil), // 5: order.payin.OptionLogQueryPageResponse
	(*OptionLogItem)(nil),              // 6: order.payin.OptionLogItem
}
var file_payin_proto_depIdxs = []int32{
	1, // 0: order.payin.OptionQueryPageResponse.items:type_name -> order.payin.OptionItem
	2, // 1: order.payin.OptionQueryPageResponse.pagination:type_name -> order.payin.Pagination
	6, // 2: order.payin.OptionLogQueryPageResponse.items:type_name -> order.payin.OptionLogItem
	2, // 3: order.payin.OptionLogQueryPageResponse.pagination:type_name -> order.payin.Pagination
	0, // 4: order.payin.payin.OptionQueryPage:input_type -> order.payin.OptionQueryPageRequest
	4, // 5: order.payin.payin.OptionLogQueryPage:input_type -> order.payin.OptionLogQueryPageRequest
	3, // 6: order.payin.payin.OptionQueryPage:output_type -> order.payin.OptionQueryPageResponse
	5, // 7: order.payin.payin.OptionLogQueryPage:output_type -> order.payin.OptionLogQueryPageResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_payin_proto_init() }
func file_payin_proto_init() {
	if File_payin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionQueryPageRequest); i {
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
		file_payin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionItem); i {
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
		file_payin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_payin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionQueryPageResponse); i {
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
		file_payin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionLogQueryPageRequest); i {
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
		file_payin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionLogQueryPageResponse); i {
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
		file_payin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionLogItem); i {
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
			RawDescriptor: file_payin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payin_proto_goTypes,
		DependencyIndexes: file_payin_proto_depIdxs,
		MessageInfos:      file_payin_proto_msgTypes,
	}.Build()
	File_payin_proto = out.File
	file_payin_proto_rawDesc = nil
	file_payin_proto_goTypes = nil
	file_payin_proto_depIdxs = nil
}