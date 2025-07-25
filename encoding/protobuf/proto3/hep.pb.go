// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: hep.proto

package protobufv3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// HEP represents HEP packet
type HEP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       uint32                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Protocol      uint32                 `protobuf:"varint,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SrcIp         string                 `protobuf:"bytes,3,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	DstIp         string                 `protobuf:"bytes,4,opt,name=dst_ip,json=dstIp,proto3" json:"dst_ip,omitempty"`
	SrcPort       uint32                 `protobuf:"varint,5,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstPort       uint32                 `protobuf:"varint,6,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
	Tsec          uint32                 `protobuf:"varint,7,opt,name=tsec,proto3" json:"tsec,omitempty"`
	Tmsec         uint32                 `protobuf:"varint,8,opt,name=tmsec,proto3" json:"tmsec,omitempty"`
	ProtoType     uint32                 `protobuf:"varint,9,opt,name=proto_type,json=protoType,proto3" json:"proto_type,omitempty"`
	NodeId        uint32                 `protobuf:"varint,10,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodePw        string                 `protobuf:"bytes,11,opt,name=node_pw,json=nodePw,proto3" json:"node_pw,omitempty"`
	Payload       []byte                 `protobuf:"bytes,12,opt,name=payload,proto3" json:"payload,omitempty"`
	CId           []byte                 `protobuf:"bytes,13,opt,name=c_id,json=cId,proto3" json:"c_id,omitempty"`
	Vlan          uint32                 `protobuf:"varint,14,opt,name=vlan,proto3" json:"vlan,omitempty"`
	NodeName      string                 `protobuf:"bytes,15,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HEP) Reset() {
	*x = HEP{}
	mi := &file_hep_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HEP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HEP) ProtoMessage() {}

func (x *HEP) ProtoReflect() protoreflect.Message {
	mi := &file_hep_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HEP.ProtoReflect.Descriptor instead.
func (*HEP) Descriptor() ([]byte, []int) {
	return file_hep_proto_rawDescGZIP(), []int{0}
}

func (x *HEP) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *HEP) GetProtocol() uint32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *HEP) GetSrcIp() string {
	if x != nil {
		return x.SrcIp
	}
	return ""
}

func (x *HEP) GetDstIp() string {
	if x != nil {
		return x.DstIp
	}
	return ""
}

func (x *HEP) GetSrcPort() uint32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *HEP) GetDstPort() uint32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

func (x *HEP) GetTsec() uint32 {
	if x != nil {
		return x.Tsec
	}
	return 0
}

func (x *HEP) GetTmsec() uint32 {
	if x != nil {
		return x.Tmsec
	}
	return 0
}

func (x *HEP) GetProtoType() uint32 {
	if x != nil {
		return x.ProtoType
	}
	return 0
}

func (x *HEP) GetNodeId() uint32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *HEP) GetNodePw() string {
	if x != nil {
		return x.NodePw
	}
	return ""
}

func (x *HEP) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *HEP) GetCId() []byte {
	if x != nil {
		return x.CId
	}
	return nil
}

func (x *HEP) GetVlan() uint32 {
	if x != nil {
		return x.Vlan
	}
	return 0
}

func (x *HEP) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

var File_hep_proto protoreflect.FileDescriptor

const file_hep_proto_rawDesc = "" +
	"\n" +
	"\thep.proto\x12\x18hep.encoding.protobuf.v3\"\xf8\x02\n" +
	"\x03HEP\x12\x18\n" +
	"\aversion\x18\x01 \x01(\rR\aversion\x12\x1a\n" +
	"\bprotocol\x18\x02 \x01(\rR\bprotocol\x12\x15\n" +
	"\x06src_ip\x18\x03 \x01(\tR\x05srcIp\x12\x15\n" +
	"\x06dst_ip\x18\x04 \x01(\tR\x05dstIp\x12\x19\n" +
	"\bsrc_port\x18\x05 \x01(\rR\asrcPort\x12\x19\n" +
	"\bdst_port\x18\x06 \x01(\rR\adstPort\x12\x12\n" +
	"\x04tsec\x18\a \x01(\rR\x04tsec\x12\x14\n" +
	"\x05tmsec\x18\b \x01(\rR\x05tmsec\x12\x1d\n" +
	"\n" +
	"proto_type\x18\t \x01(\rR\tprotoType\x12\x17\n" +
	"\anode_id\x18\n" +
	" \x01(\rR\x06nodeId\x12\x17\n" +
	"\anode_pw\x18\v \x01(\tR\x06nodePw\x12\x18\n" +
	"\apayload\x18\f \x01(\fR\apayload\x12\x11\n" +
	"\x04c_id\x18\r \x01(\fR\x03cId\x12\x12\n" +
	"\x04vlan\x18\x0e \x01(\rR\x04vlan\x12\x1b\n" +
	"\tnode_name\x18\x0f \x01(\tR\bnodeNameB\xdf\x01\n" +
	"\x1ccom.hep.encoding.protobuf.v3B\bHepProtoP\x01Z2go.voiplens.io/hep/encoding/protobuf/v3;protobufv3\xa2\x02\x03HEP\xaa\x02\x18Hep.Encoding.Protobuf.V3\xca\x02\x18Hep\\Encoding\\Protobuf\\V3\xe2\x02$Hep\\Encoding\\Protobuf\\V3\\GPBMetadata\xea\x02\x1bHep::Encoding::Protobuf::V3b\x06proto3"

var (
	file_hep_proto_rawDescOnce sync.Once
	file_hep_proto_rawDescData []byte
)

func file_hep_proto_rawDescGZIP() []byte {
	file_hep_proto_rawDescOnce.Do(func() {
		file_hep_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_hep_proto_rawDesc), len(file_hep_proto_rawDesc)))
	})
	return file_hep_proto_rawDescData
}

var file_hep_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hep_proto_goTypes = []any{
	(*HEP)(nil), // 0: hep.encoding.protobuf.v3.HEP
}
var file_hep_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hep_proto_init() }
func file_hep_proto_init() {
	if File_hep_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_hep_proto_rawDesc), len(file_hep_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hep_proto_goTypes,
		DependencyIndexes: file_hep_proto_depIdxs,
		MessageInfos:      file_hep_proto_msgTypes,
	}.Build()
	File_hep_proto = out.File
	file_hep_proto_goTypes = nil
	file_hep_proto_depIdxs = nil
}
