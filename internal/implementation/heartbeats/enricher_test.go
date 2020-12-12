package heartbeats_test

import (
	"testing"

	mocks "kafmesh-example/internal/definitions/heartbeats"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/implementation/heartbeats"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"gotest.tools/assert"
)

func Test_Enricher_ShouldNotOutputWithNullCustomerMap(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := mocks.NewMockEnricher_ProcessorContext(ctrl)
	ctx.EXPECT().Key().Return("42")
	ctx.EXPECT().Join_DeviceIDCustomer().Return(nil)

	heartbeat := &deviceId.Heartbeat{
		Time: ptypes.TimestampNow(),
	}

	subject := heartbeats.NewEnricher()
	err := subject.HandleDeviceIDHeartbeat(ctx, heartbeat)
	assert.NilError(t, err)
}

func Test_Enricher_ShouldNotOutputWithNullCustomerDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := mocks.NewMockEnricher_ProcessorContext(ctrl)
	ctx.EXPECT().Key().Return("42")
	ctx.EXPECT().Join_DeviceIDCustomer().Return(&deviceId.Customer{
		Id: 78,
	})
	ctx.EXPECT().Lookup_CustomerIDDetails("78").Return(nil)

	heartbeat := &deviceId.Heartbeat{
		Time: ptypes.TimestampNow(),
	}

	subject := heartbeats.NewEnricher()
	err := subject.HandleDeviceIDHeartbeat(ctx, heartbeat)
	assert.NilError(t, err)
}

func Test_Enricher_ShouldOutput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := mocks.NewMockEnricher_ProcessorContext(ctrl)
	ctx.EXPECT().Key().Return("42")
	ctx.EXPECT().Join_DeviceIDCustomer().Return(&deviceId.Customer{
		Id: 78,
	})
	ctx.EXPECT().Lookup_CustomerIDDetails("78").Return(&customerId.Details{
		Name: "testing customer",
	})

	heartbeat := &deviceId.Heartbeat{
		Time:      ptypes.TimestampNow(),
		IsHealthy: true,
	}
	ctx.EXPECT().Output_DeviceIDEnrichedHeartbeat("42", &deviceId.EnrichedHeartbeat{
		CustomerId:   78,
		CustomerName: "testing customer",
		Time:         heartbeat.Time,
		IsHealthy:    true,
	})

	subject := heartbeats.NewEnricher()
	err := subject.HandleDeviceIDHeartbeat(ctx, heartbeat)
	assert.NilError(t, err)
}
