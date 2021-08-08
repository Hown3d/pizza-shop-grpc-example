// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: bakery/v1/bakery.proto

package bakery

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

type BakePizzaResponse_BakeStatus int32

const (
	BakePizzaResponse_BAKE_STATUS_UNSPECIFIED BakePizzaResponse_BakeStatus = 0
	BakePizzaResponse_BAKE_STATUS_PREPARING   BakePizzaResponse_BakeStatus = 1
	BakePizzaResponse_BAKE_STATUS_BAKING      BakePizzaResponse_BakeStatus = 2
	BakePizzaResponse_BAKE_STATUS_READY       BakePizzaResponse_BakeStatus = 3
)

// Enum value maps for BakePizzaResponse_BakeStatus.
var (
	BakePizzaResponse_BakeStatus_name = map[int32]string{
		0: "BAKE_STATUS_UNSPECIFIED",
		1: "BAKE_STATUS_PREPARING",
		2: "BAKE_STATUS_BAKING",
		3: "BAKE_STATUS_READY",
	}
	BakePizzaResponse_BakeStatus_value = map[string]int32{
		"BAKE_STATUS_UNSPECIFIED": 0,
		"BAKE_STATUS_PREPARING":   1,
		"BAKE_STATUS_BAKING":      2,
		"BAKE_STATUS_READY":       3,
	}
)

func (x BakePizzaResponse_BakeStatus) Enum() *BakePizzaResponse_BakeStatus {
	p := new(BakePizzaResponse_BakeStatus)
	*p = x
	return p
}

func (x BakePizzaResponse_BakeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BakePizzaResponse_BakeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_bakery_v1_bakery_proto_enumTypes[0].Descriptor()
}

func (BakePizzaResponse_BakeStatus) Type() protoreflect.EnumType {
	return &file_bakery_v1_bakery_proto_enumTypes[0]
}

func (x BakePizzaResponse_BakeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BakePizzaResponse_BakeStatus.Descriptor instead.
func (BakePizzaResponse_BakeStatus) EnumDescriptor() ([]byte, []int) {
	return file_bakery_v1_bakery_proto_rawDescGZIP(), []int{2, 0}
}

type BakePizzaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Toppings []*Toppings `protobuf:"bytes,1,rep,name=toppings,proto3" json:"toppings,omitempty"`
	Name     string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BakePizzaRequest) Reset() {
	*x = BakePizzaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakery_v1_bakery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BakePizzaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BakePizzaRequest) ProtoMessage() {}

