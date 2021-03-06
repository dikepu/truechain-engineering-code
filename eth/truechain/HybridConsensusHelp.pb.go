// Code generated by protoc-gen-go. DO NOT EDIT.
// source: HybridConsensusHelp.proto

package truechain

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmptyParam struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyParam) Reset()         { *m = EmptyParam{} }
func (m *EmptyParam) String() string { return proto.CompactTextString(m) }
func (*EmptyParam) ProtoMessage()    {}
func (*EmptyParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{0}
}
func (m *EmptyParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyParam.Unmarshal(m, b)
}
func (m *EmptyParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyParam.Marshal(b, m, deterministic)
}
func (dst *EmptyParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyParam.Merge(dst, src)
}
func (m *EmptyParam) XXX_Size() int {
	return xxx_messageInfo_EmptyParam.Size(m)
}
func (m *EmptyParam) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyParam.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyParam proto.InternalMessageInfo

type BftPrivateKey struct {
	Pkey                 string   `protobuf:"bytes,1,opt,name=pkey" json:"pkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BftPrivateKey) Reset()         { *m = BftPrivateKey{} }
func (m *BftPrivateKey) String() string { return proto.CompactTextString(m) }
func (*BftPrivateKey) ProtoMessage()    {}
func (*BftPrivateKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{1}
}
func (m *BftPrivateKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BftPrivateKey.Unmarshal(m, b)
}
func (m *BftPrivateKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BftPrivateKey.Marshal(b, m, deterministic)
}
func (dst *BftPrivateKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BftPrivateKey.Merge(dst, src)
}
func (m *BftPrivateKey) XXX_Size() int {
	return xxx_messageInfo_BftPrivateKey.Size(m)
}
func (m *BftPrivateKey) XXX_DiscardUnknown() {
	xxx_messageInfo_BftPrivateKey.DiscardUnknown(m)
}

var xxx_messageInfo_BftPrivateKey proto.InternalMessageInfo

func (m *BftPrivateKey) GetPkey() string {
	if m != nil {
		return m.Pkey
	}
	return ""
}

type SignCommittee struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	Sigs                 []string `protobuf:"bytes,2,rep,name=sigs" json:"sigs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignCommittee) Reset()         { *m = SignCommittee{} }
func (m *SignCommittee) String() string { return proto.CompactTextString(m) }
func (*SignCommittee) ProtoMessage()    {}
func (*SignCommittee) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{2}
}
func (m *SignCommittee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignCommittee.Unmarshal(m, b)
}
func (m *SignCommittee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignCommittee.Marshal(b, m, deterministic)
}
func (dst *SignCommittee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignCommittee.Merge(dst, src)
}
func (m *SignCommittee) XXX_Size() int {
	return xxx_messageInfo_SignCommittee.Size(m)
}
func (m *SignCommittee) XXX_DiscardUnknown() {
	xxx_messageInfo_SignCommittee.DiscardUnknown(m)
}

var xxx_messageInfo_SignCommittee proto.InternalMessageInfo

func (m *SignCommittee) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SignCommittee) GetSigs() []string {
	if m != nil {
		return m.Sigs
	}
	return nil
}

type CommonReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonReply) Reset()         { *m = CommonReply{} }
func (m *CommonReply) String() string { return proto.CompactTextString(m) }
func (*CommonReply) ProtoMessage()    {}
func (*CommonReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{3}
}
func (m *CommonReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonReply.Unmarshal(m, b)
}
func (m *CommonReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonReply.Marshal(b, m, deterministic)
}
func (dst *CommonReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonReply.Merge(dst, src)
}
func (m *CommonReply) XXX_Size() int {
	return xxx_messageInfo_CommonReply.Size(m)
}
func (m *CommonReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonReply.DiscardUnknown(m)
}

var xxx_messageInfo_CommonReply proto.InternalMessageInfo

