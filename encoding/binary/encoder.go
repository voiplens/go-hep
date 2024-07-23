package binary

import (
	"fmt"

	"go.voiplens.io/hep"
)

var _ hep.Encoder = (*binaryEncoder)(nil)

type binaryEncoder struct{}

func NewBinaryEncoder() hep.Encoder {
	return &binaryEncoder{}
}

// Encode encodes the Message struct into a byte slice.
// It returns the encoded data and an error, if any.
func (b *binaryEncoder) Encode(msg *hep.Message) (data []byte, err error) {
	size := Size(msg)
	data = make([]byte, size)
	offset := copy(data[0:], []byte{0x48, 0x45, 0x50, 0x33, byte(size >> 8), byte(size)})
	offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x07, msg.IPProtocolFamily})
	offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x02, 0x00, 0x07, msg.IPProtocolID})

	if msg.IPProtocolFamily == 0x02 {
		if msg.SrcIP != nil {
			srcIPLen := ChunkHeaderLength + len(msg.SrcIP)
			offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x03, byte(srcIPLen >> 8), byte(srcIPLen)})
			offset += copy(data[offset:], msg.SrcIP)
		}

		if msg.DstIP != nil {
			dstIPLen := ChunkHeaderLength + len(msg.DstIP)
			offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x04, byte(dstIPLen >> 8), byte(dstIPLen)})
			offset += copy(data[offset:], msg.DstIP)
		}
	} else {
		if msg.SrcIP != nil {
			srcIPLen := ChunkHeaderLength + len(msg.SrcIP)
			offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x05, byte(srcIPLen >> 8), byte(srcIPLen)})
			offset += copy(data[offset:], msg.SrcIP)
		}

		if msg.DstIP != nil {
			dstIPLen := ChunkHeaderLength + len(msg.DstIP)
			offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x06, byte(dstIPLen >> 8), byte(dstIPLen)})
			offset += copy(data[offset:], msg.DstIP)
		}
	}

	offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x07, 0x00, 0x08, byte(msg.SrcPort >> 8), byte(msg.SrcPort)})
	offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x08, 0x00, 0x08, byte(msg.DstPort >> 8), byte(msg.DstPort)})
	if !msg.Timestamp.IsZero() && msg.Tsec == 0 && msg.Tmsec == 0 {
		nano := msg.Timestamp.UnixNano()
		msg.Tsec = uint32(nano / 1e9)
		msg.Tmsec = uint32((nano % 1e9) / 1000)
	}
	offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x09, 0x00, 0x0a, byte(msg.Tsec >> 24), byte(msg.Tsec >> 16), byte(msg.Tsec >> 8), byte(msg.Tsec)})
	offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x0a, 0x00, 0x0a, byte(msg.Tmsec >> 24), byte(msg.Tmsec >> 16), byte(msg.Tmsec >> 8), byte(msg.Tmsec)})
	offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x0b, 0x00, 0x07, msg.ProtoType})
	offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x0c, 0x00, 0x0a, byte(msg.NodeID >> 24), byte(msg.NodeID >> 16), byte(msg.NodeID >> 8), byte(msg.NodeID)})

	if msg.NodePW != "" {
		nodePWLen := ChunkHeaderLength + len(msg.NodePW)
		offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x0e, byte(nodePWLen >> 8), byte(nodePWLen)})
		offset += copy(data[offset:], msg.NodePW)
	}

	if msg.Payload != nil {
		payloadLen := ChunkHeaderLength + len(msg.Payload)
		offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x0f, byte(payloadLen >> 8), byte(payloadLen)})
		offset += copy(data[offset:], msg.Payload)
	}

	if msg.CID != nil {
		cidLen := ChunkHeaderLength + len(msg.CID)
		offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x11, byte(cidLen >> 8), byte(cidLen)})
		offset += copy(data[offset:], msg.CID)
	}

	offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x12, 0x00, 0x08, byte(msg.Vlan >> 8), byte(msg.Vlan)})

	if msg.NodeName != "" {
		nodeNameLen := ChunkHeaderLength + len(msg.NodeName)
		offset += copy(data[offset:], []byte{0x00, 0x00, 0x00, 0x13, byte(nodeNameLen >> 8), byte(nodeNameLen)})
		offset += copy(data[offset:], msg.NodeName)
	}

	if size != offset {
		return nil, fmt.Errorf("Invalid size")
	}

	return data, nil
}

// Size returns the size of the Message in bytes.
// It calculates the size based on the length of various fields in the Message.
// The size includes the length of the vendor, chunk, and the corresponding field.
// If the field is nil or empty, it is not included in the size calculation.
// The total size is the sum of all the field sizes plus a constant value of 81.
func Size(msg *hep.Message) int {
	// n := len("HEP3") + 2       // len("HEP3") + len(TotalLength)
	// n += ChunkHeaderLength + 1 // len(vendor) + len(chunk) + len(Version)
	// n += ChunkHeaderLength + 1 // len(vendor) + len(chunk) + len(Protocol)
	// n += ChunkHeaderLength + 2 // len(vendor) + len(chunk) + len(Vlan)
	// n += ChunkHeaderLength + 2 // len(vendor) + len(chunk) + len(SrcPort)
	// n += ChunkHeaderLength + 2 // len(vendor) + len(chunk) + len(DstPort)
	// n += ChunkHeaderLength + 4 // len(vendor) + len(chunk) + len(Tsec)
	// n += ChunkHeaderLength + 4 // len(vendor) + len(chunk) + len(Tmsec)
	// n += ChunkHeaderLength + 1 // len(vendor) + len(chunk) + len(ProtoType)
	// n += ChunkHeaderLength + 4 // len(vendor) + len(chunk) + len(NodeID)
	n := 81
	if msg.SrcIP != nil {
		n += ChunkHeaderLength + len(msg.SrcIP) // len(vendor) + len(chunk) + len(SrcIP)
	}
	if msg.DstIP != nil {
		n += ChunkHeaderLength + len(msg.DstIP) // len(vendor) + len(chunk) + len(DstIP)
	}
	if msg.NodePW != "" {
		n += ChunkHeaderLength + len(msg.NodePW) // len(vendor) + len(chunk) + len(NodePW)
	}
	if msg.Payload != nil {
		n += ChunkHeaderLength + len(msg.Payload) // len(vendor) + len(chunk) + len(Payload)
	}
	if msg.CID != nil {
		n += ChunkHeaderLength + len(msg.CID) // len(vendor) + len(chunk) + len(CID)
	}
	if msg.NodeName != "" {
		n += ChunkHeaderLength + len(msg.NodeName) // len(vendor) + len(chunk) + len(NodeName)
	}
	return n
}
