// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/collection.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/collection.proto
	common/common.proto
	common/configtx.proto
	common/configuration.proto
	common/ledger.proto
	common/policies.proto

It has these top-level messages:
	CollectionConfig
	StaticCollectionConfig
	CollectionPolicyConfig
	CollectionCriteria
	LastConfig
	Metadata
	MetadataSignature
	Header
	ChannelHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigEnvelope
	ConfigGroupSchema
	ConfigValueSchema
	ConfigPolicySchema
	Config
	ConfigUpdateEnvelope
	ConfigUpdate
	ConfigGroup
	ConfigValue
	ConfigPolicy
	ConfigSignature
	HashingAlgorithm
	BlockDataHashingStructure
	OrdererAddresses
	Consortium
	Capabilities
	Capability
	BlockchainInfo
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
	ImplicitMetaPolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// CollectionConfig defines the configuration of a collection object;
// it currently contains a single, static type.
// Dynamic collections are deferred.
type CollectionConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionConfig_StaticCollectionConfig
	Payload isCollectionConfig_Payload `protobuf_oneof:"payload"`
}

func (m *CollectionConfig) Reset()                    { *m = CollectionConfig{} }
func (m *CollectionConfig) String() string            { return proto.CompactTextString(m) }
func (*CollectionConfig) ProtoMessage()               {}
func (*CollectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isCollectionConfig_Payload interface {
	isCollectionConfig_Payload()
}

type CollectionConfig_StaticCollectionConfig struct {
	StaticCollectionConfig *StaticCollectionConfig `protobuf:"bytes,1,opt,name=static_collection_config,json=staticCollectionConfig,oneof"`
}

func (*CollectionConfig_StaticCollectionConfig) isCollectionConfig_Payload() {}

func (m *CollectionConfig) GetPayload() isCollectionConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionConfig) GetStaticCollectionConfig() *StaticCollectionConfig {
	if x, ok := m.GetPayload().(*CollectionConfig_StaticCollectionConfig); ok {
		return x.StaticCollectionConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CollectionConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CollectionConfig_OneofMarshaler, _CollectionConfig_OneofUnmarshaler, _CollectionConfig_OneofSizer, []interface{}{
		(*CollectionConfig_StaticCollectionConfig)(nil),
	}
}

func _CollectionConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CollectionConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionConfig_StaticCollectionConfig:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StaticCollectionConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CollectionConfig.Payload has unexpected type %T", x)
	}
	return nil
}

func _CollectionConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CollectionConfig)
	switch tag {
	case 1: // payload.static_collection_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StaticCollectionConfig)
		err := b.DecodeMessage(msg)
		m.Payload = &CollectionConfig_StaticCollectionConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CollectionConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CollectionConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionConfig_StaticCollectionConfig:
		s := proto.Size(x.StaticCollectionConfig)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// StaticCollectionConfig constitutes the configuration parameters of a
// static collection object. Static collections are collections that are
// known at chaincode instantiation time, and that cannot be changed.
// Dynamic collections are deferred.
type StaticCollectionConfig struct {
	// the name of the collection inside the denoted chaincode
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define which orgs have access to this collection’s private data
	MemberOrgsPolicy *CollectionPolicyConfig `protobuf:"bytes,2,opt,name=member_orgs_policy,json=memberOrgsPolicy" json:"member_orgs_policy,omitempty"`
}

func (m *StaticCollectionConfig) Reset()                    { *m = StaticCollectionConfig{} }
func (m *StaticCollectionConfig) String() string            { return proto.CompactTextString(m) }
func (*StaticCollectionConfig) ProtoMessage()               {}
func (*StaticCollectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StaticCollectionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StaticCollectionConfig) GetMemberOrgsPolicy() *CollectionPolicyConfig {
	if m != nil {
		return m.MemberOrgsPolicy
	}
	return nil
}

// Collection policy configuration. Initially, the configuration can only
// contain a SignaturePolicy. In the future, the SignaturePolicy may be a
// more general Policy. Instead of containing the actual policy, the
// configuration may in the future contain a string reference to a policy.
type CollectionPolicyConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionPolicyConfig_SignaturePolicy
	Payload isCollectionPolicyConfig_Payload `protobuf_oneof:"payload"`
}

