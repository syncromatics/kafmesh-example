package details_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"

	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/implementation/details"
	"kafmesh-example/internal/warehouse"

	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/runner"
	"gotest.tools/assert"
)

func Test_Sink_CollectandFlush(t *testing.T) {
	repo := &repository{}

	var savedDetails []warehouse.Details
	repo.saveDetails = func(ctx context.Context, details []warehouse.Details) error {
		savedDetails = details
		return nil
	}

	sink := details.NewWarehouseSink(repo)

	now := time.Now()
	tNow, _ := ptypes.TimestampProto(now)

	err := sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedDetails{
		Name:         "test",
		Time:         tNow,
		CustomerId:   42,
		CustomerName: "testing customer",
	})
	assert.NilError(t, err)

	err = sink.Flush()
	assert.NilError(t, err)

	assert.DeepEqual(t, savedDetails, []warehouse.Details{
		warehouse.Details{
			DeviceID:     67,
			Time:         now,
			Name:         "test",
			CustomerID:   42,
			CustomerName: "testing customer",
		},
	})
}

func Test_Sink_CollectShouldFailWithBadKey(t *testing.T) {
	repo := &repository{}

	sink := details.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "a67", &deviceId.EnrichedDetails{})
	assert.ErrorContains(t, err, "failed to convert 'a67' to int")
}

func Test_Sink_FlushShouldFailWhenRepoFails(t *testing.T) {
	repo := &repository{}
	repo.saveDetails = func(ctx context.Context, details []warehouse.Details) error {
		return errors.Errorf("boom")
	}

	sink := details.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedDetails{
		Time: ptypes.TimestampNow(),
	})
	assert.NilError(t, err)

	err = sink.Flush()
	assert.ErrorContains(t, err, "failed to save details to database: boom")
}

func Test_Sink_CollectShouldFailWhenTimestampConvertFails(t *testing.T) {
	repo := &repository{}

	sink := details.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedDetails{})
	assert.ErrorContains(t, err, "failed to convert timestamp to time: timestamp: nil Timestamp")
}
