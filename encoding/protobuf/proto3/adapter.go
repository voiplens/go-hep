package protobufv3

import (
	"go.voiplens.io/hep"
)

var _ hep.Encoder = (*protobufEncoder)(nil)
var _ hep.Decoder = (*protobufDecoder)(nil)

type protobufEncoder struct{}
type protobufDecoder struct{}

func NewProtobufEncoder() hep.Encoder {
	return &protobufEncoder{}
}

func NewProtobufDecoder() hep.Decoder {
	return &protobufDecoder{}
}

func (p *protobufEncoder) Encode(msg *hep.Message) (data []byte, err error) {
	x := &HEP{
		Protocol:  uint32(msg.IPProtocolID),
		Version:   uint32(msg.IPProtocolFamily),
		SrcIp:     string(msg.SrcIP),
		DstIp:     string(msg.DstIP),
		SrcPort:   uint32(msg.SrcPort),
		DstPort:   uint32(msg.DstPort),
		Tsec:      msg.Tsec,
		Tmsec:     msg.Tmsec,
		ProtoType: uint32(msg.ProtoType),
		NodeId:    msg.NodeID,
		NodePw:    msg.NodePW,
		Payload:   msg.Payload,
		CId:       msg.CID,
	}
	return x.MarshalVT()
}

func (p *protobufDecoder) Decode(data []byte) (msg *hep.Message, err error) {
	x := HEPFromVTPool()
	err = x.UnmarshalVT(data)
	hepMsg := &hep.Message{
		IPProtocolID:     byte(x.GetProtocol()),
		IPProtocolFamily: byte(x.GetVersion()),
		SrcIP:            []byte(x.GetSrcIp()),
		DstIP:            []byte(x.GetDstIp()),
		SrcPort:          uint16(x.GetSrcPort()),
		DstPort:          uint16(x.GetDstPort()),
		Tsec:             x.GetTsec(),
		Tmsec:            x.GetTmsec(),
		ProtoType:        byte(x.GetProtoType()),
		NodeID:           x.GetNodeId(),
		NodePW:           x.GetNodePw(),
		Payload:          x.GetPayload(),
		CID:              x.GetCId(),
	}
	x.ReturnToVTPool()
	return hepMsg, err
}
