package heartbeats

import (
	"kafmesh-example/internal/definitions/heartbeats"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"strconv"
)

var _ heartbeats.Enricher_Processor = &Enricher{}

// Enricher enrichers device details with customer information
type Enricher struct{}

// NewEnricher creates a new enricher
func NewEnricher() *Enricher {
	return &Enricher{}
}

// HandleDeviceIDHeartbeat handles device heartbeat input
func (e *Enricher) HandleDeviceIDHeartbeat(ctx heartbeats.Enricher_ProcessorContext, message *deviceId.Heartbeat) error {
	key := ctx.Key()
	customer := ctx.Join_DeviceIDCustomer()
	if customer == nil {
		return nil
	}

	customerID := strconv.FormatInt(customer.Id, 10)
	customerDetails := ctx.Lookup_CustomerIDDetails(customerID)
	if customerDetails == nil {
		return nil
	}

	output := &deviceId.EnrichedHeartbeat{
		Time:         message.Time,
		IsHealthy:    message.IsHealthy,
		CustomerId:   customer.Id,
		CustomerName: customerDetails.Name,
	}
	ctx.Output_DeviceIDEnrichedHeartbeat(key, output)

	return nil
}