func (x *BakePizzaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bakery_v1_bakery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BakePizzaRequest.ProtoReflect.Descriptor instead.
func (*BakePizzaRequest) Descriptor() ([]byte, []int) {
	return file_bakery_v1_bakery_proto_rawDescGZIP(), []int{0}
}

func (x *BakePizzaRequest) GetToppings() []*Toppings {
	if x != nil {
		return x.Toppings
	}
	return nil
}

func (x *BakePizzaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Toppings for the pizza
type Toppings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the topping
	// represents the id of the topping in the whole system
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name of the topping
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Toppings) Reset() {
	*x = Toppings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakery_v1_bakery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Toppings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Toppings) ProtoMessage() {}

func (x *Toppings) ProtoReflect() protoreflect.Message {
	mi := &file_bakery_v1_bakery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Toppings.ProtoReflect.Descriptor instead.
func (*Toppings) Descriptor() ([]byte, []int) {
	return file_bakery_v1_bakery_proto_rawDescGZIP(), []int{1}
}

func (x *Toppings) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Toppings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response to send, when Pizza is baking
type BakePizzaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the pizza
	// represents the id of the pizza in the whole system
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// current Status of the pizza (PREPARING, BAKING, READY)
	Status BakePizzaResponse_BakeStatus `protobuf:"varint,3,opt,name=status,proto3,enum=bakery.v1.BakePizzaResponse_BakeStatus" json:"status,omitempty"`
}

func (x *BakePizzaResponse) Reset() {
	*x = BakePizzaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bakery_v1_bakery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BakePizzaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BakePizzaResponse) ProtoMessage() {}

func (x *BakePizzaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bakery_v1_bakery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BakePizzaResponse.ProtoReflect.Descriptor instead.
func (*BakePizzaResponse) Descriptor() ([]byte, []int) {
	return file_bakery_v1_bakery_proto_rawDescGZIP(), []int{2}
}

func (x *BakePizzaResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BakePizzaResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BakePizzaResponse) GetStatus() BakePizzaResponse_BakeStatus {
	if x != nil {
		return x.Status
	}
	return BakePizzaResponse_BAKE_STATUS_UNSPECIFIED
}

var File_bakery_v1_bakery_proto protoreflect.FileDescriptor

var file_bakery_v1_bakery_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6b, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x22, 0x57, 0x0a, 0x10, 0x42, 0x61, 0x6b, 0x65, 0x50, 0x69, 0x7a, 0x7a, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x6b, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08,
	0x74, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x08,
	0x54, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xed, 0x01, 0x0a,
	0x11, 0x42, 0x61, 0x6b, 0x65, 0x50, 0x69, 0x7a, 0x7a, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x6b, 0x65, 0x50, 0x69, 0x7a, 0x7a, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x73, 0x0a, 0x0a, 0x42, 0x61, 0x6b, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x42, 0x41, 0x4b, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x41, 0x4b, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x42, 0x41, 0x4b, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x41, 0x4b,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x41, 0x4b, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x03, 0x32, 0x5b, 0x0a, 0x0d,
	0x42, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x09, 0x42, 0x61, 0x6b, 0x65, 0x50, 0x69, 0x7a, 0x7a, 0x61, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x6b,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6b, 0x65, 0x50, 0x69, 0x7a, 0x7a, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6b, 0x65, 0x50, 0x69, 0x7a, 0x7a, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x77, 0x6e, 0x33, 0x64, 0x2f, 0x70,
	0x69, 0x7a, 0x7a, 0x61, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6b,
	0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bakery_v1_bakery_proto_rawDescOnce sync.Once
	file_bakery_v1_bakery_proto_rawDescData = file_bakery_v1_bakery_proto_rawDesc
)

func file_bakery_v1_bakery_proto_rawDescGZIP() []byte {
	file_bakery_v1_bakery_proto_rawDescOnce.Do(func() {
		file_bakery_v1_bakery_proto_rawDescData = protoimpl.X.CompressGZIP(file_bakery_v1_bakery_proto_rawDescData)
	})
	return file_bakery_v1_bakery_proto_rawDescData
}

var file_bakery_v1_bakery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bakery_v1_bakery_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bakery_v1_bakery_proto_goTypes = []interface{}{
	(BakePizzaResponse_BakeStatus)(0), // 0: bakery.v1.BakePizzaResponse.BakeStatus
	(*BakePizzaRequest)(nil),          // 1: bakery.v1.BakePizzaRequest
	(*Toppings)(nil),                  // 2: bakery.v1.Toppings
	(*BakePizzaResponse)(nil),         // 3: bakery.v1.BakePizzaResponse
}
var file_bakery_v1_bakery_proto_depIdxs = []int32{
	2, // 0: bakery.v1.BakePizzaRequest.toppings:type_name -> bakery.v1.Toppings
	0, // 1: bakery.v1.BakePizzaResponse.status:type_name -> bakery.v1.BakePizzaResponse.BakeStatus
	1, // 2: bakery.v1.BakeryService.BakePizza:input_type -> bakery.v1.BakePizzaRequest
	3, // 3: bakery.v1.BakeryService.BakePizza:output_type -> bakery.v1.BakePizzaResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bakery_v1_bakery_proto_init() }
func file_bakery_v1_bakery_proto_init() {
	if File_bakery_v1_bakery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bakery_v1_bakery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BakePizzaRequest); i {
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
		file_bakery_v1_bakery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Toppings); i {
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
		file_bakery_v1_bakery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BakePizzaResponse); i {
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
			RawDescriptor: file_bakery_v1_bakery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bakery_v1_bakery_proto_goTypes,
		DependencyIndexes: file_bakery_v1_bakery_proto_depIdxs,
		EnumInfos:         file_bakery_v1_bakery_proto_enumTypes,
		MessageInfos:      file_bakery_v1_bakery_proto_msgTypes,
	}.Build()
	File_bakery_v1_bakery_proto = out.File
	file_bakery_v1_bakery_proto_rawDesc = nil
	file_bakery_v1_bakery_proto_goTypes = nil
	file_bakery_v1_bakery_proto_depIdxs = nil
}
