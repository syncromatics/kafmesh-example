package heartbeats_test

import (
	"testing"

	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/implementation/heartbeats"

	"github.com/golang/protobuf/ptypes"
	"gotest.tools/assert"
)

func Test_Processor_ShouldNotOutputWithNullCustomerMap(t *testing.T) {
	context := &contextMock{}

	p := heartbeats.NewProcessor()

	context.output_DeviceIdEnrichedHeartbeat = func(string, *deviceId.EnrichedHeartbeat) {
		t.Fatal("should not output")
	}

	context.join_DeviceIdCustomer = func() *deviceId.Customer {
		return nil
	}

	heartbeat := &deviceId.Heartbeat{
		Time: ptypes.TimestampNow(),
	}

	err := p.HandleKafmeshDeviceIDHeartbeat(context, heartbeat)
	assert.NilError(t, err)
}

func Test_Processor_ShouldNotOutputWithNullCustomerDetails(t *testing.T) {
	context := &contextMock{}

	p := heartbeats.NewProcessor()

	context.lookup_CustomerIdDetails = func(key string) *customerId.Details {
		assert.Equal(t, key, "78")
		return nil
	}

	context.join_DeviceIdCustomer = func() *deviceId.Customer {
		return &deviceId.Customer{Id: 78}
	}

	heartbeat := &deviceId.Heartbeat{
		Time: ptypes.TimestampNow(),
	}

	err := p.HandleKafmeshDeviceIDHeartbeat(context, heartbeat)
	assert.NilError(t, err)
}

func Test_Processor_ShouldOutput(t *testing.T) {
	context := &contextMock{}

	p := heartbeats.NewProcessor()

	var output *deviceId.EnrichedHeartbeat
	var outputKey string
	context.output_DeviceIdEnrichedHeartbeat = func(key string, message *deviceId.EnrichedHeartbeat) {
		outputKey = key
		output = message
	}

	context.lookup_CustomerIdDetails = func(key string) *customerId.Details {
		assert.Equal(t, key, "42")
		return &customerId.Details{
			Name: "testing customer",
		}
	}

	context.key = "423"

	context.lookup_CustomerIdDetails = func(key string) *customerId.Details {
		assert.Equal(t, key, "78")
		return &customerId.Details{
			Name: "testing customer",
		}
	}

	context.join_DeviceIdCustomer = func() *deviceId.Customer {
		return &deviceId.Customer{Id: 78}
	}

	heartbeat := &deviceId.Heartbeat{
		Time:      ptypes.TimestampNow(),
		IsHealthy: true,
	}

	err := p.HandleKafmeshDeviceIDHeartbeat(context, heartbeat)
	assert.NilError(t, err)

	assert.Equal(t, outputKey, "423")
	assert.DeepEqual(t, output, &deviceId.EnrichedHeartbeat{
		CustomerId:   78,
		CustomerName: "testing customer",
		Time:         heartbeat.Time,
		IsHealthy:    true,
	})
}
