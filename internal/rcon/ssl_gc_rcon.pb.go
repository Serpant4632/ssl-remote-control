// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: ssl_gc_rcon.proto

package rcon

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

type ControllerReply_StatusCode int32

const (
	ControllerReply_UNKNOWN_STATUS_CODE ControllerReply_StatusCode = 0
	ControllerReply_OK                  ControllerReply_StatusCode = 1
	ControllerReply_REJECTED            ControllerReply_StatusCode = 2
)

// Enum value maps for ControllerReply_StatusCode.
var (
	ControllerReply_StatusCode_name = map[int32]string{
		0: "UNKNOWN_STATUS_CODE",
		1: "OK",
		2: "REJECTED",
	}
	ControllerReply_StatusCode_value = map[string]int32{
		"UNKNOWN_STATUS_CODE": 0,
		"OK":                  1,
		"REJECTED":            2,
	}
)

func (x ControllerReply_StatusCode) Enum() *ControllerReply_StatusCode {
	p := new(ControllerReply_StatusCode)
	*p = x
	return p
}

func (x ControllerReply_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControllerReply_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ssl_gc_rcon_proto_enumTypes[0].Descriptor()
}

func (ControllerReply_StatusCode) Type() protoreflect.EnumType {
	return &file_ssl_gc_rcon_proto_enumTypes[0]
}

func (x ControllerReply_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ControllerReply_StatusCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ControllerReply_StatusCode(num)
	return nil
}

// Deprecated: Use ControllerReply_StatusCode.Descriptor instead.
func (ControllerReply_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_proto_rawDescGZIP(), []int{0, 0}
}

type ControllerReply_Verification int32

const (
	ControllerReply_UNKNOWN_VERIFICATION ControllerReply_Verification = 0
	ControllerReply_VERIFIED             ControllerReply_Verification = 1
	ControllerReply_UNVERIFIED           ControllerReply_Verification = 2
)

// Enum value maps for ControllerReply_Verification.
var (
	ControllerReply_Verification_name = map[int32]string{
		0: "UNKNOWN_VERIFICATION",
		1: "VERIFIED",
		2: "UNVERIFIED",
	}
	ControllerReply_Verification_value = map[string]int32{
		"UNKNOWN_VERIFICATION": 0,
		"VERIFIED":             1,
		"UNVERIFIED":           2,
	}
)

func (x ControllerReply_Verification) Enum() *ControllerReply_Verification {
	p := new(ControllerReply_Verification)
	*p = x
	return p
}

func (x ControllerReply_Verification) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControllerReply_Verification) Descriptor() protoreflect.EnumDescriptor {
	return file_ssl_gc_rcon_proto_enumTypes[1].Descriptor()
}

func (ControllerReply_Verification) Type() protoreflect.EnumType {
	return &file_ssl_gc_rcon_proto_enumTypes[1]
}

func (x ControllerReply_Verification) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ControllerReply_Verification) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ControllerReply_Verification(num)
	return nil
}

// Deprecated: Use ControllerReply_Verification.Descriptor instead.
func (ControllerReply_Verification) EnumDescriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_proto_rawDescGZIP(), []int{0, 1}
}

// a reply that is sent by the controller for each request from teams or autoRefs
type ControllerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// status_code is an optional code that indicates the result of the last request
	StatusCode *ControllerReply_StatusCode `protobuf:"varint,1,opt,name=status_code,json=statusCode,enum=ControllerReply_StatusCode" json:"status_code,omitempty"`
	// reason is an optional explanation of the status code
	Reason *string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	// next_token must be send with the next request, if secure communication is used
	// the token is used to avoid replay attacks
	// the token is always present in the very first message before the registration starts
	// the token is not present, if secure communication is not used
	NextToken *string `protobuf:"bytes,3,opt,name=next_token,json=nextToken" json:"next_token,omitempty"`
	// verification indicates if the last request could be verified (secure communication)
	Verification *ControllerReply_Verification `protobuf:"varint,4,opt,name=verification,enum=ControllerReply_Verification" json:"verification,omitempty"`
}

