package hep

import "context"

// Exporter is the interface that must be implemented by all HEP exporters.
//
// Once created an Exporter follows the following typical lifecycle:
//
//  1. Start: The exporter's Start method is called.
//  2. Running: The exporter is up and running. And process messages via the Export method.
//  3. Shutdown: The exporter's Shutdown method is called and the lifecycle is complete.
//
// Once the lifecycle is complete it may be repeated by starting the exporter again.
type Exporter interface {

	// Start tells the exporter to start.
	// If an error is returned by Start() then the exporter startup will be aborted.
	// This function can be used to prepare for exporting by connecting to the endpoint.
	Start(ctx context.Context) error

	// Export transmits HEP messages to an endpoint.
	//
	// The deadline or cancellation of the passed context must be honoured. An
	// appropriate error should be returned in these situations.
	//
	// The exporter must implement all retry logic in this function.
	//
	// Implementations must not retain the records slice.
	Export(ctx context.Context, messages []*Message) error

	// Shutdown is invoked during the exporter shutdown. After Shutdown() is called, the exporter
	// should not accept data anymore.
	//
	// This method must be safe to call:
	//   - without Start() having been called
	//   - if the exporter is in a shutdown state already
	//
	// If there are any background operations running by the component they must be aborted before
	// this function returns. Remember that if you started any long-running background operations from
	// the Start() method, those operations must be also cancelled. If there are any buffers in the
	// exporter, they should be cleared and the data sent immediately to the endpoint.
	Shutdown(ctx context.Context) error
}
