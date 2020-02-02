package details

import (
	"strconv"

	"kafmesh-example/internal/kafmesh/definitions/details"
	"kafmesh-example/internal/kafmesh/definitions/models/kafmesh/deviceId"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// Processor enriches device details with customer information
type Processor struct{}

// NewProcessor creates a new processor
func NewProcessor() *Processor {
	return &Processor{}
}

// HandleDeviceIdDetails handles device details input
func (p *Processor) HandleDeviceIdDetails(ctx details.KafmeshDeviceIdEnrichedDetails_ProcessorContext, message *deviceId.Details) error {
	state := ctx.State()
	state.Name = &wrappers.StringValue{Value: message.Name}

	ctx.SaveState(state)

	p.Output(ctx, state)

	return nil
}

// HandleDeviceIdCustomer handles device to customer input
func (p *Processor) HandleDeviceIdCustomer(ctx details.KafmeshDeviceIdEnrichedDetails_ProcessorContext, message *deviceId.Customer) error {
	state := ctx.State()
	state.CustomerId = &wrappers.Int64Value{Value: message.Id}

	ctx.SaveState(state)

	p.Output(ctx, state)

	return nil
}

// Output will output the enriched details if they are valid
func (p *Processor) Output(ctx details.KafmeshDeviceIdEnrichedDetails_ProcessorContext, state *deviceId.EnrichedDetailsState) {
	if state.CustomerId == nil {
		return
	}

	if state.Name == nil {
		return
	}

	customerDetails := ctx.Lookup_CustomerIdDetails(strconv.Itoa(int(state.CustomerId.Value)))
	if customerDetails == nil {
		return
	}

	ctx.Output_DeviceIdEnrichedDetails(ctx.Key(), &deviceId.EnrichedDetails{
		Name:         state.Name.Value,
		CustomerId:   state.CustomerId.Value,
		CustomerName: customerDetails.Name,
	})
}
