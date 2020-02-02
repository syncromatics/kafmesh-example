package details

import (
	"kafmesh-example/internal/kafmesh/definitions/models/kafmesh/deviceId"

	"github.com/syncromatics/kafmesh/pkg/runner"
)

// WarehouseSink sinks enriched device details to the warehouse
type WarehouseSink struct {
}

// NewWarehouseSink creates a new warehouse sink
func NewWarehouseSink() *WarehouseSink {
	return &WarehouseSink{}
}

// Collect collects enriched details till flush
func (s *WarehouseSink) Collect(ctx runner.MessageContext, key string, msg *deviceId.EnrichedDetails) error {
	return nil
}

// Flush saves the buffered details to the database
func (s *WarehouseSink) Flush() error {
	return nil
}