func (m *CommonReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TruePbftNode struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Pubkey               string   `protobuf:"bytes,2,opt,name=pubkey" json:"pubkey,omitempty"`
	Privkey              string   `protobuf:"bytes,3,opt,name=privkey" json:"privkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TruePbftNode) Reset()         { *m = TruePbftNode{} }
func (m *TruePbftNode) String() string { return proto.CompactTextString(m) }
func (*TruePbftNode) ProtoMessage()    {}
func (*TruePbftNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{4}
}
func (m *TruePbftNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TruePbftNode.Unmarshal(m, b)
}
func (m *TruePbftNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TruePbftNode.Marshal(b, m, deterministic)
}
func (dst *TruePbftNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TruePbftNode.Merge(dst, src)
}
func (m *TruePbftNode) XXX_Size() int {
	return xxx_messageInfo_TruePbftNode.Size(m)
}
func (m *TruePbftNode) XXX_DiscardUnknown() {
	xxx_messageInfo_TruePbftNode.DiscardUnknown(m)
}

var xxx_messageInfo_TruePbftNode proto.InternalMessageInfo

func (m *TruePbftNode) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *TruePbftNode) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *TruePbftNode) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

type Nodes struct {
	Nodes                []*TruePbftNode `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}
func (*Nodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{5}
}
func (m *Nodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nodes.Unmarshal(m, b)
}
func (m *Nodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nodes.Marshal(b, m, deterministic)
}
func (dst *Nodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nodes.Merge(dst, src)
}
func (m *Nodes) XXX_Size() int {
	return xxx_messageInfo_Nodes.Size(m)
}
func (m *Nodes) XXX_DiscardUnknown() {
	xxx_messageInfo_Nodes.DiscardUnknown(m)
}

var xxx_messageInfo_Nodes proto.InternalMessageInfo

func (m *Nodes) GetNodes() []*TruePbftNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type TruePbftBlockHeader struct {
	Number               int64    `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	GasLimit             int64    `protobuf:"varint,2,opt,name=GasLimit" json:"GasLimit,omitempty"`
	GasUsed              int64    `protobuf:"varint,3,opt,name=GasUsed" json:"GasUsed,omitempty"`
	Time                 int64    `protobuf:"varint,4,opt,name=Time" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TruePbftBlockHeader) Reset()         { *m = TruePbftBlockHeader{} }
