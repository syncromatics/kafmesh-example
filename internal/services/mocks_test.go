package services_test

import (
	"context"

	"kafmesh-example/internal/definitions/assignments"
	"kafmesh-example/internal/definitions/details"
	"kafmesh-example/internal/definitions/heartbeats"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
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