func (m *CollectionPolicyConfig) Reset()                    { *m = CollectionPolicyConfig{} }
func (m *CollectionPolicyConfig) String() string            { return proto.CompactTextString(m) }
func (*CollectionPolicyConfig) ProtoMessage()               {}
func (*CollectionPolicyConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isCollectionPolicyConfig_Payload interface {
	isCollectionPolicyConfig_Payload()
}

type CollectionPolicyConfig_SignaturePolicy struct {
	SignaturePolicy *SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=signature_policy,json=signaturePolicy,oneof"`
}

func (*CollectionPolicyConfig_SignaturePolicy) isCollectionPolicyConfig_Payload() {}

func (m *CollectionPolicyConfig) GetPayload() isCollectionPolicyConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionPolicyConfig) GetSignaturePolicy() *SignaturePolicyEnvelope {
	if x, ok := m.GetPayload().(*CollectionPolicyConfig_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CollectionPolicyConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CollectionPolicyConfig_OneofMarshaler, _CollectionPolicyConfig_OneofUnmarshaler, _CollectionPolicyConfig_OneofSizer, []interface{}{
		(*CollectionPolicyConfig_SignaturePolicy)(nil),
	}
}

func _CollectionPolicyConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CollectionPolicyConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionPolicyConfig_SignaturePolicy:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SignaturePolicy); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CollectionPolicyConfig.Payload has unexpected type %T", x)
	}
	return nil
}

func _CollectionPolicyConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CollectionPolicyConfig)
	switch tag {
	case 1: // payload.signature_policy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignaturePolicyEnvelope)
		err := b.DecodeMessage(msg)
		m.Payload = &CollectionPolicyConfig_SignaturePolicy{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CollectionPolicyConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CollectionPolicyConfig)
	// payload
	switch x := m.Payload.(type) {
	case *CollectionPolicyConfig_SignaturePolicy:
		s := proto.Size(x.SignaturePolicy)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// CollectionCriteria defines an element of a private data that corresponds
// to a certain transaction and collection
type CollectionCriteria struct {
	Channel    string `protobuf:"bytes,1,opt,name=channel" json:"channel,omitempty"`
	TxId       string `protobuf:"bytes,2,opt,name=tx_id,json=txId" json:"tx_id,omitempty"`
	Collection string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
	Namespace  string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *CollectionCriteria) Reset()                    { *m = CollectionCriteria{} }
func (m *CollectionCriteria) String() string            { return proto.CompactTextString(m) }
func (*CollectionCriteria) ProtoMessage()               {}
func (*CollectionCriteria) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CollectionCriteria) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *CollectionCriteria) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *CollectionCriteria) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *CollectionCriteria) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*CollectionConfig)(nil), "common.CollectionConfig")
	proto.RegisterType((*StaticCollectionConfig)(nil), "common.StaticCollectionConfig")
	proto.RegisterType((*CollectionPolicyConfig)(nil), "common.CollectionPolicyConfig")
	proto.RegisterType((*CollectionCriteria)(nil), "common.CollectionCriteria")
}

