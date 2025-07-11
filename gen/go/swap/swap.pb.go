// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: swap/swap.proto

package swaprpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetStatusRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TransactionSign string                 `protobuf:"bytes,1,opt,name=TransactionSign,proto3" json:"TransactionSign,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetStatusRequest) Reset() {
	*x = GetStatusRequest{}
	mi := &file_swap_swap_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusRequest) ProtoMessage() {}

func (x *GetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swap_swap_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusRequest.ProtoReflect.Descriptor instead.
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return file_swap_swap_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatusRequest) GetTransactionSign() string {
	if x != nil {
		return x.TransactionSign
	}
	return ""
}

type SwapRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	WalletPrivateKey string                 `protobuf:"bytes,1,opt,name=WalletPrivateKey,proto3" json:"WalletPrivateKey,omitempty"`
	PayMint          string                 `protobuf:"bytes,2,opt,name=PayMint,proto3" json:"PayMint,omitempty"`
	BuyMint          string                 `protobuf:"bytes,3,opt,name=BuyMint,proto3" json:"BuyMint,omitempty"`
	AmountDouble     float64                `protobuf:"fixed64,4,opt,name=AmountDouble,proto3" json:"AmountDouble,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *SwapRequest) Reset() {
	*x = SwapRequest{}
	mi := &file_swap_swap_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapRequest) ProtoMessage() {}

func (x *SwapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_swap_swap_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapRequest.ProtoReflect.Descriptor instead.
func (*SwapRequest) Descriptor() ([]byte, []int) {
	return file_swap_swap_proto_rawDescGZIP(), []int{1}
}

func (x *SwapRequest) GetWalletPrivateKey() string {
	if x != nil {
		return x.WalletPrivateKey
	}
	return ""
}

func (x *SwapRequest) GetPayMint() string {
	if x != nil {
		return x.PayMint
	}
	return ""
}

func (x *SwapRequest) GetBuyMint() string {
	if x != nil {
		return x.BuyMint
	}
	return ""
}

func (x *SwapRequest) GetAmountDouble() float64 {
	if x != nil {
		return x.AmountDouble
	}
	return 0
}

type SwapResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	TransactionSign string                 `protobuf:"bytes,1,opt,name=TransactionSign,proto3" json:"TransactionSign,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *SwapResponse) Reset() {
	*x = SwapResponse{}
	mi := &file_swap_swap_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SwapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapResponse) ProtoMessage() {}

func (x *SwapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_swap_swap_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapResponse.ProtoReflect.Descriptor instead.
func (*SwapResponse) Descriptor() ([]byte, []int) {
	return file_swap_swap_proto_rawDescGZIP(), []int{2}
}

func (x *SwapResponse) GetTransactionSign() string {
	if x != nil {
		return x.TransactionSign
	}
	return ""
}

var File_swap_swap_proto protoreflect.FileDescriptor

const file_swap_swap_proto_rawDesc = "" +
	"\n" +
	"\x0fswap/swap.proto\x12\abalance\x1a\x1bgoogle/protobuf/empty.proto\"<\n" +
	"\x10GetStatusRequest\x12(\n" +
	"\x0fTransactionSign\x18\x01 \x01(\tR\x0fTransactionSign\"\x91\x01\n" +
	"\vSwapRequest\x12*\n" +
	"\x10WalletPrivateKey\x18\x01 \x01(\tR\x10WalletPrivateKey\x12\x18\n" +
	"\aPayMint\x18\x02 \x01(\tR\aPayMint\x12\x18\n" +
	"\aBuyMint\x18\x03 \x01(\tR\aBuyMint\x12\"\n" +
	"\fAmountDouble\x18\x04 \x01(\x01R\fAmountDouble\"8\n" +
	"\fSwapResponse\x12(\n" +
	"\x0fTransactionSign\x18\x01 \x01(\tR\x0fTransactionSign2{\n" +
	"\x04Swap\x123\n" +
	"\x04Swap\x12\x14.balance.SwapRequest\x1a\x15.balance.SwapResponse\x12>\n" +
	"\tGetStatus\x12\x19.balance.GetStatusRequest\x1a\x16.google.protobuf.EmptyB*Z(github.com/autumnterror/tb-proto;swaprpcb\x06proto3"

var (
	file_swap_swap_proto_rawDescOnce sync.Once
	file_swap_swap_proto_rawDescData []byte
)

func file_swap_swap_proto_rawDescGZIP() []byte {
	file_swap_swap_proto_rawDescOnce.Do(func() {
		file_swap_swap_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_swap_swap_proto_rawDesc), len(file_swap_swap_proto_rawDesc)))
	})
	return file_swap_swap_proto_rawDescData
}

var file_swap_swap_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_swap_swap_proto_goTypes = []any{
	(*GetStatusRequest)(nil), // 0: balance.GetStatusRequest
	(*SwapRequest)(nil),      // 1: balance.SwapRequest
	(*SwapResponse)(nil),     // 2: balance.SwapResponse
	(*emptypb.Empty)(nil),    // 3: google.protobuf.Empty
}
var file_swap_swap_proto_depIdxs = []int32{
	1, // 0: balance.Swap.Swap:input_type -> balance.SwapRequest
	0, // 1: balance.Swap.GetStatus:input_type -> balance.GetStatusRequest
	2, // 2: balance.Swap.Swap:output_type -> balance.SwapResponse
	3, // 3: balance.Swap.GetStatus:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_swap_swap_proto_init() }
func file_swap_swap_proto_init() {
	if File_swap_swap_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_swap_swap_proto_rawDesc), len(file_swap_swap_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_swap_swap_proto_goTypes,
		DependencyIndexes: file_swap_swap_proto_depIdxs,
		MessageInfos:      file_swap_swap_proto_msgTypes,
	}.Build()
	File_swap_swap_proto = out.File
	file_swap_swap_proto_goTypes = nil
	file_swap_swap_proto_depIdxs = nil
}
