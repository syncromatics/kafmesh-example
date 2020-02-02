package details_test

import (
	"testing"

	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/implementation/details"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"gotest.tools/assert"
)

func Test_Processor_ShouldNotOutputWithNullCustomer(t *testing.T) {
	context := &contextMock{}

	p := details.NewProcessor()
	context.state = func() *deviceId.EnrichedDetailsState {
		return &deviceId.EnrichedDetailsState{}
	}

	var savedState *deviceId.EnrichedDetailsState
	context.saveState = func(state *deviceId.EnrichedDetailsState) {
		savedState = state
	}

	context.output_DeviceIdEnrichedDetails = func(string, *deviceId.EnrichedDetails) {
		t.Fatal("should not output")
	}

	err := p.HandleDeviceIdDetails(context, &deviceId.Details{
		Name: "testing",
	})
	assert.NilError(t, err)

	assert.Assert(t, proto.Equal(savedState, &deviceId.EnrichedDetailsState{
		Name: &wrappers.StringValue{Value: "testing"},
	}))
}

func Test_Processor_ShouldNotOutputWithNullDetails(t *testing.T) {
	context := &contextMock{}

	p := details.NewProcessor()
	context.state = func() *deviceId.EnrichedDetailsState {
		return &deviceId.EnrichedDetailsState{}
	}

	var savedState *deviceId.EnrichedDetailsState
	context.saveState = func(state *deviceId.EnrichedDetailsState) {
		savedState = state
	}

	context.output_DeviceIdEnrichedDetails = func(string, *deviceId.EnrichedDetails) {
		t.Fatal("should not output")
	}

	err := p.HandleDeviceIdCustomer(context, &deviceId.Customer{
		Id: 42,
	})
	assert.NilError(t, err)

	assert.Assert(t, proto.Equal(savedState, &deviceId.EnrichedDetailsState{
		CustomerId: &wrappers.Int64Value{Value: 42},
	}))
}

func Test_Processor_ShouldNotOutputWithNullCustomerDetails(t *testing.T) {
	context := &contextMock{}

	p := details.NewProcessor()
	context.state = func() *deviceId.EnrichedDetailsState {
		return &deviceId.EnrichedDetailsState{
			Name: &wrappers.StringValue{Value: "testing"},
		}
	}

	var savedState *deviceId.EnrichedDetailsState
	context.saveState = func(state *deviceId.EnrichedDetailsState) {
		savedState = state
	}

	context.output_DeviceIdEnrichedDetails = func(string, *deviceId.EnrichedDetails) {
		t.Fatal("should not output")
	}

	context.lookup_CustomerIdDetails = func(key string) *customerId.Details {
		assert.Equal(t, key, "42")
		return nil
	}

	err := p.HandleDeviceIdCustomer(context, &deviceId.Customer{
		Id: 42,
	})
	assert.NilError(t, err)

	assert.Assert(t, proto.Equal(savedState, &deviceId.EnrichedDetailsState{
		CustomerId: &wrappers.Int64Value{Value: 42},
		Name:       &wrappers.StringValue{Value: "testing"},
	}))
}

func Test_Processor_ShouldOutput(t *testing.T) {
	context := &contextMock{}

	p := details.NewProcessor()
	context.state = func() *deviceId.EnrichedDetailsState {
		return &deviceId.EnrichedDetailsState{
			Name: &wrappers.StringValue{Value: "testing"},
		}
	}

	var savedState *deviceId.EnrichedDetailsState
	context.saveState = func(state *deviceId.EnrichedDetailsState) {
		savedState = state
	}

	var output *deviceId.EnrichedDetails
	var outputKey string
	context.output_DeviceIdEnrichedDetails = func(key string, message *deviceId.EnrichedDetails) {
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

	err := p.HandleDeviceIdCustomer(context, &deviceId.Customer{
		Id: 42,
	})
	assert.NilError(t, err)

	assert.Assert(t, proto.Equal(savedState, &deviceId.EnrichedDetailsState{
		CustomerId: &wrappers.Int64Value{Value: 42},
		Name:       &wrappers.StringValue{Value: "testing"},
	}))

	assert.Equal(t, outputKey, "423")
	assert.Assert(t, proto.Equal(output, &deviceId.EnrichedDetails{
		CustomerId:   42,
		CustomerName: "testing customer",
		Name:         "testing",
	}))
}
