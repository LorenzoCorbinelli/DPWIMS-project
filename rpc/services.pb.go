// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: services.proto

package rpc

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

type Ship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imo  string `protobuf:"bytes,1,opt,name=imo,proto3" json:"imo,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Ship) Reset() {
	*x = Ship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ship) ProtoMessage() {}

func (x *Ship) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ship.ProtoReflect.Descriptor instead.
func (*Ship) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{0}
}

func (x *Ship) GetImo() string {
	if x != nil {
		return x.Imo
	}
	return ""
}

func (x *Ship) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DepartingShip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imo         string `protobuf:"bytes,1,opt,name=imo,proto3" json:"imo,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Destination string `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *DepartingShip) Reset() {
	*x = DepartingShip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepartingShip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepartingShip) ProtoMessage() {}

func (x *DepartingShip) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepartingShip.ProtoReflect.Descriptor instead.
func (*DepartingShip) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{1}
}

func (x *DepartingShip) GetImo() string {
	if x != nil {
		return x.Imo
	}
	return ""
}

func (x *DepartingShip) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DepartingShip) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{2}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BunkeringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imo string `protobuf:"bytes,1,opt,name=imo,proto3" json:"imo,omitempty"`
}

func (x *BunkeringRequest) Reset() {
	*x = BunkeringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BunkeringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BunkeringRequest) ProtoMessage() {}

func (x *BunkeringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BunkeringRequest.ProtoReflect.Descriptor instead.
func (*BunkeringRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{3}
}

func (x *BunkeringRequest) GetImo() string {
	if x != nil {
		return x.Imo
	}
	return ""
}

type BunkeringReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMessage string `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Tanker       *Ship  `protobuf:"bytes,2,opt,name=tanker,proto3" json:"tanker,omitempty"`
}

func (x *BunkeringReply) Reset() {
	*x = BunkeringReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BunkeringReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BunkeringReply) ProtoMessage() {}

func (x *BunkeringReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BunkeringReply.ProtoReflect.Descriptor instead.
func (*BunkeringReply) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{4}
}

func (x *BunkeringReply) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *BunkeringReply) GetTanker() *Ship {
	if x != nil {
		return x.Tanker
	}
	return nil
}

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x2c, 0x0a, 0x04, 0x53, 0x68, 0x69, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x68, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a,
	0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x24, 0x0a, 0x10, 0x42, 0x75, 0x6e, 0x6b, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x69, 0x6d, 0x6f, 0x22, 0x58, 0x0a, 0x0e, 0x42, 0x75, 0x6e, 0x6b, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x06,
	0x74, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x52, 0x06, 0x74, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x32, 0x9e, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x07, 0x41, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x12, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x53, 0x68, 0x69, 0x70, 0x1a, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x68, 0x69, 0x70, 0x1a, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x42, 0x75, 0x6e, 0x6b, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x75, 0x6e, 0x6b, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x42, 0x75, 0x6e, 0x6b, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_proto_rawDescOnce sync.Once
	file_services_proto_rawDescData = file_services_proto_rawDesc
)

func file_services_proto_rawDescGZIP() []byte {
	file_services_proto_rawDescOnce.Do(func() {
		file_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_proto_rawDescData)
	})
	return file_services_proto_rawDescData
}

var file_services_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_proto_goTypes = []interface{}{
	(*Ship)(nil),             // 0: main.Ship
	(*DepartingShip)(nil),    // 1: main.DepartingShip
	(*Reply)(nil),            // 2: main.Reply
	(*BunkeringRequest)(nil), // 3: main.BunkeringRequest
	(*BunkeringReply)(nil),   // 4: main.BunkeringReply
}
var file_services_proto_depIdxs = []int32{
	0, // 0: main.BunkeringReply.tanker:type_name -> main.Ship
	0, // 1: main.Register.Arrival:input_type -> main.Ship
	1, // 2: main.Register.Departure:input_type -> main.DepartingShip
	3, // 3: main.Register.Bunkering:input_type -> main.BunkeringRequest
	2, // 4: main.Register.Arrival:output_type -> main.Reply
	2, // 5: main.Register.Departure:output_type -> main.Reply
	4, // 6: main.Register.Bunkering:output_type -> main.BunkeringReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ship); i {
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
		file_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepartingShip); i {
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
		file_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BunkeringRequest); i {
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
		file_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BunkeringReply); i {
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
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
		MessageInfos:      file_services_proto_msgTypes,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}
