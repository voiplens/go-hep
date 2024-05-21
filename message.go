package hep

import (
	"net"
	"time"
)

// Message represents a HEP message
type Message struct {
	IPProtocolFamily byte
	IPProtocolID     byte
	SrcIP            net.IP
	DstIP            net.IP
	SrcPort          uint16
	DstPort          uint16
	Tsec             uint32
	Tmsec            uint32
	ProtoType        byte
	NodeID           uint32
	NodePW           string
	Payload          []byte
	CID              []byte
	Vlan             uint16
	ProtoString      string
	Timestamp        time.Time
	NodeName         string
	TargetName       string
	SID              string
}
