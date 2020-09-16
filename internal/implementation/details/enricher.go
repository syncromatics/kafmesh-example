package details

import (
	"strconv"

	"kafmesh-example/internal/definitions/details"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// Enricher enriches device details with customer information
type Enricher struct{}

// NewEnricher creates a new processor
func NewEnricher() *Enricher {
	return &Enricher{}
}

// HandleDeviceIDDetails handles device details input
func (p *Enricher) HandleDeviceIDDetails(ctx details.Enricher_ProcessorContext, message *deviceId.Details) error {
	state := ctx.State()
	state.Details = message

	ctx.SaveState(state)

	p.output(ctx, state)

	return nil
}

// HandleKafmeshDeviceIDCustomer handles device to customer input
func (p *Enricher) HandleKafmeshDeviceIDCustomer(ctx details.Enricher_ProcessorContext, message *deviceId.Customer) error {
	state := ctx.State()
	state.CustomerId = &wrappers.Int64Value{Value: message.Id}

	ctx.SaveState(state)

	p.output(ctx, state)

	return nil
}

func (p *Enricher) output(ctx details.Enricher_ProcessorContext, state *deviceId.EnrichedDetailsState) {
	if state.CustomerId == nil {
		return
	}

	if state.Details == nil {
		return
	}

	customerID := strconv.Itoa(int(state.CustomerId.Value))
	customerDetails := ctx.Lookup_CustomerIDDetails(customerID)
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
