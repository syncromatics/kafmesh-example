package services_test

import (
	"context"
	"testing"

	"kafmesh-example/internal/definitions/details"
	"kafmesh-example/internal/definitions/heartbeats"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	gatewayv1 "kafmesh-example/internal/definitions/models/kafmesh/gateway/v1"
	"kafmesh-example/internal/services"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_Gateway_Details(t *testing.T) {
	detailsEmitter := &detailsEmitter{}
	heartbeatEmitter := &heartbeatEmitter{}

	api := services.NewGatewayService(detailsEmitter, heartbeatEmitter)

	var emitted details.DeviceIdDetails_Emitter_Message
	detailsEmitter.emit = func(msg details.DeviceIdDetails_Emitter_Message) error {
		emitted = msg
		return nil
	}

	_, err := api.Details(context.Background(), &gatewayv1.DetailsRequest{
		DeviceId: 42,
		Name:     "testing",
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, emitted, details.DeviceIdDetails_Emitter_Message{
		Key: "42",
		Value: &deviceId.Details{
			Name: "testing",
		},
	})
}

func Test_Gateway_Details_ShouldReturnErrorIfEmitterFails(t *testing.T) {
	detailsEmitter := &detailsEmitter{}
	heartbeatEmitter := &heartbeatEmitter{}

	api := services.NewGatewayService(detailsEmitter, heartbeatEmitter)

	detailsEmitter.emit = func(msg details.DeviceIdDetails_Emitter_Message) error {
		return errors.Errorf("boom")
	}

	_, err := api.Details(context.Background(), &gatewayv1.DetailsRequest{
		DeviceId: 42,
		Name:     "testing",
	})
	assert.ErrorContains(t, err, "failed to emit device details: boom")
}

func Test_Gateway_Heartbeat(t *testing.T) {
	detailsEmitter := &detailsEmitter{}
	heartbeatEmitter := &heartbeatEmitter{}

	api := services.NewGatewayService(detailsEmitter, heartbeatEmitter)

	var emitted heartbeats.DeviceIdHeartbeat_Emitter_Message
	heartbeatEmitter.emit = func(msg heartbeats.DeviceIdHeartbeat_Emitter_Message) error {
		emitted = msg
		return nil
	}

	now := ptypes.TimestampNow()

	_, err := api.Heartbeat(context.Background(), &gatewayv1.HeartbeatRequest{
		DeviceId:  42,
		Time:      now,
		IsHealthy: true,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, emitted, heartbeats.DeviceIdHeartbeat_Emitter_Message{
		Key: "42",
		Value: &deviceId.Heartbeat{
			Time:      now,
			IsHealthy: true,
		},
	})
}

func Test_Gateway_Heartbeat_ShouldReturnErrorIfEmitterFails(t *testing.T) {
	detailsEmitter := &detailsEmitter{}
	heartbeatEmitter := &heartbeatEmitter{}

	api := services.NewGatewayService(detailsEmitter, heartbeatEmitter)

	heartbeatEmitter.emit = func(msg heartbeats.DeviceIdHeartbeat_Emitter_Message) error {
		return errors.Errorf("boom")
	}

	now := ptypes.TimestampNow()

	_, err := api.Heartbeat(context.Background(), &gatewayv1.HeartbeatRequest{
		DeviceId:  42,
		Time:      now,
		IsHealthy: true,
	})
	assert.ErrorContains(t, err, "failed to emit device heartbeat: boom")
}