func (x *ControllerReply) Reset() {
	*x = ControllerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerReply) ProtoMessage() {}

func (x *ControllerReply) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerReply.ProtoReflect.Descriptor instead.
func (*ControllerReply) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_proto_rawDescGZIP(), []int{0}
}

func (x *ControllerReply) GetStatusCode() ControllerReply_StatusCode {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return ControllerReply_UNKNOWN_STATUS_CODE
}

func (x *ControllerReply) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *ControllerReply) GetNextToken() string {
	if x != nil && x.NextToken != nil {
		return *x.NextToken
	}
	return ""
}

func (x *ControllerReply) GetVerification() ControllerReply_Verification {
	if x != nil && x.Verification != nil {
		return *x.Verification
	}
	return ControllerReply_UNKNOWN_VERIFICATION
}

// Signature can be added to a request to let it be verfied by the controller
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the token that was received with the last controller reply
	Token *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`
	// the PKCS1v15 signature of this message
	Pkcs1V15 []byte `protobuf:"bytes,2,req,name=pkcs1v15" json:"pkcs1v15,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *Signature) GetPkcs1V15() []byte {
	if x != nil {
		return x.Pkcs1V15
	}
	return nil
}

var File_ssl_gc_rcon_proto protoreflect.FileDescriptor

var file_ssl_gc_rcon_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x41, 0x0a, 0x0c,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3b, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a,
	0x13, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x22, 0x46, 0x0a, 0x0c,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x14,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x02, 0x22, 0x3d, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6b, 0x63, 0x73, 0x31,
	0x76, 0x31, 0x35, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x6b, 0x63, 0x73, 0x31,
	0x76, 0x31, 0x35, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73,
	0x6c, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x72, 0x63, 0x6f, 0x6e,
}

var (
	file_ssl_gc_rcon_proto_rawDescOnce sync.Once
	file_ssl_gc_rcon_proto_rawDescData = file_ssl_gc_rcon_proto_rawDesc
)

func file_ssl_gc_rcon_proto_rawDescGZIP() []byte {
	file_ssl_gc_rcon_proto_rawDescOnce.Do(func() {
		file_ssl_gc_rcon_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_gc_rcon_proto_rawDescData)
	})
	return file_ssl_gc_rcon_proto_rawDescData
}

var file_ssl_gc_rcon_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ssl_gc_rcon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ssl_gc_rcon_proto_goTypes = []interface{}{
	(ControllerReply_StatusCode)(0),   // 0: ControllerReply.StatusCode
	(ControllerReply_Verification)(0), // 1: ControllerReply.Verification
	(*ControllerReply)(nil),           // 2: ControllerReply
	(*Signature)(nil),                 // 3: Signature
}
var file_ssl_gc_rcon_proto_depIdxs = []int32{
	0, // 0: ControllerReply.status_code:type_name -> ControllerReply.StatusCode
	1, // 1: ControllerReply.verification:type_name -> ControllerReply.Verification
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ssl_gc_rcon_proto_init() }
func file_ssl_gc_rcon_proto_init() {
	if File_ssl_gc_rcon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssl_gc_rcon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerReply); i {
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
		file_ssl_gc_rcon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
			RawDescriptor: file_ssl_gc_rcon_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_gc_rcon_proto_goTypes,
		DependencyIndexes: file_ssl_gc_rcon_proto_depIdxs,
		EnumInfos:         file_ssl_gc_rcon_proto_enumTypes,
		MessageInfos:      file_ssl_gc_rcon_proto_msgTypes,
	}.Build()
	File_ssl_gc_rcon_proto = out.File
	file_ssl_gc_rcon_proto_rawDesc = nil
	file_ssl_gc_rcon_proto_goTypes = nil
	file_ssl_gc_rcon_proto_depIdxs = nil
}
