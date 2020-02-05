package heartbeats

import (
	"strconv"

	"kafmesh-example/internal/definitions/heartbeats"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
)

// Processor enriches device details with customer information
type Processor struct{}

// NewProcessor creates a new processor
func NewProcessor() *Processor {
	return &Processor{}
}

// HandleKafmeshDeviceIDHeartbeat handles device heartbeat input
func (p *Processor) HandleKafmeshDeviceIDHeartbeat(ctx heartbeats.KafmeshDeviceIdEnrichedHeartbeat_ProcessorContext, message *deviceId.Heartbeat) error {
	customer := ctx.Join_DeviceIdCustomer()

	if customer == nil {
		return nil
	}

	customerDetails := ctx.Lookup_CustomerIdDetails(strconv.Itoa(int(customer.Id)))
	if customerDetails == nil {
		return nil
	}

	ctx.Output_DeviceIdEnrichedHeartbeat(ctx.Key(), &deviceId.EnrichedHeartbeat{
		Time:         message.Time,
		IsHealthy:    message.IsHealthy,
		CustomerId:   customer.Id,
		CustomerName: customerDetails.Name,
	})

	return nil
}
