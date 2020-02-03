package details

import (
	"strconv"

	"kafmesh-example/internal/definitions/details"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// Processor enriches device details with customer information
type Processor struct{}

// NewProcessor creates a new processor
func NewProcessor() *Processor {
	return &Processor{}
}

// HandleKafmeshDeviceIDDetails handles device details input
func (p *Processor) HandleKafmeshDeviceIDDetails(ctx details.KafmeshDeviceIdEnrichedDetails_ProcessorContext, message *deviceId.Details) error {
	state := ctx.State()
	state.Details = message

	ctx.SaveState(state)

	p.output(ctx, state)

	return nil
}

// HandleKafmeshDeviceIDCustomer handles device to customer input
func (p *Processor) HandleKafmeshDeviceIDCustomer(ctx details.KafmeshDeviceIdEnrichedDetails_ProcessorContext, message *deviceId.Customer) error {
	state := ctx.State()
	state.CustomerId = &wrappers.Int64Value{Value: message.Id}

	ctx.SaveState(state)

	p.output(ctx, state)

	return nil
}

func (p *Processor) output(ctx details.KafmeshDeviceIdEnrichedDetails_ProcessorContext, state *deviceId.EnrichedDetailsState) {
	if state.CustomerId == nil {
		return
	}

	if state.Details == nil {
		return
	}

	customerDetails := ctx.Lookup_CustomerIdDetails(strconv.Itoa(int(state.CustomerId.Value)))
	if customerDetails == nil {
		return
	}

	ctx.Output_DeviceIdEnrichedDetails(ctx.Key(), &deviceId.EnrichedDetails{
		Time:         state.Details.Time,
		Name:         state.Details.Name,
		CustomerId:   state.CustomerId.Value,
		CustomerName: customerDetails.Name,
	})
}
