package binaryexporter // import "go.voiplens.io/hep/exporter/binaryexporter"

import (
	"errors"
	"fmt"
	"time"
)

// Config defines the configuration for the HEP Binary exporter.
type Config struct {
	// Endpoint configures the HEP server address for this network connection.
	// For TCP and UDP networks, the address has the form "host:port". The host must be a literal IP address,
	// or a host name that can be resolved to IP addresses. The port must be a literal port number or a service name.
	// If the host is a literal IPv6 address it must be enclosed in square brackets, as in "[2001:db8::1]:80" or
	// "[fe80::1%zone]:80". The zone specifies the scope of the literal IPv6 address as defined in RFC 4007.
	Endpoint string

	// Network for HEP Binary communication
	// Supported networks are: tcp, udp, tls
	Network NetworktType

	// InsecureSkipVerify will enable TLS but not verify the certificate.
	InsecureSkipVerify bool

	// KeepAlive specifies the interval between keep-alive
	// probes for an active network connection.
	// If zero, keep-alive probes are sent with a default value
	// (currently 15 seconds), if supported by the protocol and operating
	// system. Network protocols or operating systems that do
	// not support keep-alives ignore this field.
	// If negative, keep-alive probes are disabled.
	Keepalive time.Duration

	// Timeout is the timeout for every attempt to send data to the backend.
	// A zero timeout means no timeout.
	// Timeout is the maximum amount of time a dial will wait for
	// a connect to complete. The default is no timeout.
	Timeout time.Duration

	// QueueSize is the maximum number of encoded messages allowed in queue at a given time.
	QueueSize int64
}

// Validate the configuration for errors.
func (cfg *Config) Validate() error {
	invalidFields := []error{}

	if cfg.Endpoint == "" {
		invalidFields = append(invalidFields, errors.New("Endpoint is required"))
	}

	switch cfg.Network {
	case
		NetworkTLS,
		NetworkTCP,
		NetworkUDP:
	default:
		invalidFields = append(invalidFields, fmt.Errorf("Network type %q is invalid", cfg.Network))
	}

	if cfg.QueueSize < 0 {
		invalidFields = append(invalidFields, errors.New("QueueSize must be non-negative"))
	}
	// Negative timeouts are not acceptable, since all sends will fail.
	if cfg.Timeout < 0 {
		invalidFields = append(invalidFields, errors.New("Timeout must be non-negative"))
	}
	return errors.Join(invalidFields...)
}
