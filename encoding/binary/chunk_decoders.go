package binary

import (
	"encoding/binary"
	"fmt"

	"go.voiplens.io/hep"
)

type ChunkDecoder interface {
	Decode(msg *hep.Message, chunkBody []byte) error
}

type ipProtocolFamilyChunk struct{}

var _ ChunkDecoder = (*ipProtocolFamilyChunk)(nil)

func (v *ipProtocolFamilyChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 1 {
		return fmt.Errorf("IP Protocol Family chunk should be 1 byte long but is %d", len(chunkBody))
	}
	msg.IPProtocolFamily = chunkBody[0]
	return nil
}

type ipProtocolIDChunk struct{}

var _ ChunkDecoder = (*ipProtocolIDChunk)(nil)

func (v *ipProtocolIDChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 1 {
		return fmt.Errorf("IP Protocol ID chunk should be 1 byte long but is %d", len(chunkBody))
	}
	msg.IPProtocolID = chunkBody[0]
	return nil
}

type protoTypeChunk struct{}

var _ ChunkDecoder = (*protoTypeChunk)(nil)

func (v *protoTypeChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 1 {
		return fmt.Errorf("Protocol Type chunk should be 1 byte long but is %d", len(chunkBody))
	}
	msg.ProtoType = chunkBody[0]
	return nil
}

type srcPortChunk struct{}

var _ ChunkDecoder = (*srcPortChunk)(nil)

func (v *srcPortChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 2 {
		return fmt.Errorf("SrcPort chunk should be 2 byte long but is %d", len(chunkBody))
	}
	msg.SrcPort = binary.BigEndian.Uint16(chunkBody)
	return nil
}

type dstPortChunk struct{}

var _ ChunkDecoder = (*dstPortChunk)(nil)

func (v *dstPortChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 2 {
		return fmt.Errorf("DstPort chunk should be 2 byte long but is %d", len(chunkBody))
	}
	msg.DstPort = binary.BigEndian.Uint16(chunkBody)
	return nil
}

type vlanChunk struct{}

var _ ChunkDecoder = (*vlanChunk)(nil)

func (v *vlanChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 2 {
		return fmt.Errorf("VLAN chunk should be 2 byte long but is %d", len(chunkBody))
	}
	msg.Vlan = binary.BigEndian.Uint16(chunkBody)
	return nil
}

type tsecChunk struct{}

var _ ChunkDecoder = (*tsecChunk)(nil)

func (v *tsecChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 4 {
		return fmt.Errorf("Tsec chunk should be 4 byte long but is %d", len(chunkBody))
	}
	msg.Tsec = binary.BigEndian.Uint32(chunkBody)
	return nil
}

type tmsecChunk struct{}

var _ ChunkDecoder = (*tmsecChunk)(nil)

func (v *tmsecChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 4 {
		return fmt.Errorf("Tmsec chunk should be 4 byte long but is %d", len(chunkBody))
	}
	msg.Tmsec = binary.BigEndian.Uint32(chunkBody)
	return nil
}

type nodeIDChunk struct{}

var _ ChunkDecoder = (*nodeIDChunk)(nil)

func (v *nodeIDChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 4 {
		return fmt.Errorf("NodeID chunk should be 4 byte long but is %d", len(chunkBody))
	}
	msg.NodeID = binary.BigEndian.Uint32(chunkBody)
	return nil
}

type srcIP4Chunk struct{}

var _ ChunkDecoder = (*srcIP4Chunk)(nil)

func (v *srcIP4Chunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 4 {
		return fmt.Errorf("SrcIP chunk should be 4 or 16 byte long but is %d", len(chunkBody))
	}
	msg.SrcIP = chunkBody
	return nil
}

type dstIP4Chunk struct{}

var _ ChunkDecoder = (*dstIP4Chunk)(nil)

func (v *dstIP4Chunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 4 {
		return fmt.Errorf("DstIP chunk should be 4 or 16 byte long but is %d", len(chunkBody))
	}
	msg.DstIP = chunkBody
	return nil
}

type srcIP6Chunk struct{}

var _ ChunkDecoder = (*srcIP6Chunk)(nil)

func (v *srcIP6Chunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 16 {
		return fmt.Errorf("SrcIP chunk should be 4 or 16 byte long but is %d", len(chunkBody))
	}
	msg.SrcIP = chunkBody
	return nil
}

type dstIP6Chunk struct{}

var _ ChunkDecoder = (*dstIP6Chunk)(nil)

func (v *dstIP6Chunk) Decode(msg *hep.Message, chunkBody []byte) error {
	if len(chunkBody) != 16 {
		return fmt.Errorf("DstIP chunk should be 4 or 16 byte long but is %d", len(chunkBody))
	}
	msg.DstIP = chunkBody
	return nil
}

type nodePWChunk struct{}

var _ ChunkDecoder = (*nodePWChunk)(nil)

func (v *nodePWChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	msg.NodePW = string(chunkBody)
	return nil
}

type payloadChunk struct{}

var _ ChunkDecoder = (*payloadChunk)(nil)

func (v *payloadChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	msg.Payload = chunkBody
	return nil
}

type cidChunk struct{}

var _ ChunkDecoder = (*cidChunk)(nil)

func (v *cidChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	msg.CID = chunkBody
	return nil
}

type nodeNameChunk struct{}

var _ ChunkDecoder = (*nodeNameChunk)(nil)

func (v *nodeNameChunk) Decode(msg *hep.Message, chunkBody []byte) error {
	msg.NodeName = string(chunkBody)
	return nil
}
