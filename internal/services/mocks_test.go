package services_test

import (
	"context"

	"kafmesh-example/internal/definitions/assignments"
	"kafmesh-example/internal/definitions/details"
	"kafmesh-example/internal/definitions/heartbeats"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/warehouse"
)

type customerEmitter struct {
	emit     func(message assignments.DeviceIdCustomer_Emitter_Message) error
	emitBulk func(ctx context.Context, messages []assignments.DeviceIdCustomer_Emitter_Message) error
}

func (e *customerEmitter) Emit(message assignments.DeviceIdCustomer_Emitter_Message) error {
	return e.emit(message)
}

func (e *customerEmitter) EmitBulk(ctx context.Context, messages []assignments.DeviceIdCustomer_Emitter_Message) error {
	return e.emitBulk(ctx, messages)
}

type customerView struct {
	keys func() []string
	get  func(key string) (*deviceId.Customer, error)
}

func (v *customerView) Keys() []string {
	return v.keys()
}

func (v *customerView) Get(key string) (*deviceId.Customer, error) {
	return v.get(key)
}

type customerDetailsEmitter struct {
	emit func(message assignments.CustomerIdDetails_Emitter_Message) error
}

func (e *customerDetailsEmitter) Emit(message assignments.CustomerIdDetails_Emitter_Message) error {
	return e.emit(message)
}

func (e *customerDetailsEmitter) EmitBulk(ctx context.Context, messages []assignments.CustomerIdDetails_Emitter_Message) error {
	return nil
}

type customerDetailsView struct {
	get func(key string) (*customerId.Details, error)
}

func (v *customerDetailsView) Get(key string) (*customerId.Details, error) {
	return v.get(key)
}

func (v *customerDetailsView) Keys() []string {
	return nil
}

type detailsEmitter struct {
	emit     func(message details.DeviceIdDetails_Emitter_Message) error
	emitBulk func(ctx context.Context, messages []details.DeviceIdDetails_Emitter_Message) error
}

func (e *detailsEmitter) Emit(message details.DeviceIdDetails_Emitter_Message) error {
	return e.emit(message)
}

func (e *detailsEmitter) EmitBulk(ctx context.Context, messages []details.DeviceIdDetails_Emitter_Message) error {
	return e.emitBulk(ctx, messages)
}

type heartbeatEmitter struct {
	emit     func(message heartbeats.DeviceIdHeartbeat_Emitter_Message) error
	emitBulk func(ctx context.Context, messages []heartbeats.DeviceIdHeartbeat_Emitter_Message) error
}

func (e *heartbeatEmitter) Emit(message heartbeats.DeviceIdHeartbeat_Emitter_Message) error {
	return e.emit(message)
}

func (e *heartbeatEmitter) EmitBulk(ctx context.Context, messages []heartbeats.DeviceIdHeartbeat_Emitter_Message) error {
	return e.emitBulk(ctx, messages)
}

type repository struct {
	save func(ctx context.Context, details warehouse.CustomerDetails) error
}

func (r *repository) Save(ctx context.Context, details warehouse.CustomerDetails) error {
	return r.save(ctx, details)
}
