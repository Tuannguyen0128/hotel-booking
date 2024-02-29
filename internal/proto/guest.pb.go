// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: guest.proto

package proto

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

type GetGuestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CitizenId string `protobuf:"bytes,2,opt,name=citizen_id,json=citizenId,proto3" json:"citizen_id,omitempty"`
	Phone     string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Page      int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Offset    int32  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetGuestsRequest) Reset() {
	*x = GetGuestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGuestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGuestsRequest) ProtoMessage() {}

func (x *GetGuestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_guest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGuestsRequest.ProtoReflect.Descriptor instead.
func (*GetGuestsRequest) Descriptor() ([]byte, []int) {
	return file_guest_proto_rawDescGZIP(), []int{0}
}

func (x *GetGuestsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetGuestsRequest) GetCitizenId() string {
	if x != nil {
		return x.CitizenId
	}
	return ""
}

func (x *GetGuestsRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetGuestsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetGuestsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetGuestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guests []*Guest `protobuf:"bytes,1,rep,name=guests,proto3" json:"guests,omitempty"`
}

func (x *GetGuestsResponse) Reset() {
	*x = GetGuestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGuestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGuestsResponse) ProtoMessage() {}

func (x *GetGuestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_guest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGuestsResponse.ProtoReflect.Descriptor instead.
func (*GetGuestsResponse) Descriptor() ([]byte, []int) {
	return file_guest_proto_rawDescGZIP(), []int{1}
}

func (x *GetGuestsResponse) GetGuests() []*Guest {
	if x != nil {
		return x.Guests
	}
	return nil
}

type DeleteGuestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGuestRequest) Reset() {
	*x = DeleteGuestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGuestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGuestRequest) ProtoMessage() {}

func (x *DeleteGuestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_guest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGuestRequest.ProtoReflect.Descriptor instead.
func (*DeleteGuestRequest) Descriptor() ([]byte, []int) {
	return file_guest_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGuestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteGuestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteGuestResponse) Reset() {
	*x = DeleteGuestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGuestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGuestResponse) ProtoMessage() {}

func (x *DeleteGuestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_guest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGuestResponse.ProtoReflect.Descriptor instead.
func (*DeleteGuestResponse) Descriptor() ([]byte, []int) {
	return file_guest_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteGuestResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type Guest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	DateOfBirth string `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Phone       string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	CitizenId   string `protobuf:"bytes,7,opt,name=citizen_id,json=citizenId,proto3" json:"citizen_id,omitempty"`
	Email       string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt   string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdateAt    string `protobuf:"bytes,10,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	DeletedAt   string `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Guest) Reset() {
	*x = Guest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Guest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guest) ProtoMessage() {}

func (x *Guest) ProtoReflect() protoreflect.Message {
	mi := &file_guest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guest.ProtoReflect.Descriptor instead.
func (*Guest) Descriptor() ([]byte, []int) {
	return file_guest_proto_rawDescGZIP(), []int{4}
}

func (x *Guest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Guest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Guest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Guest) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *Guest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Guest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Guest) GetCitizenId() string {
	if x != nil {
		return x.CitizenId
	}
	return ""
}

func (x *Guest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Guest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Guest) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

func (x *Guest) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

var File_guest_proto protoreflect.FileDescriptor

var file_guest_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3d,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x67, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x67, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x24, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0xb7, 0x02, 0x0a, 0x05, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x8a, 0x02, 0x0a,
	0x0c, 0x47, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73,
	0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x73, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x73, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x73, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_guest_proto_rawDescOnce sync.Once
	file_guest_proto_rawDescData = file_guest_proto_rawDesc
)

func file_guest_proto_rawDescGZIP() []byte {
	file_guest_proto_rawDescOnce.Do(func() {
		file_guest_proto_rawDescData = protoimpl.X.CompressGZIP(file_guest_proto_rawDescData)
	})
	return file_guest_proto_rawDescData
}

var file_guest_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_guest_proto_goTypes = []interface{}{
	(*GetGuestsRequest)(nil),    // 0: protobufs.GetGuestsRequest
	(*GetGuestsResponse)(nil),   // 1: protobufs.GetGuestsResponse
	(*DeleteGuestRequest)(nil),  // 2: protobufs.DeleteGuestRequest
	(*DeleteGuestResponse)(nil), // 3: protobufs.DeleteGuestResponse
	(*Guest)(nil),               // 4: protobufs.Guest
}
var file_guest_proto_depIdxs = []int32{
	4, // 0: protobufs.GetGuestsResponse.guests:type_name -> protobufs.Guest
	0, // 1: protobufs.GuestService.GetGuests:input_type -> protobufs.GetGuestsRequest
	4, // 2: protobufs.GuestService.CreateGuest:input_type -> protobufs.Guest
	2, // 3: protobufs.GuestService.DeleteGuest:input_type -> protobufs.DeleteGuestRequest
	4, // 4: protobufs.GuestService.UpdateGuest:input_type -> protobufs.Guest
	1, // 5: protobufs.GuestService.GetGuests:output_type -> protobufs.GetGuestsResponse
	4, // 6: protobufs.GuestService.CreateGuest:output_type -> protobufs.Guest
	3, // 7: protobufs.GuestService.DeleteGuest:output_type -> protobufs.DeleteGuestResponse
	4, // 8: protobufs.GuestService.UpdateGuest:output_type -> protobufs.Guest
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_guest_proto_init() }
func file_guest_proto_init() {
	if File_guest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_guest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGuestsRequest); i {
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
		file_guest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGuestsResponse); i {
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
		file_guest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGuestRequest); i {
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
		file_guest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGuestResponse); i {
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
		file_guest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Guest); i {
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
			RawDescriptor: file_guest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_guest_proto_goTypes,
		DependencyIndexes: file_guest_proto_depIdxs,
		MessageInfos:      file_guest_proto_msgTypes,
	}.Build()
	File_guest_proto = out.File
	file_guest_proto_rawDesc = nil
	file_guest_proto_goTypes = nil
	file_guest_proto_depIdxs = nil
}
