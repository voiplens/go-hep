package hep

// Encoder is an interface defining functions that an encoder plugin must
// satisfy.
type Encoder interface {
	// Encode takes a single HEP Message and turns it into a byte array.
	Encode(*Message) ([]byte, error)
}

// Decoder is an interface defining functions that a decoder plugin must
// satisfy.
type Decoder interface {
	// Encode takes a byte array and turns it into single HEP Message.
	Decode([]byte) (*Message, error)
}
