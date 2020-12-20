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

func Test_Processor_ShouldNotOutputWithNullCustomerMap(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := mocks.NewMockHeartbeatEnricher_ProcessorContext(ctrl)

	p := heartbeats.NewProcessor()

	context.EXPECT().Join_DeviceIDCustomer().Return(nil)

	heartbeat := &deviceId.Heartbeat{
		Time: ptypes.TimestampNow(),
	}

	err := p.HandleDeviceIDHeartbeat(context, heartbeat)
	assert.NilError(t, err)
}

func Test_Processor_ShouldNotOutputWithNullCustomerDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := mocks.NewMockHeartbeatEnricher_ProcessorContext(ctrl)

	p := heartbeats.NewProcessor()

	context.EXPECT().Lookup_CustomerIDDetails("78").Return(nil)
	context.EXPECT().Join_DeviceIDCustomer().Return(&deviceId.Customer{
		Id: 78,
	})

	heartbeat := &deviceId.Heartbeat{
		Time: ptypes.TimestampNow(),
	}

	err := p.HandleDeviceIDHeartbeat(context, heartbeat)
	assert.NilError(t, err)
}

func Test_Processor_ShouldOutput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	context := mocks.NewMockHeartbeatEnricher_ProcessorContext(ctrl)

	p := heartbeats.NewProcessor()

	heartbeat := &deviceId.Heartbeat{
		Time:      ptypes.TimestampNow(),
		IsHealthy: true,
	}
	context.EXPECT().Key().Return("423")
	context.EXPECT().Output_DeviceIDEnrichedHeartbeat("423", &deviceId.EnrichedHeartbeat{
		CustomerId:   78,
		CustomerName: "testing customer",
		Time:         heartbeat.Time,
		IsHealthy:    true,
	})

	context.EXPECT().Lookup_CustomerIDDetails("78").Return(&customerId.Details{
		Name: "testing customer",
	})
	context.EXPECT().Join_DeviceIDCustomer().Return(&deviceId.Customer{
		Id: 78,
	})

	err := p.HandleDeviceIDHeartbeat(context, heartbeat)
	assert.NilError(t, err)
}
