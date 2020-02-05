package details_test

import (
	"context"

	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/warehouse"
)

type contextMock struct {
	key                            string
	lookup_CustomerIdDetails       func(key string) *customerId.Details
	output_DeviceIdEnrichedDetails func(key string, message *deviceId.EnrichedDetails)
	saveState                      func(state *deviceId.EnrichedDetailsState)
	state                          func() *deviceId.EnrichedDetailsState
}

func (c *contextMock) Key() string {
	return c.key
}

func (c *contextMock) Lookup_CustomerIdDetails(key string) *customerId.Details {
	return c.lookup_CustomerIdDetails(key)
}

func (c *contextMock) Output_DeviceIdEnrichedDetails(key string, message *deviceId.EnrichedDetails) {
	c.output_DeviceIdEnrichedDetails(key, message)
}

func (c *contextMock) SaveState(state *deviceId.EnrichedDetailsState) {
	c.saveState(state)
}

func (c *contextMock) State() *deviceId.EnrichedDetailsState {
	return c.state()
}

type repository struct {
	saveDetails func(ctx context.Context, details []warehouse.Details) error
}

func (r *repository) SaveDetails(ctx context.Context, details []warehouse.Details) error {
	return r.saveDetails(ctx, details)
}
