// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: liquidity.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LiquidityPoolType struct {
	PoolTypeIndex         uint32 `protobuf:"varint,1,opt,name=poolTypeIndex,proto3" json:"poolTypeIndex,omitempty" yaml:"pool_type_index"`
	NumOfReserveTokens    uint32 `protobuf:"varint,2,opt,name=NumOfReserveTokens,proto3" json:"NumOfReserveTokens,omitempty" yaml:"num_of_reserve_tokens"`
	SwapPriceFunctionName string `protobuf:"bytes,3,opt,name=SwapPriceFunctionName,proto3" json:"SwapPriceFunctionName,omitempty" yaml:"swap_price_function_name"`
	Description           string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty" yaml:"description"`
}

func (m *LiquidityPoolType) Reset()         { *m = LiquidityPoolType{} }
func (m *LiquidityPoolType) String() string { return proto.CompactTextString(m) }
func (*LiquidityPoolType) ProtoMessage()    {}
func (*LiquidityPoolType) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b46dab34d3c00e, []int{0}
}
func (m *LiquidityPoolType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidityPoolType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidityPoolType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidityPoolType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityPoolType.Merge(m, src)
}
func (m *LiquidityPoolType) XXX_Size() int {
	return m.Size()
}
func (m *LiquidityPoolType) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityPoolType.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityPoolType proto.InternalMessageInfo

type Params struct {
	LiquidityPoolTypes      []LiquidityPoolType                    `protobuf:"bytes,1,rep,name=LiquidityPoolTypes,proto3" json:"LiquidityPoolTypes" yaml:"liquidity_pool_types"`
	MinInitDepositToPool    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=MinInitDepositToPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"MinInitDepositToPool" yaml:"min_init_deposit_to_pool"`
	InitPoolTokenMintAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=InitPoolTokenMintAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"InitPoolTokenMintAmount" yaml:"init_pool_token_mint_amount"`
	SwapFeeRate             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=SwapFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"SwapFeeRate" yaml:"swap_fee_rate"`
	LiquidityPoolFeeRate    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=LiquidityPoolFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"LiquidityPoolFeeRate" yaml:"liquidity_pool_fee_rate"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b46dab34d3c00e, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

type LiquidityPool struct {
	PoolID             uint64                                        `protobuf:"varint,1,opt,name=PoolID,proto3" json:"PoolID,omitempty" yaml:"pool_id"`
	PoolTypeIndex      uint32                                        `protobuf:"varint,2,opt,name=poolTypeIndex,proto3" json:"poolTypeIndex,omitempty" yaml:"pool_type_index"`
	ReserveTokenDenoms []string                                      `protobuf:"bytes,3,rep,name=ReserveTokenDenoms,proto3" json:"ReserveTokenDenoms,omitempty" yaml:"reserve_token_denoms"`
	ReserveAccount     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,opt,name=ReserveAccount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"ReserveAccount,omitempty" yaml:"reserve_account"`
	PoolTokenDenom     string                                        `protobuf:"bytes,5,opt,name=PoolTokenDenom,proto3" json:"PoolTokenDenom,omitempty" yaml:"pool_token_denom"`
	SwapFeeRate        github_com_cosmos_cosmos_sdk_types.Dec        `protobuf:"bytes,6,opt,name=SwapFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"SwapFeeRate" yaml:"swap_fee_rate"`
	PoolFeeRate        github_com_cosmos_cosmos_sdk_types.Dec        `protobuf:"bytes,7,opt,name=PoolFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"PoolFeeRate" yaml:"pool_fee_rate"`
	BatchSize          uint32                                        `protobuf:"varint,8,opt,name=BatchSize,proto3" json:"BatchSize,omitempty" yaml:"batch_size"`
}

func (m *LiquidityPool) Reset()      { *m = LiquidityPool{} }
func (*LiquidityPool) ProtoMessage() {}
func (*LiquidityPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b46dab34d3c00e, []int{2}
}
func (m *LiquidityPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidityPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidityPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidityPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityPool.Merge(m, src)
}
func (m *LiquidityPool) XXX_Size() int {
	return m.Size()
}
func (m *LiquidityPool) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityPool.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityPool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LiquidityPoolType)(nil), "cosmos.liquidity.LiquidityPoolType")
	proto.RegisterType((*Params)(nil), "cosmos.liquidity.Params")
	proto.RegisterType((*LiquidityPool)(nil), "cosmos.liquidity.LiquidityPool")
}

