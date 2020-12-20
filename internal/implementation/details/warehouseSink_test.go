package details_test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"

	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/implementation/details"
	"kafmesh-example/internal/warehouse"

	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/runner"
	"gotest.tools/assert"
)

func Test_Sink_CollectandFlush(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockRepository(ctrl)

	now := time.Now()
	tNow, _ := ptypes.TimestampProto(now)
	resolvedNow, err := ptypes.Timestamp(tNow)
	assert.NilError(t, err)
	repo.EXPECT().SaveDetails(gomock.Any(), []warehouse.Details{
		{
			DeviceID:     67,
			Time:         resolvedNow,
			Name:         "test",
			CustomerID:   42,
			CustomerName: "testing customer",
		},
	})

	sink := details.NewWarehouseSink(repo)
	err = sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedDetails{
		Name:         "test",
		Time:         tNow,
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

	sink := details.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "a67", &deviceId.EnrichedDetails{})
	assert.ErrorContains(t, err, "failed to convert 'a67' to int")
}

func Test_Sink_FlushShouldFailWhenRepoFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockRepository(ctrl)
	repo.EXPECT().SaveDetails(gomock.Any(), gomock.Any()).Return(errors.Errorf("boom"))

	sink := details.NewWarehouseSink(repo)
	err := sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedDetails{
		Time: ptypes.TimestampNow(),
	})
	assert.NilError(t, err)

	err = sink.Flush()
	assert.ErrorContains(t, err, "failed to save details to database: boom")
}

func Test_Sink_CollectShouldFailWhenTimestampConvertFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := NewMockRepository(ctrl)

	sink := details.NewWarehouseSink(repo)

	err := sink.Collect(runner.MessageContext{}, "67", &deviceId.EnrichedDetails{})
	assert.ErrorContains(t, err, "failed to convert timestamp to time: timestamp: nil Timestamp")
}
