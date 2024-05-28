package queue // import "go.voiplens.io/hep/exporter/internal/queue"

import (
	"context"
	"errors"
)

var (
	// ErrQueueIsFull is the error returned when an item is offered to the Queue and the queue is full.
	ErrQueueIsFull = errors.New("sending queue is full")
)

type Queue interface {
	// Run is called to start the worker
	Start(ctx context.Context) error

	// Shutdown is called if stop all worker
	Shutdown(ctx context.Context) error

	// Offer inserts the specified element into this queue if it is possible to do so immediately
	// without violating capacity restrictions. If success returns no error.
	// It returns ErrQueueIsFull if no space is currently available.
	Push(ctx context.Context, payload []byte) error

	// Pop applies the provided function on the head of queue.
	// The call blocks until there is an item available or the queue is stopped.
	// The function returns true when an item is consumed or false if the queue is stopped.
	Pop(func(ctx context.Context, payload []byte) error) bool
	// Size returns the current Size of the queue
	Size() int
	// Capacity returns the capacity of the queue.
	Capacity() int
}
