package heartbeats_test

import (
	"testing"
	"time"

	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/implementation/heartbeats"
	"kafmesh-example/internal/warehouse"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/runner"
	"gotest.tools/assert"
)

func Test_Sink_CollectandFlush(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockRepository(ctrl)

	now := time.Now()
	pnow, err := ptypes.TimestampProto(now)
	assert.NilError(t, err)
	resolvedNow, err := ptypes.Timestamp(pnow)
	assert.NilError(t, err)
	repo.EXPECT().SaveHeartbeats(gomock.Any(), []warehouse.Heartbeat{
		{
			DeviceID:     67,
			Time:         resolvedNow,
			IsHealthy:    true,
			CustomerID:   42,
			CustomerName: "testing customer",
		},
	})

	sink := heartbeats.NewWarehouseSink(repo)

	err = sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedHeartbeat{
		Time:         pnow,
		IsHealthy:    true,
		CustomerId:   42,
		CustomerName: "testing customer",
	})
	assert.NilError(t, err)

	err = sink.Flush()
	assert.NilError(t, err)
}

func Test_Sink_CollectShouldFailWithBadKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockRepository(ctrl)

	sink := heartbeats.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "a67", &deviceId.EnrichedHeartbeat{})
	assert.ErrorContains(t, err, "failed to convert 'a67' to int")
}

func Test_Sink_FlushShouldFailWithNilTimestamp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockRepository(ctrl)

	sink := heartbeats.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedHeartbeat{})
	assert.ErrorContains(t, err, "failed to convert timestamp to time: timestamp: nil Timestamp")
}

func Test_Sink_FlushShouldFailWhenRepoFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockRepository(ctrl)
	repo.EXPECT().SaveHeartbeats(gomock.Any(), gomock.Any()).Return(errors.Errorf("boom"))

	sink := heartbeats.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedHeartbeat{
		Time: ptypes.TimestampNow(),
	})
	assert.NilError(t, err)

	err = sink.Flush()
	assert.ErrorContains(t, err, "failed to save heartbeats to database: boom")
}