func init() { proto.RegisterFile("liquidity.proto", fileDescriptor_65b46dab34d3c00e) }

var fileDescriptor_65b46dab34d3c00e = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x4f, 0xeb, 0x46,
	0x14, 0x8d, 0x09, 0xa4, 0x64, 0x52, 0x68, 0x19, 0x05, 0x88, 0xa0, 0xf2, 0x44, 0x83, 0x5a, 0x45,
	0x95, 0x48, 0xa4, 0xb2, 0xa9, 0x58, 0x35, 0x26, 0x42, 0x8a, 0x5a, 0x20, 0x35, 0xd9, 0xb4, 0x9b,
	0x91, 0xb1, 0x27, 0x64, 0x04, 0x9e, 0x71, 0xed, 0x49, 0x21, 0xac, 0xdb, 0x6d, 0xd5, 0x55, 0x55,
	0xa9, 0x1b, 0x7e, 0x48, 0x7f, 0x00, 0x4b, 0x96, 0x55, 0x17, 0x56, 0x0b, 0x9b, 0xb7, 0xf6, 0xf2,
	0xad, 0x9e, 0x3c, 0x0e, 0xb1, 0xf3, 0xf1, 0xa4, 0x17, 0x3d, 0xbd, 0x55, 0xac, 0xb9, 0xe7, 0x9e,
	0x73, 0x74, 0xef, 0xb9, 0x0a, 0xf8, 0xe4, 0x9a, 0xfd, 0x34, 0x60, 0x0e, 0x93, 0xc3, 0xba, 0xe7,
	0x0b, 0x29, 0xe0, 0xa7, 0xb6, 0x08, 0x5c, 0x11, 0xd4, 0xc7, 0xef, 0x3b, 0xe5, 0x4b, 0x71, 0x29,
	0x54, 0xb1, 0x11, 0x7f, 0x25, 0x38, 0xfc, 0xf7, 0x12, 0xd8, 0xf8, 0xee, 0x05, 0xd3, 0x11, 0xe2,
	0xba, 0x3b, 0xf4, 0x28, 0xfc, 0x06, 0xac, 0x79, 0xa3, 0xef, 0x36, 0x77, 0xe8, 0x6d, 0x45, 0xab,
	0x6a, 0xb5, 0x35, 0x63, 0x27, 0x0a, 0xd1, 0xd6, 0xd0, 0x72, 0xaf, 0x0f, 0x71, 0x5c, 0x26, 0x72,
	0xe8, 0x51, 0xc2, 0x62, 0x00, 0x36, 0x27, 0x1b, 0x60, 0x07, 0xc0, 0xd3, 0x81, 0x7b, 0xd6, 0x33,
	0x69, 0x40, 0xfd, 0x9f, 0x69, 0x57, 0x5c, 0x51, 0x1e, 0x54, 0x96, 0x14, 0x4d, 0x35, 0x0a, 0xd1,
	0x67, 0x09, 0x0d, 0x1f, 0xb8, 0x44, 0xf4, 0x88, 0x9f, 0xa0, 0x88, 0x54, 0x30, 0x6c, 0xce, 0xe9,
	0x85, 0x3f, 0x80, 0xcd, 0xf3, 0x1b, 0xcb, 0xeb, 0xf8, 0xcc, 0xa6, 0xc7, 0x03, 0x6e, 0x4b, 0x26,
	0xf8, 0xa9, 0xe5, 0xd2, 0x4a, 0xbe, 0xaa, 0xd5, 0x8a, 0xc6, 0x5e, 0x14, 0x22, 0x94, 0x90, 0x06,
	0x37, 0x96, 0x47, 0xbc, 0x18, 0x47, 0x7a, 0x23, 0x20, 0xe1, 0x96, 0x4b, 0xb1, 0x39, 0x9f, 0x01,
	0x7e, 0x0d, 0x4a, 0x2d, 0x1a, 0xd8, 0x3e, 0xf3, 0xe2, 0xa7, 0xca, 0xb2, 0x22, 0xdc, 0x8a, 0x42,
	0x04, 0x13, 0x42, 0x27, 0x2d, 0x62, 0x33, 0x0b, 0xc5, 0x7f, 0xac, 0x80, 0x42, 0xc7, 0xf2, 0x2d,
	0x37, 0x80, 0xb7, 0x00, 0xce, 0x0c, 0x32, 0xa8, 0x68, 0xd5, 0x7c, 0xad, 0xf4, 0xd5, 0x5e, 0x7d,
	0x7a, 0x1d, 0xf5, 0x19, 0xac, 0xb1, 0xf7, 0x10, 0xa2, 0x5c, 0x14, 0xa2, 0xdd, 0x44, 0x74, 0x0c,
	0x25, 0xe3, 0x59, 0xc7, 0x93, 0x99, 0xd5, 0x80, 0xbf, 0x6a, 0xa0, 0x7c, 0xc2, 0x78, 0x9b, 0x33,
	0xd9, 0xa2, 0x9e, 0x08, 0x98, 0xec, 0x8a, 0xb8, 0xaa, 0xc6, 0x5d, 0x34, 0xbe, 0x8f, 0x79, 0xff,
	0x0d, 0xd1, 0x17, 0x97, 0x4c, 0xf6, 0x07, 0x17, 0x75, 0x5b, 0xb8, 0x8d, 0xc4, 0xce, 0xe8, 0x67,
	0x3f, 0x70, 0xae, 0x1a, 0x8a, 0xbf, 0xde, 0xe6, 0x32, 0x9d, 0xa3, 0xcb, 0x38, 0x61, 0x9c, 0x49,
	0xe2, 0x24, 0xac, 0x44, 0x0a, 0xe5, 0x05, 0x9b, 0x73, 0xe5, 0xe0, 0x6f, 0x1a, 0xd8, 0x8e, 0x5f,
	0x95, 0xb3, 0x78, 0x69, 0x27, 0x8c, 0xcb, 0xa6, 0x2b, 0x06, 0x5c, 0x8e, 0x96, 0xd4, 0x5d, 0xd8,
	0x0a, 0x4e, 0xac, 0x28, 0x1b, 0xc9, 0x1c, 0x62, 0x62, 0xe2, 0x32, 0x2e, 0x89, 0xa5, 0xa8, 0xb1,
	0xf9, 0x36, 0x51, 0xd8, 0x07, 0xa5, 0x78, 0xe1, 0xc7, 0x94, 0x9a, 0x96, 0xa4, 0x6a, 0xaf, 0x1f,
	0x1b, 0xc7, 0x0b, 0x78, 0x68, 0x51, 0x3b, 0x0a, 0x51, 0x39, 0x13, 0xab, 0x1e, 0xa5, 0xc4, 0xb7,
	0x24, 0xc5, 0x66, 0x96, 0x1a, 0xfe, 0xa2, 0x81, 0xf2, 0xc4, 0x66, 0x5e, 0x34, 0x57, 0x94, 0x66,
	0x67, 0x61, 0x4d, 0x7d, 0x6e, 0x08, 0x52, 0xf5, 0xb9, 0x6a, 0x87, 0xab, 0x7f, 0xde, 0xa3, 0xdc,
	0xab, 0x7b, 0xa4, 0xe1, 0xbf, 0x56, 0xc0, 0xda, 0x04, 0x04, 0x7e, 0x09, 0x0a, 0xf1, 0x6f, 0xbb,
	0xa5, 0x8e, 0x79, 0xd9, 0x80, 0x51, 0x88, 0xd6, 0x33, 0xc7, 0xcc, 0x1c, 0x6c, 0x8e, 0x10, 0xb3,
	0xf7, 0xbf, 0xb4, 0xe8, 0xfd, 0x9f, 0x01, 0x98, 0x3d, 0xdf, 0x16, 0xe5, 0xc2, 0x0d, 0x2a, 0xf9,
	0x6a, 0xbe, 0x56, 0x34, 0x50, 0x1a, 0xf2, 0x89, 0xc3, 0x27, 0x8e, 0x42, 0x61, 0x73, 0x4e, 0x2b,
	0x0c, 0xc0, 0xfa, 0xe8, 0xb5, 0x69, 0xdb, 0x2a, 0x52, 0xc9, 0x3a, 0xbf, 0x4d, 0x3d, 0xbd, 0x90,
	0x59, 0x09, 0x00, 0xbf, 0x0e, 0xd1, 0xfe, 0x3b, 0x0c, 0xbc, 0x69, 0xdb, 0x4d, 0xc7, 0xf1, 0x69,
	0x10, 0x98, 0x53, 0x12, 0xf0, 0x08, 0xac, 0x8f, 0x73, 0xa5, 0x7c, 0xa8, 0x7d, 0x16, 0x8d, 0xdd,
	0x28, 0x44, 0xdb, 0xd9, 0x41, 0xa4, 0xf6, 0xb1, 0x39, 0xd5, 0x32, 0x9d, 0xc2, 0xc2, 0x87, 0x4b,
	0x61, 0x1f, 0x94, 0xb2, 0xd9, 0xfb, 0xe8, 0xfd, 0x94, 0xa6, 0x12, 0x97, 0xa5, 0x86, 0x07, 0xa0,
	0x68, 0x58, 0xd2, 0xee, 0x9f, 0xb3, 0x3b, 0x5a, 0x59, 0x55, 0xe1, 0xd8, 0x8c, 0x42, 0xb4, 0x91,
	0x74, 0x5e, 0xc4, 0x25, 0x12, 0xb0, 0x3b, 0x8a, 0xcd, 0x14, 0x97, 0xa6, 0xd3, 0x38, 0x7a, 0xf8,
	0x5f, 0xcf, 0x3d, 0x3c, 0xe9, 0xda, 0xe3, 0x93, 0xae, 0xfd, 0xf7, 0xa4, 0x6b, 0xbf, 0x3f, 0xeb,
	0xb9, 0xc7, 0x67, 0x3d, 0xf7, 0xcf, 0xb3, 0x9e, 0xfb, 0xf1, 0xf3, 0x8c, 0x53, 0x49, 0xb9, 0x43,
	0xfd, 0xf8, 0xdc, 0x1b, 0xe3, 0x5b, 0x48, 0xcc, 0x5e, 0x14, 0xd4, 0x3f, 0xd8, 0xc1, 0x9b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xca, 0x82, 0x75, 0xdc, 0xfc, 0x06, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.LiquidityPoolTypes) != len(that1.LiquidityPoolTypes) {
		return false
	}
	for i := range this.LiquidityPoolTypes {
		if !this.LiquidityPoolTypes[i].Equal(&that1.LiquidityPoolTypes[i]) {
			return false
		}
	}
	if !this.MinInitDepositToPool.Equal(that1.MinInitDepositToPool) {
		return false
	}
	if !this.InitPoolTokenMintAmount.Equal(that1.InitPoolTokenMintAmount) {
		return false
	}
	if !this.SwapFeeRate.Equal(that1.SwapFeeRate) {
		return false
	}
	if !this.LiquidityPoolFeeRate.Equal(that1.LiquidityPoolFeeRate) {
		return false
	}
	return true
}
func (this *LiquidityPool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LiquidityPool)
	if !ok {
		that2, ok := that.(LiquidityPool)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PoolID != that1.PoolID {
		return false
	}
	if this.PoolTypeIndex != that1.PoolTypeIndex {
		return false
	}
	if len(this.ReserveTokenDenoms) != len(that1.ReserveTokenDenoms) {
		return false
	}
	for i := range this.ReserveTokenDenoms {
		if this.ReserveTokenDenoms[i] != that1.ReserveTokenDenoms[i] {
			return false
		}
	}
	if !bytes.Equal(this.ReserveAccount, that1.ReserveAccount) {
		return false
	}
	if this.PoolTokenDenom != that1.PoolTokenDenom {
		return false
	}
	if !this.SwapFeeRate.Equal(that1.SwapFeeRate) {
		return false
	}
	if !this.PoolFeeRate.Equal(that1.PoolFeeRate) {
		return false
	}
	if this.BatchSize != that1.BatchSize {
		return false
	}
	return true
}
func (m *LiquidityPoolType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidityPoolType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidityPoolType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SwapPriceFunctionName) > 0 {
		i -= len(m.SwapPriceFunctionName)
		copy(dAtA[i:], m.SwapPriceFunctionName)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.SwapPriceFunctionName)))
		i--
		dAtA[i] = 0x1a
	}
	if m.NumOfReserveTokens != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.NumOfReserveTokens))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolTypeIndex != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.PoolTypeIndex))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LiquidityPoolFeeRate.Size()
		i -= size
		if _, err := m.LiquidityPoolFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SwapFeeRate.Size()
		i -= size
		if _, err := m.SwapFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InitPoolTokenMintAmount.Size()
		i -= size
		if _, err := m.InitPoolTokenMintAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinInitDepositToPool.Size()
		i -= size
		if _, err := m.MinInitDepositToPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.LiquidityPoolTypes) > 0 {
		for iNdEx := len(m.LiquidityPoolTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LiquidityPoolTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidity(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LiquidityPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidityPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidityPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BatchSize != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.BatchSize))
		i--
		dAtA[i] = 0x40
	}
	{
		size := m.PoolFeeRate.Size()
		i -= size
		if _, err := m.PoolFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.SwapFeeRate.Size()
		i -= size
		if _, err := m.SwapFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidity(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.PoolTokenDenom) > 0 {
		i -= len(m.PoolTokenDenom)
		copy(dAtA[i:], m.PoolTokenDenom)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.PoolTokenDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ReserveAccount) > 0 {
		i -= len(m.ReserveAccount)
		copy(dAtA[i:], m.ReserveAccount)
		i = encodeVarintLiquidity(dAtA, i, uint64(len(m.ReserveAccount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ReserveTokenDenoms) > 0 {
		for iNdEx := len(m.ReserveTokenDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ReserveTokenDenoms[iNdEx])
			copy(dAtA[i:], m.ReserveTokenDenoms[iNdEx])
			i = encodeVarintLiquidity(dAtA, i, uint64(len(m.ReserveTokenDenoms[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolTypeIndex != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.PoolTypeIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolID != 0 {
		i = encodeVarintLiquidity(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidity(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LiquidityPoolType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolTypeIndex != 0 {
		n += 1 + sovLiquidity(uint64(m.PoolTypeIndex))
	}
	if m.NumOfReserveTokens != 0 {
		n += 1 + sovLiquidity(uint64(m.NumOfReserveTokens))
	}
	l = len(m.SwapPriceFunctionName)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LiquidityPoolTypes) > 0 {
		for _, e := range m.LiquidityPoolTypes {
			l = e.Size()
			n += 1 + l + sovLiquidity(uint64(l))
		}
	}
	l = m.MinInitDepositToPool.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	l = m.InitPoolTokenMintAmount.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	l = m.SwapFeeRate.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	l = m.LiquidityPoolFeeRate.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	return n
}

func (m *LiquidityPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovLiquidity(uint64(m.PoolID))
	}
	if m.PoolTypeIndex != 0 {
		n += 1 + sovLiquidity(uint64(m.PoolTypeIndex))
	}
	if len(m.ReserveTokenDenoms) > 0 {
		for _, s := range m.ReserveTokenDenoms {
			l = len(s)
			n += 1 + l + sovLiquidity(uint64(l))
		}
	}
	l = len(m.ReserveAccount)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	l = len(m.PoolTokenDenom)
	if l > 0 {
		n += 1 + l + sovLiquidity(uint64(l))
	}
	l = m.SwapFeeRate.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	l = m.PoolFeeRate.Size()
	n += 1 + l + sovLiquidity(uint64(l))
	if m.BatchSize != 0 {
		n += 1 + sovLiquidity(uint64(m.BatchSize))
	}
	return n
}

func sovLiquidity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidity(x uint64) (n int) {
	return sovLiquidity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LiquidityPoolType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LiquidityPoolType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidityPoolType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolTypeIndex", wireType)
			}
			m.PoolTypeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolTypeIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumOfReserveTokens", wireType)
			}
			m.NumOfReserveTokens = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumOfReserveTokens |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapPriceFunctionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SwapPriceFunctionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLiquidity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLiquidity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityPoolTypes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidityPoolTypes = append(m.LiquidityPoolTypes, LiquidityPoolType{})
			if err := m.LiquidityPoolTypes[len(m.LiquidityPoolTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitDepositToPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinInitDepositToPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitPoolTokenMintAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitPoolTokenMintAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityPoolFeeRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityPoolFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLiquidity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLiquidity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LiquidityPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidity
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LiquidityPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidityPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolTypeIndex", wireType)
			}
			m.PoolTypeIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolTypeIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveTokenDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReserveTokenDenoms = append(m.ReserveTokenDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveAccount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReserveAccount = append(m.ReserveAccount[:0], dAtA[iNdEx:postIndex]...)
			if m.ReserveAccount == nil {
				m.ReserveAccount = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolTokenDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolTokenDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolFeeRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLiquidity
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchSize", wireType)
			}
			m.BatchSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLiquidity
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLiquidity
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLiquidity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidity
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLiquidity
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLiquidity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidity = fmt.Errorf("proto: unexpected end of group")
)
