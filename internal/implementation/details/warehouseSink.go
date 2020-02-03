package details

import (
	"context"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"

	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/warehouse"

	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/runner"
)

// Repository is the repository to the warehouse.
type Repository interface {
	SaveDetails(ctx context.Context, details []warehouse.Details) error
}

// WarehouseSink sinks enriched device details to the warehouse
type WarehouseSink struct {
	repo   Repository
	buffer []warehouse.Details
}

// NewWarehouseSink creates a new warehouse sink
func NewWarehouseSink(repository Repository) *WarehouseSink {
	return &WarehouseSink{
		repo:   repository,
		buffer: []warehouse.Details{},
	}
}

// Collect collects enriched details till flush
func (s *WarehouseSink) Collect(ctx runner.MessageContext, key string, msg *deviceId.EnrichedDetails) error {
	id, err := strconv.Atoi(key)
	if err != nil {
		return errors.Wrapf(err, "failed to convert '%s' to int", key)
	}

	t, err := ptypes.Timestamp(msg.Time)
	if err != nil {
		return errors.Wrap(err, "failed to convert timestamp to time")
	}

	s.buffer = append(s.buffer, warehouse.Details{
		DeviceID:     int64(id),
		Time:         t,
		Name:         msg.Name,
		CustomerID:   msg.CustomerId,
		CustomerName: msg.CustomerName,
	})

	return nil
}

// Flush saves the buffered details to the database
func (s *WarehouseSink) Flush() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := s.repo.SaveDetails(ctx, s.buffer)
	if err != nil {
		return errors.Wrap(err, "failed to save details to database")
	}

	s.buffer = s.buffer[:0]

	return nil
}
