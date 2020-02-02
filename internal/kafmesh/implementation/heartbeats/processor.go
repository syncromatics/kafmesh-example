package heartbeats

import (
	"strconv"

	"kafmesh-example/internal/kafmesh/definitions/heartbeats"
	"kafmesh-example/internal/kafmesh/definitions/models/kafmesh/deviceId"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// Processor enriches device details with customer information
type Processor struct{}

// NewProcessor creates a new processor
func NewProcessor() *Processor {
	return &Processor{}
}

// HandleDeviceIdHeartbeat handles device heartbeat input
func (p *Processor) HandleDeviceIdHeartbeat(ctx heartbeats.KafmeshDeviceIdEnrichedHeartbeat_ProcessorContext, message *deviceId.Heartbeat) error {
	state := ctx.State()
	state.Heartbeat = message

	ctx.SaveState(state)

	p.Output(ctx, state)

	return nil
}

// HandleDeviceIdCustomer handles device to customer input
func (p *Processor) HandleDeviceIdCustomer(ctx heartbeats.KafmeshDeviceIdEnrichedHeartbeat_ProcessorContext, message *deviceId.Customer) error {
	state := ctx.State()
	state.CustomerId = &wrappers.Int64Value{Value: message.Id}

	ctx.SaveState(state)

	p.Output(ctx, state)

	return nil
}

// Output will output the enriched details if they are valid
func (p *Processor) Output(ctx heartbeats.KafmeshDeviceIdEnrichedHeartbeat_ProcessorContext, state *deviceId.EnrichedHeartbeatState) {
	if state.CustomerId == nil {
		return
	}

	if state.Heartbeat == nil {
		return
	}

	customerDetails := ctx.Lookup_CustomerIdDetails(strconv.Itoa(int(state.CustomerId.Value)))
	if customerDetails == nil {
		return
	}

	ctx.Output_DeviceIdEnrichedHeartbeat(ctx.Key(), &deviceId.EnrichedHeartbeat{
		Time:         state.Heartbeat.Time,
		IsHealthy:    state.Heartbeat.IsHealthy,
		CustomerId:   state.CustomerId.Value,
		CustomerName: customerDetails.Name,
	})
}
