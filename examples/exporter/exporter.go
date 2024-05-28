package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"time"

	"go.voiplens.io/hep"
	"go.voiplens.io/hep/encoding/binary"
	"go.voiplens.io/hep/exporter/binaryexporter"
)

const SIP_OPTION = `OPTIONS sip:172.16.105.138:5060;transport=tcp SIP/2.0
Via: SIP/2.0/TCP 172.16.105.10;branch=z9hG4bK9c6c.b38f9343000000000000000000000000.0
To: <sip:172.16.105.138:5060;transport=tcp>
From: <sip:UWT-SBC-1@172.16.106.10>;tag=5f941dfd99695412a5639025073fb7e8-69947a8e
CSeq: 10 OPTIONS
Call-ID: 4c128aba-8ee7-123d-6aa4-0242ac106a80
Max-Forwards: 70
Content-Length: 0

`

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("Starting HEP Exporter...")
	encoder := binary.NewBinaryEncoder()
	cfg := &binaryexporter.Config{
		Endpoint:  "localhost:9060",
		Network:   binaryexporter.NetworkUDP,
		QueueSize: 1000,
	}
	exporter, err := binaryexporter.NewBinaryExporter(cfg, logger, encoder)
	if err != nil {
		logger.Error("Failed to start HEP Exporter", "err", err)
		return
	}

	ctx := context.Background()
	err = exporter.Start(ctx)
	if err != nil {
		logger.Error("Failed to start HEP Exporter", "err", err)
		return
	}

	timestamp := time.Now()
	tsec := timestamp.Unix()
	tmsec := timestamp.UnixMicro() - tsec*1000000
	hepMsg := &hep.Message{
		IPProtocolFamily: byte(0x2),
		IPProtocolID:     byte(0x11),
		SrcIP:            net.ParseIP("172.16.105.10").To4(),
		DstIP:            net.ParseIP("172.16.105.138").To4(),
		DstPort:          5060,
		SrcPort:          5060,
		Tsec:             uint32(tsec),
		Tmsec:            uint32(tmsec),
		Payload:          []byte(SIP_OPTION),
		Vlan:             0,
		ProtoType:        1,
		NodeID:           101,
		NodePW:           "passwd",
	}
	logger.Info("Exporting", "msg", hepMsg)
	err = exporter.Export(ctx, []*hep.Message{hepMsg})
	if err != nil {
		logger.Error("Failed to export HEP message", "err", err)
		return
	}
	exporter.Shutdown(ctx)
}
