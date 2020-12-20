package details_test

import (
	"testing"

	mocks "kafmesh-example/internal/definitions/details"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/implementation/details"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/wrappers"
	"gotest.tools/assert"
)

func Test_Processor_ShouldNotOutputWithNullCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	context := mocks.NewMockEnricher_ProcessorContext(ctrl)

	p := details.NewEnricherProcessor()
	context.EXPECT().State().Return(&deviceId.EnrichedDetailsState{})

	context.EXPECT().SaveState(&deviceId.EnrichedDetailsState{
		Details: &deviceId.Details{
			Name: "testing",
		},
	})

	err := p.HandleDeviceIDDetails(context, &deviceId.Details{
		Name: "testing",
	})
	assert.NilError(t, err)
}

func Test_Processor_ShouldNotOutputWithNullDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	context := mocks.NewMockEnricher_ProcessorContext(ctrl)
	context.EXPECT().State().Return(&deviceId.EnrichedDetailsState{})

	p := details.NewEnricherProcessor()

	context.EXPECT().SaveState(&deviceId.EnrichedDetailsState{
		CustomerId: &wrappers.Int64Value{Value: 42},
	})

	err := p.HandleDeviceIDCustomer(context, &deviceId.Customer{
		Id: 42,
	})
	assert.NilError(t, err)
}

func Test_Processor_ShouldNotOutputWithNullCustomerDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := details.NewEnricherProcessor()
	context := mocks.NewMockEnricher_ProcessorContext(ctrl)
	context.EXPECT().State().Return(&deviceId.EnrichedDetailsState{
		Details: &deviceId.Details{
			Name: "testing",
		},
	})

	context.EXPECT().SaveState(&deviceId.EnrichedDetailsState{
		CustomerId: &wrappers.Int64Value{Value: 42},
		Details: &deviceId.Details{
			Name: "testing",
		},
	})

	context.EXPECT().Lookup_CustomerIDDetails("42").Return(nil)

	err := p.HandleDeviceIDCustomer(context, &deviceId.Customer{
		Id: 42,
	})
	assert.NilError(t, err)
}

func Test_Processor_ShouldOutput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	p := details.NewEnricherProcessor()
	context := mocks.NewMockEnricher_ProcessorContext(ctrl)
	context.EXPECT().State().Return(&deviceId.EnrichedDetailsState{
		Details: &deviceId.Details{
			Name: "testing",
		},
	})

	context.EXPECT().SaveState(&deviceId.EnrichedDetailsState{
		CustomerId: &wrappers.Int64Value{Value: 42},
		Details: &deviceId.Details{
			Name: "testing",
		},
	})

	context.EXPECT().Output_DeviceIDEnrichedDetails("423", &deviceId.EnrichedDetails{
		CustomerId:   42,
		CustomerName: "testing customer",
		Name:         "testing",
	})

	context.EXPECT().Lookup_CustomerIDDetails("42").Return(&customerId.Details{
		Name: "testing customer",
	})

	context.EXPECT().Key().Return("423")

	err := p.HandleDeviceIDCustomer(context, &deviceId.Customer{
		Id: 42,
	})
	assert.NilError(t, err)
}
