// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pkg/event/pb/event.proto

package pb

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

type AddDisciplineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddDisciplineRequest) Reset() {
	*x = AddDisciplineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDisciplineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDisciplineRequest) ProtoMessage() {}

func (x *AddDisciplineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDisciplineRequest.ProtoReflect.Descriptor instead.
func (*AddDisciplineRequest) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{0}
}

func (x *AddDisciplineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddDisciplineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddDisciplineResponse) Reset() {
	*x = AddDisciplineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDisciplineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDisciplineResponse) ProtoMessage() {}

func (x *AddDisciplineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDisciplineResponse.ProtoReflect.Descriptor instead.
func (*AddDisciplineResponse) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{1}
}

func (x *AddDisciplineResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteDisciplineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteDisciplineRequest) Reset() {
	*x = DeleteDisciplineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDisciplineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDisciplineRequest) ProtoMessage() {}

func (x *DeleteDisciplineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDisciplineRequest.ProtoReflect.Descriptor instead.
func (*DeleteDisciplineRequest) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteDisciplineRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteDisciplineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteDisciplineResponse) Reset() {
	*x = DeleteDisciplineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDisciplineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDisciplineResponse) ProtoMessage() {}

func (x *DeleteDisciplineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDisciplineResponse.ProtoReflect.Descriptor instead.
func (*DeleteDisciplineResponse) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteDisciplineResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetDisciplinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDisciplinesRequest) Reset() {
	*x = GetDisciplinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDisciplinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDisciplinesRequest) ProtoMessage() {}

func (x *GetDisciplinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDisciplinesRequest.ProtoReflect.Descriptor instead.
func (*GetDisciplinesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{4}
}

func (x *GetDisciplinesRequest) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type DisciplineData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DisciplineData) Reset() {
	*x = DisciplineData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisciplineData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisciplineData) ProtoMessage() {}

func (x *DisciplineData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisciplineData.ProtoReflect.Descriptor instead.
func (*DisciplineData) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{5}
}

func (x *DisciplineData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DisciplineData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDisciplinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disciplines []*DisciplineData `protobuf:"bytes,1,rep,name=disciplines,proto3" json:"disciplines,omitempty"`
}

func (x *GetDisciplinesResponse) Reset() {
	*x = GetDisciplinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDisciplinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDisciplinesResponse) ProtoMessage() {}

func (x *GetDisciplinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDisciplinesResponse.ProtoReflect.Descriptor instead.
func (*GetDisciplinesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{6}
}

func (x *GetDisciplinesResponse) GetDisciplines() []*DisciplineData {
	if x != nil {
		return x.Disciplines
	}
	return nil
}

type AddEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisciplineId int64  `protobuf:"varint,1,opt,name=discipline_id,json=disciplineId,proto3" json:"discipline_id,omitempty"`
	Place        string `protobuf:"bytes,2,opt,name=place,proto3" json:"place,omitempty"`
	Date         string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	StartTime    string `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime      string `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Price        string `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	Limit        string `protobuf:"bytes,7,opt,name=limit,proto3" json:"limit,omitempty"`
	Description  string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AddEventRequest) Reset() {
	*x = AddEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventRequest) ProtoMessage() {}

func (x *AddEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventRequest.ProtoReflect.Descriptor instead.
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{7}
}

func (x *AddEventRequest) GetDisciplineId() int64 {
	if x != nil {
		return x.DisciplineId
	}
	return 0
}

func (x *AddEventRequest) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *AddEventRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *AddEventRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *AddEventRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *AddEventRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *AddEventRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *AddEventRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AddEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddEventResponse) Reset() {
	*x = AddEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventResponse) ProtoMessage() {}

func (x *AddEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventResponse.ProtoReflect.Descriptor instead.
func (*AddEventResponse) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{8}
}

func (x *AddEventResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEventRequest) Reset() {
	*x = DeleteEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRequest) ProtoMessage() {}

func (x *DeleteEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEventRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteEventResponse) Reset() {
	*x = DeleteEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventResponse) ProtoMessage() {}

func (x *DeleteEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventResponse.ProtoReflect.Descriptor instead.
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteEventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type EventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DisciplineName string `protobuf:"bytes,2,opt,name=discipline_name,json=disciplineName,proto3" json:"discipline_name,omitempty"`
	Place          string `protobuf:"bytes,3,opt,name=place,proto3" json:"place,omitempty"`
	Date           string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	StartTime      string `protobuf:"bytes,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime        string `protobuf:"bytes,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Price          string `protobuf:"bytes,7,opt,name=price,proto3" json:"price,omitempty"`
	Limit          string `protobuf:"bytes,8,opt,name=limit,proto3" json:"limit,omitempty"`
	Description    string `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *EventData) Reset() {
	*x = EventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventData) ProtoMessage() {}

func (x *EventData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventData.ProtoReflect.Descriptor instead.
func (*EventData) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{11}
}

func (x *EventData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EventData) GetDisciplineName() string {
	if x != nil {
		return x.DisciplineName
	}
	return ""
}

func (x *EventData) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *EventData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *EventData) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *EventData) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *EventData) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *EventData) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *EventData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisciplineId []int64 `protobuf:"varint,1,rep,packed,name=discipline_id,json=disciplineId,proto3" json:"discipline_id,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{12}
}

