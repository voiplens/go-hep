package binary

import (
	"encoding/binary"
	"fmt"

	"go.voiplens.io/hep"
)

type binaryDecoder struct {
	vendocrChunkDecoder map[uint16]map[uint16]ChunkDecoder
}

var _ hep.Decoder = (*binaryDecoder)(nil)

func NewBinaryDecoder() hep.Decoder {
	vendorChunkDecoder := make(map[uint16]map[uint16]ChunkDecoder)
	vendorChunkDecoder[VendorGeneric] = genericChunkDecoders
	return &binaryDecoder{
		vendocrChunkDecoder: vendorChunkDecoder,
	}
}

var genericChunkDecoders = map[uint16]ChunkDecoder{
	TypeIPProtocolFamily: &ipProtocolFamilyChunk{},
	TypeIPProtocolID:     &ipProtocolIDChunk{},
	TypeProtoType:        &protoTypeChunk{},
	TypeSrcPort:          &srcPortChunk{},
	TypeDstPort:          &dstPortChunk{},
	TypeIP4SrcIP:         &srcIP4Chunk{},
	TypeIP4DstIP:         &dstIP4Chunk{},
	TypeIP6SrcIP:         &srcIP6Chunk{},
	TypeIP6DstIP:         &dstIP6Chunk{},
	TypeTsec:             &tsecChunk{},
	TypeTmsec:            &tmsecChunk{},
	TypeNodeID:           &nodeIDChunk{},
	TypeNodePW:           &nodePWChunk{},
	TypePayload:          &payloadChunk{},
	TypeCID:              &cidChunk{},
	TypeVlanID:           &vlanChunk{},
	TypeNodeName:         &nodeNameChunk{},
}

// Decode decodes the given packet and populates the fields of the Message struct.
// It returns an error if the packet is not a valid HEP packet or if there is an error during decoding.
func (b *binaryDecoder) Decode(data []byte) (*hep.Message, error) {
	encoding := binary.BigEndian
	length := encoding.Uint16(data[4:6])
	if int(length) != len(data) {
		return nil, fmt.Errorf("HEP packet length is %d but should be %d", len(data), length)
	}
	currentByte := uint16(6)

	msg := &hep.Message{}
	for currentByte < length {
		chunk := data[currentByte:]
		if len(chunk) < ChunkHeaderLength {
			return nil, fmt.Errorf("HEP chunk must be >= 6 byte long but is %d", len(chunk))
		}
		chunkType := encoding.Uint16(chunk[2:4])
		chunkLength := encoding.Uint16(chunk[4:6])
		if len(chunk) < int(chunkLength) || int(chunkLength) < ChunkHeaderLength {
			return nil, fmt.Errorf("HEP chunk with %d byte < chunkLength %d or chunkLength < 6", len(chunk), chunkLength)
		}

		chunkVendorId := encoding.Uint16(chunk[:2])
		vendorChunkDecoder, ok := b.vendocrChunkDecoder[chunkVendorId]
		if !ok {
			currentByte += chunkLength
			continue
		}
		decoder, ok := vendorChunkDecoder[chunkType]
		currentByte += chunkLength
		if !ok {
			continue
		}
		chunkBody := chunk[ChunkHeaderLength:chunkLength]
		if err := decoder.Decode(msg, chunkBody); err != nil {
			return nil, err
		}
	}
	return msg, nil
}
