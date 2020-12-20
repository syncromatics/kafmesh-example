package heartbeats

import (
	"strconv"

	"kafmesh-example/internal/definitions/heartbeats"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
)

var _ heartbeats.HeartbeatEnricher_Processor = &Processor{}

// Processor enriches device details with customer information
type Processor struct{}

// NewProcessor creates a new processor
func NewProcessor() *Processor {
	return &Processor{}
}

// HandleDeviceIDHeartbeat handles device heartbeat input
func (p *Processor) HandleDeviceIDHeartbeat(ctx heartbeats.HeartbeatEnricher_ProcessorContext, message *deviceId.Heartbeat) error {
	customer := ctx.Join_DeviceIDCustomer()

	if customer == nil {
		return nil
	}

	customerDetails := ctx.Lookup_CustomerIDDetails(strconv.Itoa(int(customer.Id)))
	if customerDetails == nil {
		return nil
	}

	ctx.Output_DeviceIDEnrichedHeartbeat(ctx.Key(), &deviceId.EnrichedHeartbeat{
		Time:         message.Time,
		IsHealthy:    message.IsHealthy,
		CustomerId:   customer.Id,
		CustomerName: customerDetails.Name,
	})

	return nil
}
