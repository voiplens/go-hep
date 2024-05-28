package binaryexporter // import "go.voiplens.io/hep/exporter/binaryexporter"

import (
	"bufio"
	"context"
	"crypto/tls"
	"errors"
	"log/slog"
	"net"
	"sync"

	"go.voiplens.io/hep"
	"go.voiplens.io/hep/exporter/internal/queue"
)

var _ hep.Exporter = (*hepexporter)(nil)

type hepexporter struct {
	config *Config
	logger *slog.Logger
	queue  queue.Queue

	processWG sync.WaitGroup
	conn      net.Conn
	writer    *bufio.Writer
	encoder   hep.Encoder
}

func NewBinaryExporter(config *Config, logger *slog.Logger, encoder hep.Encoder) (*hepexporter, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	queue := queue.NewSizedChannelQueue(config.QueueSize)

	return &hepexporter{
		config:  config,
		logger:  logger,
		encoder: encoder,
		queue:   queue,
	}, nil
}

// Connect implements hep.Exporter.
func (h *hepexporter) Start(ctx context.Context) (err error) {
	cfg := h.config
	dialer := net.Dialer{
		Timeout:   cfg.Timeout,
		KeepAlive: cfg.Keepalive,
	}
	switch cfg.Network {
	case NetworkTCP,
		NetworkUDP:
		if h.conn, err = dialer.DialContext(ctx, string(cfg.Network), cfg.Endpoint); err != nil {
			return err
		}
	case NetworkTLS:
		tlsConfig := tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify}
		tlsDialer := tls.Dialer{NetDialer: &dialer, Config: &tlsConfig}
		if h.conn, err = tlsDialer.DialContext(ctx, "tcp", cfg.Endpoint); err != nil {
			return err
		}
	}

	h.writer = bufio.NewWriterSize(h.conn, 8192)

	if err := h.queue.Start(ctx); err != nil {
		return err
	}
	h.startQueueProcessing()

	return nil
}

func (h *hepexporter) startQueueProcessing() {
	h.processWG.Add(1)
	go func() {
		defer h.processWG.Done()
		for {
			if !h.queue.Pop(h.send) {
				return
			}
		}
	}()
}

func (h *hepexporter) send(ctx context.Context, payload []byte) error {
	_, err := h.writer.Write(payload)
	if err != nil {
		return err
	}
	return h.writer.Flush()
}

// Export implements hep.Exporter.
func (h *hepexporter) Export(ctx context.Context, records []*hep.Message) error {
	errs := []error{}
	for _, record := range records {
		data, err := h.encoder.Encode(record)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if err = h.queue.Push(ctx, data); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

// Shutdown implements hep.Exporter.
func (h *hepexporter) Shutdown(ctx context.Context) error {
	errs := []error{}
	errs = append(errs, h.queue.Shutdown(ctx))
	h.processWG.Wait()
	errs = append(errs, h.writer.Flush())
	errs = append(errs, h.conn.Close())
	return errors.Join(errs...)
}
