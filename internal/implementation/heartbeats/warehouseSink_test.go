package heartbeats_test

import (
	"context"
	"testing"
	"time"

	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/implementation/heartbeats"
	"kafmesh-example/internal/warehouse"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/runner"
	"gotest.tools/assert"
)

func Test_Sink_CollectandFlush(t *testing.T) {
	repo := &repository{}

	var savedHeartbeats []warehouse.Heartbeat
	repo.saveHeartbeats = func(ctx context.Context, heartbeats []warehouse.Heartbeat) error {
		savedHeartbeats = heartbeats
		return nil
	}

	sink := heartbeats.NewWarehouseSink(repo)

	now := time.Now()
	pnow, err := ptypes.TimestampProto(now)
	assert.NilError(t, err)

	err = sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedHeartbeat{
		Time:         pnow,
		IsHealthy:    true,
		CustomerId:   42,
		CustomerName: "testing customer",
	})
	assert.NilError(t, err)

	err = sink.Flush()
	assert.NilError(t, err)

	assert.DeepEqual(t, savedHeartbeats, []warehouse.Heartbeat{
		warehouse.Heartbeat{
			DeviceID:     67,
			Time:         now,
			IsHealthy:    true,
			CustomerID:   42,
			CustomerName: "testing customer",
		},
	})
}

func Test_Sink_CollectShouldFailWithBadKey(t *testing.T) {
	repo := &repository{}

	sink := heartbeats.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "a67", &deviceId.EnrichedHeartbeat{})
	assert.ErrorContains(t, err, "failed to convert 'a67' to int")
}

func Test_Sink_FlushShouldFailWithNilTimestamp(t *testing.T) {
	repo := &repository{}
	repo.saveHeartbeats = func(ctx context.Context, details []warehouse.Heartbeat) error {
		return errors.Errorf("boom")
	}

	sink := heartbeats.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedHeartbeat{})
	assert.ErrorContains(t, err, "failed to convert timestamp to time: timestamp: nil Timestamp")
}

func Test_Sink_FlushShouldFailWhenRepoFails(t *testing.T) {
	repo := &repository{}
	repo.saveHeartbeats = func(ctx context.Context, details []warehouse.Heartbeat) error {
		return errors.Errorf("boom")
	}

	sink := heartbeats.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedHeartbeat{
		Time: ptypes.TimestampNow(),
	})
	assert.NilError(t, err)

	err = sink.Flush()
	assert.ErrorContains(t, err, "failed to save heartbeats to database: boom")
}
