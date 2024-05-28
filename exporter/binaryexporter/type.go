package binaryexporter // import "go.voiplens.io/hep/exporter/binaryexporter"

// NetworktType represents a type of network used by the HEP Binary exporter.
type NetworktType string

const (
	NetworkTLS NetworktType = "tls"
	NetworkTCP NetworktType = "tcp"
	NetworkUDP NetworktType = "udp"
)
