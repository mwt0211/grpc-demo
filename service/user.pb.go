// 指定的当前proto语法的版本，有2和3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: Pbfile/user.proto

// 指定等会文件生成出来的package

package service

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Age       int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Score     int32    `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	Height    float32  `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty"`
	StuNumber int64    `protobuf:"varint,5,opt,name=stu_number,json=stuNumber,proto3" json:"stu_number,omitempty"`
	Password  *string  `protobuf:"bytes,6,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Addresses []string `protobuf:"bytes,7,rep,name=addresses,proto3" json:"addresses,omitempty"`
	IdCard    *string  `protobuf:"bytes,8,opt,name=IdCard,proto3,oneof" json:"IdCard,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_Pbfile_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *User) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *User) GetStuNumber() int64 {
	if x != nil {
		return x.StuNumber
	}
	return 0
}

func (x *User) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *User) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *User) GetIdCard() string {
	if x != nil && x.IdCard != nil {
		return *x.IdCard
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Age      int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Password *string  `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Address  []string `protobuf:"bytes,4,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[1]
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
	return file_Pbfile_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *UserRequest) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Age      int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Password *string  `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Address  []string `protobuf:"bytes,4,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[2]
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
	return file_Pbfile_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserResponse) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserResponse) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *UserResponse) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

type PersonInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*PersonInfo_Person `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *PersonInfo) Reset() {
	*x = PersonInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonInfo) ProtoMessage() {}

func (x *PersonInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonInfo.ProtoReflect.Descriptor instead.
func (*PersonInfo) Descriptor() ([]byte, []int) {
	return file_Pbfile_user_proto_rawDescGZIP(), []int{3}
}

func (x *PersonInfo) GetInfo() []*PersonInfo_Person {
	if x != nil {
		return x.Info
	}
	return nil
}

type PersonMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *PersonInfo_Person `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PersonMessage) Reset() {
	*x = PersonMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonMessage) ProtoMessage() {}

func (x *PersonMessage) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonMessage.ProtoReflect.Descriptor instead.
func (*PersonMessage) Descriptor() ([]byte, []int) {
	return file_Pbfile_user_proto_rawDescGZIP(), []int{4}
}

func (x *PersonMessage) GetInfo() *PersonInfo_Person {
	if x != nil {
		return x.Info
	}
	return nil
}

type Grandpa struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Grandpa) Reset() {
	*x = Grandpa{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grandpa) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grandpa) ProtoMessage() {}

func (x *Grandpa) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grandpa.ProtoReflect.Descriptor instead.
func (*Grandpa) Descriptor() ([]byte, []int) {
	return file_Pbfile_user_proto_rawDescGZIP(), []int{5}
}

type PersonInfo_Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Height int32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Weight []int32 `protobuf:"varint,3,rep,packed,name=weight,proto3" json:"weight,omitempty"`
}

func (x *PersonInfo_Person) Reset() {
	*x = PersonInfo_Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonInfo_Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonInfo_Person) ProtoMessage() {}

func (x *PersonInfo_Person) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonInfo_Person.ProtoReflect.Descriptor instead.
func (*PersonInfo_Person) Descriptor() ([]byte, []int) {
	return file_Pbfile_user_proto_rawDescGZIP(), []int{3, 0}
}

func (x *PersonInfo_Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonInfo_Person) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *PersonInfo_Person) GetWeight() []int32 {
	if x != nil {
		return x.Weight
	}
	return nil
}

type Grandpa_Father struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Grandpa_Father) Reset() {
	*x = Grandpa_Father{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grandpa_Father) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grandpa_Father) ProtoMessage() {}

func (x *Grandpa_Father) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grandpa_Father.ProtoReflect.Descriptor instead.
func (*Grandpa_Father) Descriptor() ([]byte, []int) {
	return file_Pbfile_user_proto_rawDescGZIP(), []int{5, 0}
}

type Grandpa_Uncle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Grandpa_Uncle) Reset() {
	*x = Grandpa_Uncle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grandpa_Uncle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grandpa_Uncle) ProtoMessage() {}

func (x *Grandpa_Uncle) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grandpa_Uncle.ProtoReflect.Descriptor instead.
func (*Grandpa_Uncle) Descriptor() ([]byte, []int) {
	return file_Pbfile_user_proto_rawDescGZIP(), []int{5, 1}
}

type Grandpa_FatherSon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Grandpa_FatherSon) Reset() {
	*x = Grandpa_FatherSon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grandpa_FatherSon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grandpa_FatherSon) ProtoMessage() {}

