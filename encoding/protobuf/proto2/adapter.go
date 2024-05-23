package proto2

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
	hepMsg := &HEP{
		Protocol:  uint32(msg.IPProtocolID),
		Version:   uint32(msg.IPProtocolFamily),
		SrcIP:     msg.SrcIP.String(),
		DstIP:     msg.DstIP.String(),
		SrcPort:   uint32(msg.SrcPort),
		DstPort:   uint32(msg.DstPort),
		Tsec:      msg.Tsec,
		Tmsec:     msg.Tmsec,
		ProtoType: uint32(msg.ProtoType),
		NodeID:    msg.NodeID,
		NodePW:    msg.NodePW,
		Payload:   string(msg.Payload),
		CID:       string(msg.CID),
	}
	return hepMsg.Marshal()
}

func (p *protobufDecoder) Decode(data []byte) (msg *hep.Message, err error) {
	x := HEP{}
	err = x.Unmarshal(data)
	hepMsg := &hep.Message{
		IPProtocolID:     byte(x.GetProtocol()),
		IPProtocolFamily: byte(x.GetVersion()),
		SrcIP:            []byte(x.GetSrcIP()),
		DstIP:            []byte(x.GetDstIP()),
		SrcPort:          uint16(x.GetSrcPort()),
		DstPort:          uint16(x.GetDstPort()),
		Tsec:             x.GetTsec(),
		Tmsec:            x.GetTmsec(),
		ProtoType:        byte(x.GetProtoType()),
		NodeID:           x.GetNodeID(),
		NodePW:           x.GetNodePW(),
		Payload:          []byte(x.GetPayload()),
		CID:              []byte(x.GetCID()),
	}
	return hepMsg, err
}
