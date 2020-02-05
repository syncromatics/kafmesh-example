package heartbeats_test

import (
	"context"

	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/warehouse"
)

type contextMock struct {
	key                              string
	lookup_CustomerIdDetails         func(key string) *customerId.Details
	join_DeviceIdCustomer            func() *deviceId.Customer
	output_DeviceIdEnrichedHeartbeat func(key string, message *deviceId.EnrichedHeartbeat)
	saveState                        func(state *deviceId.EnrichedHeartbeatState)
}

func (c *contextMock) Key() string {
	return c.key
}

func (c *contextMock) Lookup_CustomerIdDetails(key string) *customerId.Details {
	return c.lookup_CustomerIdDetails(key)
}

func (c *contextMock) Join_DeviceIdCustomer() *deviceId.Customer {
	return c.join_DeviceIdCustomer()
}

func (c *contextMock) Output_DeviceIdEnrichedHeartbeat(key string, message *deviceId.EnrichedHeartbeat) {
	c.output_DeviceIdEnrichedHeartbeat(key, message)
}

type repository struct {
	saveHeartbeats func(ctx context.Context, heartbeats []warehouse.Heartbeat) error
}

func (r *repository) SaveHeartbeats(ctx context.Context, heartbeats []warehouse.Heartbeat) error {
	return r.saveHeartbeats(ctx, heartbeats)
}