func init() { proto.RegisterFile("common/collection.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xb7, 0xff, 0x7f, 0x6e, 0xf4, 0xfa, 0xe0, 0x88, 0x38, 0x8b, 0xc8, 0x94, 0xe1, 0x83,
	0x20, 0xb4, 0xa0, 0xdf, 0x60, 0x43, 0x98, 0x30, 0x50, 0xba, 0xb7, 0xbd, 0x94, 0x34, 0xbd, 0xeb,
	0x02, 0x6d, 0x52, 0x93, 0x4c, 0x56, 0xdf, 0xfc, 0xe6, 0xb2, 0x64, 0x5b, 0x3b, 0xd9, 0x5b, 0xef,
	0xfd, 0x9d, 0x73, 0xd2, 0xd3, 0x06, 0xae, 0x99, 0x2c, 0x0a, 0x29, 0x42, 0x26, 0xf3, 0x1c, 0x99,
	0xe1, 0x52, 0x04, 0xa5, 0x92, 0x46, 0x92, 0xae, 0x03, 0x37, 0x57, 0x3b, 0x41, 0x29, 0x73, 0xce,
	0x38, 0x6a, 0x87, 0x47, 0x15, 0xf4, 0x27, 0x07, 0xcb, 0x44, 0x8a, 0x25, 0xcf, 0xc8, 0x02, 0x7c,
	0x6d, 0xa8, 0xe1, 0x2c, 0xae, 0xd3, 0x62, 0x66, 0x99, 0xdf, 0xbe, 0x6f, 0x3f, 0x9e, 0x3f, 0x0f,
	0x03, 0x97, 0x16, 0xcc, 0xad, 0xee, 0x6f, 0xc2, 0xb4, 0x15, 0x0d, 0xf4, 0x49, 0x32, 0xf6, 0xa0,
	0x57, 0xd2, 0x2a, 0x97, 0x34, 0x1d, 0x7d, 0xc3, 0xe0, 0xb4, 0x9d, 0x10, 0xe8, 0x08, 0x5a, 0xa0,
	0x3d, 0xcc, 0x8b, 0xec, 0x33, 0x99, 0x01, 0x29, 0xb0, 0x48, 0x50, 0xc5, 0x52, 0x65, 0x3a, 0xb6,
	0x35, 0x2a, 0xff, 0xdf, 0xf1, 0xeb, 0xd4, 0x49, 0x1f, 0x96, 0xbb, 0xbc, 0xa8, 0xef, 0x9c, 0xef,
	0x2a, 0xd3, 0x6e, 0x3f, 0xfa, 0x84, 0xc1, 0x69, 0x2d, 0x99, 0x41, 0x5f, 0xf3, 0x4c, 0x50, 0xb3,
	0x56, 0xb8, 0x3f, 0xc5, 0x95, 0xbe, 0x3b, 0x94, 0xde, 0x73, 0x67, 0x7c, 0x15, 0x5f, 0x98, 0xcb,
	0x12, 0xa7, 0xad, 0xe8, 0x42, 0x1f, 0xa3, 0x66, 0xdd, 0x9f, 0x36, 0x90, 0x46, 0x53, 0xc5, 0x0d,
	0x2a, 0x4e, 0x89, 0x0f, 0x3d, 0xb6, 0xa2, 0x42, 0x60, 0xbe, 0xab, 0xbb, 0x1f, 0xc9, 0x25, 0x9c,
	0x99, 0x4d, 0xcc, 0x53, 0x5b, 0xd2, 0x8b, 0x3a, 0x66, 0xf3, 0x96, 0x92, 0x21, 0x40, 0xfd, 0x53,
	0xfc, 0xff, 0x96, 0x34, 0x36, 0xe4, 0x16, 0xbc, 0xed, 0xe7, 0xd2, 0x25, 0x65, 0xe8, 0x77, 0x2c,
	0xae, 0x17, 0xe3, 0x39, 0x3c, 0x48, 0x95, 0x05, 0xab, 0xaa, 0x44, 0x95, 0x63, 0x9a, 0xa1, 0x0a,
	0x96, 0x34, 0x51, 0x9c, 0xb9, 0xdb, 0xa0, 0x77, 0x0d, 0x17, 0x4f, 0x19, 0x37, 0xab, 0x75, 0xb2,
	0x1d, 0xc3, 0x86, 0x38, 0x74, 0xe2, 0xd0, 0x89, 0x43, 0x27, 0x4e, 0xba, 0x76, 0x7c, 0xf9, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x1d, 0x13, 0xa8, 0x68, 0x83, 0x02, 0x00, 0x00,
}