func (m *TruePbftBlockHeader) String() string { return proto.CompactTextString(m) }
func (*TruePbftBlockHeader) ProtoMessage()    {}
func (*TruePbftBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{6}
}
func (m *TruePbftBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TruePbftBlockHeader.Unmarshal(m, b)
}
func (m *TruePbftBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TruePbftBlockHeader.Marshal(b, m, deterministic)
}
func (dst *TruePbftBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TruePbftBlockHeader.Merge(dst, src)
}
func (m *TruePbftBlockHeader) XXX_Size() int {
	return xxx_messageInfo_TruePbftBlockHeader.Size(m)
}
func (m *TruePbftBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_TruePbftBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_TruePbftBlockHeader proto.InternalMessageInfo

func (m *TruePbftBlockHeader) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *TruePbftBlockHeader) GetGasLimit() int64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *TruePbftBlockHeader) GetGasUsed() int64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *TruePbftBlockHeader) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type TxData struct {
	AccountNonce         uint64   `protobuf:"varint,1,opt,name=AccountNonce" json:"AccountNonce,omitempty"`
	Price                int64    `protobuf:"varint,2,opt,name=Price" json:"Price,omitempty"`
	GasLimit             int64    `protobuf:"varint,3,opt,name=GasLimit" json:"GasLimit,omitempty"`
	Recipient            []byte   `protobuf:"bytes,4,opt,name=Recipient,proto3" json:"Recipient,omitempty"`
	Amount               int64    `protobuf:"varint,5,opt,name=Amount" json:"Amount,omitempty"`
	Payload              []byte   `protobuf:"bytes,6,opt,name=Payload,proto3" json:"Payload,omitempty"`
	V                    int64    `protobuf:"varint,7,opt,name=V" json:"V,omitempty"`
	R                    int64    `protobuf:"varint,8,opt,name=R" json:"R,omitempty"`
	S                    int64    `protobuf:"varint,9,opt,name=S" json:"S,omitempty"`
	Hash                 []byte   `protobuf:"bytes,10,opt,name=Hash,proto3" json:"Hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxData) Reset()         { *m = TxData{} }
func (m *TxData) String() string { return proto.CompactTextString(m) }
func (*TxData) ProtoMessage()    {}
func (*TxData) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{7}
}
func (m *TxData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxData.Unmarshal(m, b)
}
func (m *TxData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxData.Marshal(b, m, deterministic)
}
func (dst *TxData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxData.Merge(dst, src)
}
func (m *TxData) XXX_Size() int {
	return xxx_messageInfo_TxData.Size(m)
}
func (m *TxData) XXX_DiscardUnknown() {
	xxx_messageInfo_TxData.DiscardUnknown(m)
}

var xxx_messageInfo_TxData proto.InternalMessageInfo

func (m *TxData) GetAccountNonce() uint64 {
	if m != nil {
		return m.AccountNonce
	}
	return 0
}

func (m *TxData) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TxData) GetGasLimit() int64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *TxData) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *TxData) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TxData) GetV() int64 {
	if m != nil {
		return m.V
	}
	return 0
}

func (m *TxData) GetR() int64 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *TxData) GetS() int64 {
	if m != nil {
		return m.S
	}
	return 0
}

func (m *TxData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Transaction struct {
	Data                 *TxData  `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{8}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetData() *TxData {
	if m != nil {
		return m.Data
	}
	return nil
}

type Transactions struct {
	Txs                  []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Transactions) Reset()         { *m = Transactions{} }
func (m *Transactions) String() string { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()    {}
func (*Transactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{9}
}
func (m *Transactions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transactions.Unmarshal(m, b)
}
func (m *Transactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transactions.Marshal(b, m, deterministic)
}
func (dst *Transactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transactions.Merge(dst, src)
}
func (m *Transactions) XXX_Size() int {
	return xxx_messageInfo_Transactions.Size(m)
}
func (m *Transactions) XXX_DiscardUnknown() {
	xxx_messageInfo_Transactions.DiscardUnknown(m)
}

var xxx_messageInfo_Transactions proto.InternalMessageInfo

func (m *Transactions) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type TruePbftBlock struct {
	Header               *TruePbftBlockHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Txs                  *Transactions        `protobuf:"bytes,2,opt,name=txs" json:"txs,omitempty"`
	Sigs                 []string             `protobuf:"bytes,3,rep,name=sigs" json:"sigs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TruePbftBlock) Reset()         { *m = TruePbftBlock{} }
func (m *TruePbftBlock) String() string { return proto.CompactTextString(m) }
func (*TruePbftBlock) ProtoMessage()    {}
func (*TruePbftBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059, []int{10}
}
func (m *TruePbftBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TruePbftBlock.Unmarshal(m, b)
}
func (m *TruePbftBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TruePbftBlock.Marshal(b, m, deterministic)
}
func (dst *TruePbftBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TruePbftBlock.Merge(dst, src)
}
func (m *TruePbftBlock) XXX_Size() int {
	return xxx_messageInfo_TruePbftBlock.Size(m)
}
func (m *TruePbftBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_TruePbftBlock.DiscardUnknown(m)
}

var xxx_messageInfo_TruePbftBlock proto.InternalMessageInfo

func (m *TruePbftBlock) GetHeader() *TruePbftBlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TruePbftBlock) GetTxs() *Transactions {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *TruePbftBlock) GetSigs() []string {
	if m != nil {
		return m.Sigs
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyParam)(nil), "truechain.EmptyParam")
	proto.RegisterType((*BftPrivateKey)(nil), "truechain.BftPrivateKey")
	proto.RegisterType((*SignCommittee)(nil), "truechain.SignCommittee")
	proto.RegisterType((*CommonReply)(nil), "truechain.CommonReply")
	proto.RegisterType((*TruePbftNode)(nil), "truechain.TruePbftNode")
	proto.RegisterType((*Nodes)(nil), "truechain.Nodes")
	proto.RegisterType((*TruePbftBlockHeader)(nil), "truechain.TruePbftBlockHeader")
	proto.RegisterType((*TxData)(nil), "truechain.TxData")
	proto.RegisterType((*Transaction)(nil), "truechain.Transaction")
	proto.RegisterType((*Transactions)(nil), "truechain.Transactions")
	proto.RegisterType((*TruePbftBlock)(nil), "truechain.TruePbftBlock")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HybridConsensusHelpClient is the client API for HybridConsensusHelp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HybridConsensusHelpClient interface {
	PutBlock(ctx context.Context, in *TruePbftBlock, opts ...grpc.CallOption) (*CommonReply, error)
	ViewChange(ctx context.Context, in *EmptyParam, opts ...grpc.CallOption) (*CommonReply, error)
	PutNewSignedCommittee(ctx context.Context, in *SignCommittee, opts ...grpc.CallOption) (*CommonReply, error)
}

type hybridConsensusHelpClient struct {
	cc *grpc.ClientConn
}

func NewHybridConsensusHelpClient(cc *grpc.ClientConn) HybridConsensusHelpClient {
	return &hybridConsensusHelpClient{cc}
}

func (c *hybridConsensusHelpClient) PutBlock(ctx context.Context, in *TruePbftBlock, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, "/truechain.HybridConsensusHelp/PutBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hybridConsensusHelpClient) ViewChange(ctx context.Context, in *EmptyParam, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, "/truechain.HybridConsensusHelp/ViewChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hybridConsensusHelpClient) PutNewSignedCommittee(ctx context.Context, in *SignCommittee, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, "/truechain.HybridConsensusHelp/PutNewSignedCommittee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HybridConsensusHelpServer is the server API for HybridConsensusHelp service.
type HybridConsensusHelpServer interface {
	PutBlock(context.Context, *TruePbftBlock) (*CommonReply, error)
	ViewChange(context.Context, *EmptyParam) (*CommonReply, error)
	PutNewSignedCommittee(context.Context, *SignCommittee) (*CommonReply, error)
}

func RegisterHybridConsensusHelpServer(s *grpc.Server, srv HybridConsensusHelpServer) {
	s.RegisterService(&_HybridConsensusHelp_serviceDesc, srv)
}

func _HybridConsensusHelp_PutBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TruePbftBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HybridConsensusHelpServer).PutBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/truechain.HybridConsensusHelp/PutBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HybridConsensusHelpServer).PutBlock(ctx, req.(*TruePbftBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _HybridConsensusHelp_ViewChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HybridConsensusHelpServer).ViewChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/truechain.HybridConsensusHelp/ViewChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HybridConsensusHelpServer).ViewChange(ctx, req.(*EmptyParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _HybridConsensusHelp_PutNewSignedCommittee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCommittee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HybridConsensusHelpServer).PutNewSignedCommittee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/truechain.HybridConsensusHelp/PutNewSignedCommittee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HybridConsensusHelpServer).PutNewSignedCommittee(ctx, req.(*SignCommittee))
	}
	return interceptor(ctx, in, info, handler)
}

var _HybridConsensusHelp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "truechain.HybridConsensusHelp",
	HandlerType: (*HybridConsensusHelpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutBlock",
			Handler:    _HybridConsensusHelp_PutBlock_Handler,
		},
		{
			MethodName: "ViewChange",
			Handler:    _HybridConsensusHelp_ViewChange_Handler,
		},
		{
			MethodName: "PutNewSignedCommittee",
			Handler:    _HybridConsensusHelp_PutNewSignedCommittee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "HybridConsensusHelp.proto",
}

// PyHybConsensusClient is the client API for PyHybConsensus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PyHybConsensusClient interface {
	Start(ctx context.Context, in *BftPrivateKey, opts ...grpc.CallOption) (*CommonReply, error)
	Stop(ctx context.Context, in *EmptyParam, opts ...grpc.CallOption) (*CommonReply, error)
	MembersNodes(ctx context.Context, in *Nodes, opts ...grpc.CallOption) (*CommonReply, error)
	SetTransactions(ctx context.Context, in *Transactions, opts ...grpc.CallOption) (*CommonReply, error)
}

type pyHybConsensusClient struct {
	cc *grpc.ClientConn
}

func NewPyHybConsensusClient(cc *grpc.ClientConn) PyHybConsensusClient {
	return &pyHybConsensusClient{cc}
}

func (c *pyHybConsensusClient) Start(ctx context.Context, in *BftPrivateKey, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, "/truechain.PyHybConsensus/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pyHybConsensusClient) Stop(ctx context.Context, in *EmptyParam, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, "/truechain.PyHybConsensus/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pyHybConsensusClient) MembersNodes(ctx context.Context, in *Nodes, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, "/truechain.PyHybConsensus/MembersNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pyHybConsensusClient) SetTransactions(ctx context.Context, in *Transactions, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, "/truechain.PyHybConsensus/SetTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PyHybConsensusServer is the server API for PyHybConsensus service.
type PyHybConsensusServer interface {
	Start(context.Context, *BftPrivateKey) (*CommonReply, error)
	Stop(context.Context, *EmptyParam) (*CommonReply, error)
	MembersNodes(context.Context, *Nodes) (*CommonReply, error)
	SetTransactions(context.Context, *Transactions) (*CommonReply, error)
}

func RegisterPyHybConsensusServer(s *grpc.Server, srv PyHybConsensusServer) {
	s.RegisterService(&_PyHybConsensus_serviceDesc, srv)
}

func _PyHybConsensus_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BftPrivateKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PyHybConsensusServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/truechain.PyHybConsensus/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PyHybConsensusServer).Start(ctx, req.(*BftPrivateKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _PyHybConsensus_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PyHybConsensusServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/truechain.PyHybConsensus/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PyHybConsensusServer).Stop(ctx, req.(*EmptyParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _PyHybConsensus_MembersNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nodes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PyHybConsensusServer).MembersNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/truechain.PyHybConsensus/MembersNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PyHybConsensusServer).MembersNodes(ctx, req.(*Nodes))
	}
	return interceptor(ctx, in, info, handler)
}

func _PyHybConsensus_SetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transactions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PyHybConsensusServer).SetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/truechain.PyHybConsensus/SetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PyHybConsensusServer).SetTransactions(ctx, req.(*Transactions))
	}
	return interceptor(ctx, in, info, handler)
}

var _PyHybConsensus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "truechain.PyHybConsensus",
	HandlerType: (*PyHybConsensusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _PyHybConsensus_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _PyHybConsensus_Stop_Handler,
		},
		{
			MethodName: "MembersNodes",
			Handler:    _PyHybConsensus_MembersNodes_Handler,
		},
		{
			MethodName: "SetTransactions",
			Handler:    _PyHybConsensus_SetTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "HybridConsensusHelp.proto",
}

func init() {
	proto.RegisterFile("HybridConsensusHelp.proto", fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059)
}

var fileDescriptor_HybridConsensusHelp_8c3bcc73b96f0059 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0x4d, 0xdb, 0xad, 0xa7, 0x1d, 0x0c, 0x8f, 0x6d, 0x61, 0x42, 0x68, 0x32, 0x42, 0x8c,
	0x0b, 0x76, 0x51, 0x60, 0x20, 0x10, 0x48, 0xfb, 0x41, 0xab, 0x34, 0x56, 0x45, 0x6e, 0xd9, 0xbd,
	0x9b, 0x9c, 0xb5, 0xd6, 0x9a, 0x1f, 0xc5, 0xce, 0xb6, 0xbc, 0x00, 0xe2, 0x35, 0xb9, 0xe7, 0x21,
	0x90, 0x9d, 0xb4, 0x4d, 0xd0, 0x3a, 0x89, 0xbb, 0xf3, 0x1d, 0x9f, 0x9f, 0xef, 0xf3, 0xf1, 0x31,
	0x3c, 0xed, 0x65, 0xa3, 0x44, 0xf8, 0xc7, 0x51, 0x28, 0x31, 0x94, 0xa9, 0xec, 0xe1, 0x34, 0xde,
	0x8f, 0x93, 0x48, 0x45, 0xa4, 0xa5, 0x92, 0x14, 0xbd, 0x09, 0x17, 0x21, 0xed, 0x00, 0x7c, 0x0b,
	0x62, 0x95, 0xb9, 0x3c, 0xe1, 0x01, 0x7d, 0x01, 0x6b, 0x47, 0x97, 0xca, 0x4d, 0xc4, 0x35, 0x57,
	0x78, 0x86, 0x19, 0x21, 0x50, 0x8f, 0xaf, 0x30, 0x73, 0xac, 0x5d, 0x6b, 0xaf, 0xc5, 0x8c, 0x4d,
	0xdf, 0xc3, 0xda, 0x40, 0x8c, 0xc3, 0xe3, 0x28, 0x08, 0x84, 0x52, 0x88, 0x64, 0x1d, 0xec, 0x40,
	0x8e, 0x8b, 0x18, 0x6d, 0xea, 0x34, 0x29, 0xc6, 0xd2, 0xa9, 0xed, 0xda, 0x3a, 0x4d, 0xdb, 0xf4,
	0x15, 0xb4, 0x75, 0x4a, 0x14, 0x32, 0x8c, 0xa7, 0x19, 0x71, 0x60, 0xe5, 0x1c, 0xa5, 0xe4, 0x63,
	0x2c, 0x12, 0x67, 0x90, 0x0e, 0xa1, 0x33, 0x4c, 0x52, 0x74, 0x47, 0x97, 0xaa, 0x1f, 0xf9, 0xa8,
	0x8b, 0x71, 0xdf, 0x4f, 0x66, 0x1c, 0xb4, 0x4d, 0xb6, 0xa0, 0x19, 0xa7, 0x23, 0xcd, 0xac, 0x66,
	0xbc, 0x05, 0xd2, 0x55, 0xe3, 0x44, 0x5c, 0xeb, 0x03, 0x3b, 0xaf, 0x5a, 0x40, 0x7a, 0x00, 0x0d,
	0x5d, 0x4d, 0x92, 0x37, 0xd0, 0x08, 0xb5, 0xe1, 0x58, 0xbb, 0xf6, 0x5e, 0xbb, 0xbb, 0xbd, 0x3f,
	0xbf, 0x8c, 0xfd, 0x72, 0x5b, 0x96, 0x47, 0xd1, 0x1b, 0xd8, 0x98, 0xb9, 0x8f, 0xa6, 0x91, 0x77,
	0xd5, 0x43, 0xee, 0xa3, 0x21, 0xd0, 0x4f, 0x83, 0x11, 0xe6, 0xb4, 0x6c, 0x56, 0x20, 0xb2, 0x03,
	0xab, 0xa7, 0x5c, 0x7e, 0x17, 0x81, 0x50, 0x86, 0x9a, 0xcd, 0xe6, 0x58, 0x93, 0x3b, 0xe5, 0xf2,
	0x87, 0x44, 0xdf, 0x90, 0xb3, 0xd9, 0x0c, 0x6a, 0x89, 0x43, 0x11, 0xa0, 0x53, 0x37, 0x6e, 0x63,
	0xd3, 0x3f, 0x16, 0x34, 0x87, 0xb7, 0x27, 0x5c, 0x71, 0x42, 0xa1, 0x73, 0xe8, 0x79, 0x51, 0x1a,
	0xaa, 0x7e, 0x14, 0x7a, 0xf9, 0x85, 0xd5, 0x59, 0xc5, 0x47, 0x9e, 0x40, 0xc3, 0x4d, 0x84, 0x87,
	0x45, 0xd7, 0x1c, 0x54, 0xe8, 0xd8, 0xff, 0xd0, 0x79, 0x06, 0x2d, 0x86, 0x9e, 0x88, 0x05, 0x86,
	0xca, 0x74, 0xee, 0xb0, 0x85, 0x43, 0x0b, 0x3c, 0x0c, 0x74, 0x79, 0xa7, 0x91, 0x0b, 0xcc, 0x91,
	0x16, 0xe1, 0xf2, 0x6c, 0x1a, 0x71, 0xdf, 0x69, 0x9a, 0x9c, 0x19, 0x24, 0x1d, 0xb0, 0x2e, 0x9c,
	0x15, 0x13, 0x6c, 0x5d, 0x68, 0xc4, 0x9c, 0xd5, 0x1c, 0x31, 0x8d, 0x06, 0x4e, 0x2b, 0x47, 0x03,
	0x2d, 0xb7, 0xc7, 0xe5, 0xc4, 0x01, 0x53, 0xc0, 0xd8, 0xf4, 0x1d, 0xb4, 0x87, 0x09, 0x0f, 0x25,
	0xf7, 0x94, 0x88, 0x42, 0xf2, 0x12, 0xea, 0x3e, 0x57, 0xdc, 0x48, 0x6d, 0x77, 0x1f, 0x97, 0x87,
	0x64, 0xee, 0x84, 0x99, 0x63, 0xfa, 0x51, 0xbf, 0x95, 0x79, 0x96, 0x24, 0x7b, 0x60, 0xab, 0xdb,
	0xd9, 0x68, 0xb7, 0x2a, 0xa3, 0x9d, 0x47, 0x31, 0x1d, 0x42, 0x7f, 0x5a, 0xb0, 0x56, 0x19, 0x2c,
	0x39, 0x80, 0xe6, 0xc4, 0x0c, 0xb7, 0x68, 0xfa, 0xfc, 0x8e, 0x97, 0x51, 0x7a, 0x02, 0xac, 0x88,
	0x26, 0xaf, 0xf3, 0x9e, 0x35, 0x93, 0xb4, 0x7d, 0x77, 0x4f, 0x69, 0x9a, 0xce, 0xf7, 0xc2, 0x5e,
	0xec, 0x45, 0xf7, 0xb7, 0x05, 0x1b, 0x77, 0xac, 0x2a, 0xf9, 0x0a, 0xab, 0x6e, 0x5a, 0x50, 0x73,
	0x96, 0x51, 0xd9, 0x29, 0x6b, 0x2c, 0xad, 0x17, 0x7d, 0x40, 0xbe, 0x00, 0x5c, 0x08, 0xbc, 0x39,
	0x9e, 0xf0, 0x70, 0x8c, 0x64, 0xb3, 0x14, 0xb7, 0x58, 0xf8, 0x7b, 0xd2, 0xcf, 0x60, 0xd3, 0x4d,
	0x55, 0x1f, 0x6f, 0xf4, 0xae, 0xa3, 0xbf, 0xd8, 0xf6, 0x32, 0x97, 0xca, 0x3f, 0xb0, 0xbc, 0x58,
	0xf7, 0x57, 0x0d, 0x1e, 0xba, 0x59, 0x2f, 0x1b, 0xcd, 0x25, 0x92, 0xcf, 0xd0, 0x18, 0x28, 0x9e,
	0xa8, 0x4a, 0xbd, 0xca, 0xe7, 0x73, 0x0f, 0xb9, 0x0f, 0x50, 0x1f, 0xa8, 0x28, 0xfe, 0x7f, 0x55,
	0x9f, 0xa0, 0x73, 0x8e, 0x7a, 0x51, 0x65, 0xfe, 0x19, 0xac, 0x97, 0x22, 0x8d, 0xe7, 0x9e, 0xdc,
	0x13, 0x78, 0x34, 0x40, 0x55, 0x79, 0x6e, 0xcb, 0xa6, 0xbd, 0xbc, 0xca, 0xa8, 0x69, 0xbe, 0xe0,
	0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x17, 0xf0, 0x9a, 0x9f, 0x05, 0x00, 0x00,
}
