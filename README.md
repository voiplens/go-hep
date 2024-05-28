# go-hep

## Description

`go.voiplens.io/hep` or `go-hep` is a Go library implementing the HEP3 encapsulation protocol.
This project doesn't provide a CLI as it's primary usage is to be integrated in other projects.
The documentation is available [here](https://pkg.go.dev/go.voiplens.io/hep).

## Features

- [X] Encoder/Decoder for HEP3 binary protocol
- [X] Encoder/Decoder for Protobuf V2 Messages (gogo-proto-fast)
- [X] Encoder/Decoder for Protobuf V3 Messages (vtprotobuf)
- [X] Basic UDP/TCP/TLS exporter
- [ ] Exporter Logs and Metrics
- [ ] Retry & Circuit breaker support for exporter
- [ ] WAL support for exporter
- [ ] Support for Vendor specific extensions/chunks

## Usage
### Basic HEP3 Exporter

See example in the `examples/exporter` directory.

## License

`go-hep` is licensed under the [license name](LICENSE).
