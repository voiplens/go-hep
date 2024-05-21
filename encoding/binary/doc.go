// Package binary provides the HEP3 network protocol Encoder/Decoder implementation as specified here:
// https://github.com/sipcapture/HEP/blob/master/docs/HEP3_Network_Protocol_Specification_REV_36.pdf
//
// The first 4 bytes are the string "HEP3". The next 2 bytes are the length of the
// whole message (len("HEP3") + length of all the chunks we have. The next bytes
// are all the chunks created by makeChunks()
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// | "HEP3" (4 bytes) | Total len (2 bytes) | Chunks (n bytes) |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
// Chunks have the following structure:
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+- +-+-+-+
// | Chunk Vendor ID (2 bytes) | Chunk Type ID (2 bytes) | Chunk Length (2 bytes) | Chunk Data (n bytes)  |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+- +-+-+-+
package binary // import "go.voiplens.io/hep/encoding/binary"
