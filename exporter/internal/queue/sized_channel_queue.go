package queue // import "go.voiplens.io/hep/exporter/internal/queue"

import (
	"context"
	"sync/atomic"
)

type sizedChannelQueue struct {
	size *atomic.Int64
	cap  int64
	ch   chan queueItem
}

type queueItem struct {
	payload []byte
	ctx     context.Context
}

var _ Queue = (*sizedChannelQueue)(nil)

func NewSizedChannelQueue(capacity int64) *sizedChannelQueue {
	size := &atomic.Int64{}
	size.Store(0)

	ch := make(chan queueItem, capacity)

	return &sizedChannelQueue{
		size: size,
		cap:  capacity,
		ch:   ch,
	}
}

// Pop implements Queue.
func (c *sizedChannelQueue) Pop(consumeFunc func(ctx context.Context, payload []byte) error) bool {
	item, ok := <-c.ch
	if !ok {
		return false
	}
	if c.size.Add(-1) < 0 {
		c.size.Store(0)
	}

	err := consumeFunc(item.ctx, item.payload)
	return err == nil
}

// Push implements Push.
func (c *sizedChannelQueue) Push(ctx context.Context, payload []byte) error {
	if c.size.Add(1) > c.cap {
		c.size.Add(-1)
		return ErrQueueIsFull
	}
	c.ch <- queueItem{payload: payload, ctx: ctx}
	return nil
}

// Run implements Queue.
func (c *sizedChannelQueue) Start(ctx context.Context) error {
	return nil
}

// Shutdown implements Queue.
func (c *sizedChannelQueue) Shutdown(ctx context.Context) error {
	close(c.ch)
	return nil
}

// Capacity implements Queue.
func (c *sizedChannelQueue) Capacity() int {
	return int(c.cap)
}

// Size implements Queue.
func (c *sizedChannelQueue) Size() int {
	return int(c.size.Load())
}