func (x *GetEventsRequest) GetDisciplineId() []int64 {
	if x != nil {
		return x.DisciplineId
	}
	return nil
}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*EventData `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_event_pb_event_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_event_pb_event_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_event_pb_event_proto_rawDescGZIP(), []int{13}
}

func (x *GetEventsResponse) GetEvents() []*EventData {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_pkg_event_pb_event_proto protoreflect.FileDescriptor

var file_pkg_event_pb_event_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x2a, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a,
	0x15, 0x41, 0x64, 0x64, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x34, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x69,
	0x70, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x34, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x51, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73,
	0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0xe6, 0x01, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xf4, 0x01,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x64,
	0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63,
	0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x0c, 0x64, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x3d, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xc6, 0x03, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x44, 0x69, 0x73,
	0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69,
	0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1c, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x69, 0x70, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08,
	0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_event_pb_event_proto_rawDescOnce sync.Once
	file_pkg_event_pb_event_proto_rawDescData = file_pkg_event_pb_event_proto_rawDesc
)

func file_pkg_event_pb_event_proto_rawDescGZIP() []byte {
	file_pkg_event_pb_event_proto_rawDescOnce.Do(func() {
		file_pkg_event_pb_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_event_pb_event_proto_rawDescData)
	})
	return file_pkg_event_pb_event_proto_rawDescData
}

var file_pkg_event_pb_event_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_pkg_event_pb_event_proto_goTypes = []interface{}{
	(*AddDisciplineRequest)(nil),     // 0: event.AddDisciplineRequest
	(*AddDisciplineResponse)(nil),    // 1: event.AddDisciplineResponse
	(*DeleteDisciplineRequest)(nil),  // 2: event.DeleteDisciplineRequest
	(*DeleteDisciplineResponse)(nil), // 3: event.DeleteDisciplineResponse
	(*GetDisciplinesRequest)(nil),    // 4: event.GetDisciplinesRequest
	(*DisciplineData)(nil),           // 5: event.DisciplineData
	(*GetDisciplinesResponse)(nil),   // 6: event.GetDisciplinesResponse
	(*AddEventRequest)(nil),          // 7: event.AddEventRequest
	(*AddEventResponse)(nil),         // 8: event.AddEventResponse
	(*DeleteEventRequest)(nil),       // 9: event.DeleteEventRequest
	(*DeleteEventResponse)(nil),      // 10: event.DeleteEventResponse
	(*EventData)(nil),                // 11: event.EventData
	(*GetEventsRequest)(nil),         // 12: event.GetEventsRequest
	(*GetEventsResponse)(nil),        // 13: event.GetEventsResponse
}
var file_pkg_event_pb_event_proto_depIdxs = []int32{
	5,  // 0: event.GetDisciplinesResponse.disciplines:type_name -> event.DisciplineData
	11, // 1: event.GetEventsResponse.events:type_name -> event.EventData
	0,  // 2: event.Event.AddDiscipline:input_type -> event.AddDisciplineRequest
	2,  // 3: event.Event.DeleteDiscipline:input_type -> event.DeleteDisciplineRequest
	4,  // 4: event.Event.GetDisciplines:input_type -> event.GetDisciplinesRequest
	7,  // 5: event.Event.AddEvent:input_type -> event.AddEventRequest
	9,  // 6: event.Event.DeleteEvent:input_type -> event.DeleteEventRequest
	12, // 7: event.Event.GetEvents:input_type -> event.GetEventsRequest
	1,  // 8: event.Event.AddDiscipline:output_type -> event.AddDisciplineResponse
	3,  // 9: event.Event.DeleteDiscipline:output_type -> event.DeleteDisciplineResponse
	6,  // 10: event.Event.GetDisciplines:output_type -> event.GetDisciplinesResponse
	8,  // 11: event.Event.AddEvent:output_type -> event.AddEventResponse
	10, // 12: event.Event.DeleteEvent:output_type -> event.DeleteEventResponse
	13, // 13: event.Event.GetEvents:output_type -> event.GetEventsResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_event_pb_event_proto_init() }
func file_pkg_event_pb_event_proto_init() {
	if File_pkg_event_pb_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_event_pb_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDisciplineRequest); i {
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
		file_pkg_event_pb_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDisciplineResponse); i {
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
		file_pkg_event_pb_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDisciplineRequest); i {
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
		file_pkg_event_pb_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDisciplineResponse); i {
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
		file_pkg_event_pb_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDisciplinesRequest); i {
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
		file_pkg_event_pb_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisciplineData); i {
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
		file_pkg_event_pb_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDisciplinesResponse); i {
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
		file_pkg_event_pb_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventRequest); i {
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
		file_pkg_event_pb_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventResponse); i {
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
		file_pkg_event_pb_event_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventRequest); i {
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
		file_pkg_event_pb_event_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventResponse); i {
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
		file_pkg_event_pb_event_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventData); i {
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
		file_pkg_event_pb_event_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
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
		file_pkg_event_pb_event_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsResponse); i {
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
			RawDescriptor: file_pkg_event_pb_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_event_pb_event_proto_goTypes,
		DependencyIndexes: file_pkg_event_pb_event_proto_depIdxs,
		MessageInfos:      file_pkg_event_pb_event_proto_msgTypes,
	}.Build()
	File_pkg_event_pb_event_proto = out.File
	file_pkg_event_pb_event_proto_rawDesc = nil
	file_pkg_event_pb_event_proto_goTypes = nil
	file_pkg_event_pb_event_proto_depIdxs = nil
}