func (x *Grandpa_FatherSon) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grandpa_FatherSon.ProtoReflect.Descriptor instead.
func (*Grandpa_FatherSon) Descriptor() ([]byte, []int) {
	return file_Pbfile_user_proto_rawDescGZIP(), []int{5, 0, 0}
}

func (x *Grandpa_FatherSon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Grandpa_FatherSon) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Grandpa_Uncle_Son struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Grandpa_Uncle_Son) Reset() {
	*x = Grandpa_Uncle_Son{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Pbfile_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grandpa_Uncle_Son) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grandpa_Uncle_Son) ProtoMessage() {}

func (x *Grandpa_Uncle_Son) ProtoReflect() protoreflect.Message {
	mi := &file_Pbfile_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grandpa_Uncle_Son.ProtoReflect.Descriptor instead.
func (*Grandpa_Uncle_Son) Descriptor() ([]byte, []int) {
	return file_Pbfile_user_proto_rawDescGZIP(), []int{5, 1, 0}
}

func (x *Grandpa_Uncle_Son) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Grandpa_Uncle_Son) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_Pbfile_user_proto protoreflect.FileDescriptor

var file_Pbfile_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x50, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xf5, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x75, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x1b, 0x0a, 0x06, 0x49, 0x64, 0x43, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x06, 0x49, 0x64, 0x43, 0x61, 0x72, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x49, 0x64,
	0x43, 0x61, 0x72, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x8a, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x1a, 0x4c, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3f,
	0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x76, 0x0a, 0x07, 0x47, 0x72, 0x61, 0x6e, 0x64, 0x70, 0x61, 0x1a, 0x35, 0x0a, 0x06, 0x46, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x1a, 0x2b, 0x0a, 0x03, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x1a, 0x34, 0x0a, 0x05, 0x55, 0x6e, 0x63, 0x6c, 0x65, 0x1a, 0x2b, 0x0a, 0x03, 0x53, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Pbfile_user_proto_rawDescOnce sync.Once
	file_Pbfile_user_proto_rawDescData = file_Pbfile_user_proto_rawDesc
)

func file_Pbfile_user_proto_rawDescGZIP() []byte {
	file_Pbfile_user_proto_rawDescOnce.Do(func() {
		file_Pbfile_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_Pbfile_user_proto_rawDescData)
	})
	return file_Pbfile_user_proto_rawDescData
}

var file_Pbfile_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_Pbfile_user_proto_goTypes = []interface{}{
	(*User)(nil),              // 0: service.User
	(*UserRequest)(nil),       // 1: service.UserRequest
	(*UserResponse)(nil),      // 2: service.UserResponse
	(*PersonInfo)(nil),        // 3: service.PersonInfo
	(*PersonMessage)(nil),     // 4: service.PersonMessage
	(*Grandpa)(nil),           // 5: service.Grandpa
	(*PersonInfo_Person)(nil), // 6: service.PersonInfo.Person
	(*Grandpa_Father)(nil),    // 7: service.Grandpa.Father
	(*Grandpa_Uncle)(nil),     // 8: service.Grandpa.Uncle
	(*Grandpa_FatherSon)(nil), // 9: service.Grandpa.Father.son
	(*Grandpa_Uncle_Son)(nil), // 10: service.Grandpa.Uncle.Son
}
var file_Pbfile_user_proto_depIdxs = []int32{
	6, // 0: service.PersonInfo.info:type_name -> service.PersonInfo.Person
	6, // 1: service.PersonMessage.info:type_name -> service.PersonInfo.Person
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Pbfile_user_proto_init() }
func file_Pbfile_user_proto_init() {
	if File_Pbfile_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Pbfile_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_Pbfile_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Pbfile_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_Pbfile_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonInfo); i {
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
		file_Pbfile_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonMessage); i {
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
		file_Pbfile_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grandpa); i {
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
		file_Pbfile_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonInfo_Person); i {
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
		file_Pbfile_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grandpa_Father); i {
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
		file_Pbfile_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grandpa_Uncle); i {
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
		file_Pbfile_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grandpa_FatherSon); i {
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
		file_Pbfile_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grandpa_Uncle_Son); i {
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
	file_Pbfile_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_Pbfile_user_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_Pbfile_user_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Pbfile_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Pbfile_user_proto_goTypes,
		DependencyIndexes: file_Pbfile_user_proto_depIdxs,
		MessageInfos:      file_Pbfile_user_proto_msgTypes,
	}.Build()
	File_Pbfile_user_proto = out.File
	file_Pbfile_user_proto_rawDesc = nil
	file_Pbfile_user_proto_goTypes = nil
	file_Pbfile_user_proto_depIdxs = nil
}
