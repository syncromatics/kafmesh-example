package details

import (
	"strconv"

	"kafmesh-example/internal/definitions/details"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"

	"github.com/golang/protobuf/ptypes/wrappers"
)

var _ details.Enricher_Processor = &EnricherProcessor{}

// EnricherProcessor enriches device details with customer information
type EnricherProcessor struct{}

// NewEnricherProcessor creates a new processor
func NewEnricherProcessor() *EnricherProcessor {
	return &EnricherProcessor{}
}

// HandleDeviceIDDetails handles device details input
func (p *EnricherProcessor) HandleDeviceIDDetails(ctx details.Enricher_ProcessorContext, message *deviceId.Details) error {
	state := ctx.State()
	state.Details = message

	ctx.SaveState(state)

	p.output(ctx, state)

	return nil
}

// HandleDeviceIDCustomer handles device to customer input
func (p *EnricherProcessor) HandleDeviceIDCustomer(ctx details.Enricher_ProcessorContext, message *deviceId.Customer) error {
	state := ctx.State()
	state.CustomerId = &wrappers.Int64Value{Value: message.Id}

	ctx.SaveState(state)

	p.output(ctx, state)

	return nil
}

func (p *EnricherProcessor) output(ctx details.Enricher_ProcessorContext, state *deviceId.EnrichedDetailsState) {
	if state.CustomerId == nil {
		return
	}

	if state.Details == nil {
		return
	}

	customerDetails := ctx.Lookup_CustomerIDDetails(strconv.Itoa(int(state.CustomerId.Value)))
	if customerDetails == nil {
		return
	}

	ctx.Output_DeviceIDEnrichedDetails(ctx.Key(), &deviceId.EnrichedDetails{
		Time:         state.Details.Time,
		Name:         state.Details.Name,
		CustomerId:   state.CustomerId.Value,
		CustomerName: customerDetails.Name,
	})
}
