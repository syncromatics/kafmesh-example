package heartbeats

import (
	"context"
	"strconv"
	"time"

	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/warehouse"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/runner"
)

// Repository saves heartbeats to the warehouse
type Repository interface {
	SaveHeartbeats(context.Context, []warehouse.Heartbeat) error
}

// WarehouseSink sinks enriched device details to the warehouse
type WarehouseSink struct {
	repository Repository

	buffer []warehouse.Heartbeat
}

// NewWarehouseSink creates a new warehouse sink
func NewWarehouseSink(repository Repository) *WarehouseSink {
	return &WarehouseSink{
		repository: repository,
	}
}

// Collect collects enriched details till flush
func (s *WarehouseSink) Collect(ctx runner.MessageContext, key string, msg *deviceId.EnrichedHeartbeat) error {
	id, err := strconv.Atoi(key)
	if err != nil {
		return errors.Wrapf(err, "failed to convert '%s' to int", key)
	}

	t, err := ptypes.Timestamp(msg.Time)
	if err != nil {
		return errors.Wrapf(err, "failed to convert timestamp to time")
	}

	s.buffer = append(s.buffer, warehouse.Heartbeat{
		DeviceID:     int64(id),
		Time:         t,
		IsHealthy:    msg.IsHealthy,
		CustomerID:   msg.CustomerId,
		CustomerName: msg.CustomerName,
	})

	return nil
}

// Flush saves the buffered details to the database
func (s *WarehouseSink) Flush() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := s.repository.SaveHeartbeats(ctx, s.buffer)
	if err != nil {
		return errors.Wrap(err, "failed to save heartbeats to database")
	}

	s.buffer = s.buffer[:0]

	return nil
}
