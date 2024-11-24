// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: accounts.proto

package generated

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

type AccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumAddress string `protobuf:"bytes,1,opt,name=ethereum_address,json=ethereumAddress,proto3" json:"ethereum_address,omitempty"`
	CryptoSignature string `protobuf:"bytes,2,opt,name=crypto_signature,json=cryptoSignature,proto3" json:"crypto_signature,omitempty"`
}

func (x *AccountRequest) Reset() {
	*x = AccountRequest{}
	mi := &file_accounts_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRequest) ProtoMessage() {}

func (x *AccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRequest.ProtoReflect.Descriptor instead.
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *AccountRequest) GetEthereumAddress() string {
	if x != nil {
		return x.EthereumAddress
	}
	return ""
}

func (x *AccountRequest) GetCryptoSignature() string {
	if x != nil {
		return x.CryptoSignature
	}
	return ""
}

type AccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumAddresses []string `protobuf:"bytes,1,rep,name=ethereum_addresses,json=ethereumAddresses,proto3" json:"ethereum_addresses,omitempty"`
	Erc20TokenAddress string   `protobuf:"bytes,2,opt,name=erc20_token_address,json=erc20TokenAddress,proto3" json:"erc20_token_address,omitempty"`
}

func (x *AccountsRequest) Reset() {
	*x = AccountsRequest{}
	mi := &file_accounts_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsRequest) ProtoMessage() {}

func (x *AccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsRequest.ProtoReflect.Descriptor instead.
func (*AccountsRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *AccountsRequest) GetEthereumAddresses() []string {
	if x != nil {
		return x.EthereumAddresses
	}
	return nil
}

func (x *AccountsRequest) GetErc20TokenAddress() string {
	if x != nil {
		return x.Erc20TokenAddress
	}
	return ""
}

type AccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GastokenBalance string `protobuf:"bytes,1,opt,name=gastoken_balance,json=gastokenBalance,proto3" json:"gastoken_balance,omitempty"`
	WalletNonce     string `protobuf:"bytes,2,opt,name=wallet_nonce,json=walletNonce,proto3" json:"wallet_nonce,omitempty"`
}

func (x *AccountResponse) Reset() {
	*x = AccountResponse{}
	mi := &file_accounts_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResponse) ProtoMessage() {}

func (x *AccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResponse.ProtoReflect.Descriptor instead.
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *AccountResponse) GetGastokenBalance() string {
	if x != nil {
		return x.GastokenBalance
	}
	return ""
}

func (x *AccountResponse) GetWalletNonce() string {
	if x != nil {
		return x.WalletNonce
	}
	return ""
}

type AccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthereumAddresses string `protobuf:"bytes,1,opt,name=ethereum_addresses,json=ethereumAddresses,proto3" json:"ethereum_addresses,omitempty"`
	Erc20Balance      string `protobuf:"bytes,2,opt,name=erc20_balance,json=erc20Balance,proto3" json:"erc20_balance,omitempty"`
}

func (x *AccountsResponse) Reset() {
	*x = AccountsResponse{}
	mi := &file_accounts_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountsResponse) ProtoMessage() {}

func (x *AccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountsResponse.ProtoReflect.Descriptor instead.
func (*AccountsResponse) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *AccountsResponse) GetEthereumAddresses() string {
	if x != nil {
		return x.EthereumAddresses
	}
	return ""
}

func (x *AccountsResponse) GetErc20Balance() string {
	if x != nil {
		return x.Erc20Balance
	}
	return ""
}

var File_accounts_proto protoreflect.FileDescriptor

var file_accounts_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x66, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x70, 0x0a, 0x0f,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x12, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2e,
	0x0a, 0x13, 0x65, 0x72, 0x63, 0x32, 0x30, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x72, 0x63,
	0x32, 0x30, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5f,
	0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x67, 0x61, 0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x61, 0x73,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x22,
	0x66, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x63, 0x32, 0x30, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x63, 0x32, 0x30,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x85, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounts_proto_rawDescOnce sync.Once
	file_accounts_proto_rawDescData = file_accounts_proto_rawDesc
)

func file_accounts_proto_rawDescGZIP() []byte {
	file_accounts_proto_rawDescOnce.Do(func() {
		file_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounts_proto_rawDescData)
	})
	return file_accounts_proto_rawDescData
}

var file_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_accounts_proto_goTypes = []any{
	(*AccountRequest)(nil),   // 0: pb.AccountRequest
	(*AccountsRequest)(nil),  // 1: pb.AccountsRequest
	(*AccountResponse)(nil),  // 2: pb.AccountResponse
	(*AccountsResponse)(nil), // 3: pb.AccountsResponse
}
var file_accounts_proto_depIdxs = []int32{
	0, // 0: pb.AccountService.GetAccount:input_type -> pb.AccountRequest
	1, // 1: pb.AccountService.GetAccounts:input_type -> pb.AccountsRequest
	2, // 2: pb.AccountService.GetAccount:output_type -> pb.AccountResponse
	3, // 3: pb.AccountService.GetAccounts:output_type -> pb.AccountsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accounts_proto_init() }
func file_accounts_proto_init() {
	if File_accounts_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accounts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounts_proto_goTypes,
		DependencyIndexes: file_accounts_proto_depIdxs,
		MessageInfos:      file_accounts_proto_msgTypes,
	}.Build()
	File_accounts_proto = out.File
	file_accounts_proto_rawDesc = nil
	file_accounts_proto_goTypes = nil
	file_accounts_proto_depIdxs = nil
}
