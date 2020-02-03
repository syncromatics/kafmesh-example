package services

import (
	"context"
	"strconv"

	"kafmesh-example/internal/definitions/details"
	"kafmesh-example/internal/definitions/heartbeats"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	gatewayv1 "kafmesh-example/internal/definitions/models/kafmesh/gateway/v1"

	"github.com/pkg/errors"
)

// GatewayService is the service for ingressing device telemetry
type GatewayService struct {
	detailsEmitter    details.DeviceIdDetails_Emitter
	heartbeatsEmitter heartbeats.DeviceIdHeartbeat_Emitter
}

// NewGatewayService creates a new gateway service
func NewGatewayService(detailsEmitter details.DeviceIdDetails_Emitter, heartbeatsEmitter heartbeats.DeviceIdHeartbeat_Emitter) *GatewayService {
	return &GatewayService{
		detailsEmitter:    detailsEmitter,
		heartbeatsEmitter: heartbeatsEmitter,
	}
}

// Details handles device details telemetry
func (s *GatewayService) Details(ctx context.Context, request *gatewayv1.DetailsRequest) (*gatewayv1.DetailsResponse, error) {
	err := s.detailsEmitter.Emit(details.DeviceIdDetails_Emitter_Message{
		Key: strconv.Itoa(int(request.DeviceId)),
		Value: &deviceId.Details{
			Name: request.Name,
			Time: request.Time,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to emit device details")
	}

	return &gatewayv1.DetailsResponse{}, nil
}

// Heartbeat handles device heartbeat telemetry
func (s *GatewayService) Heartbeat(ctx context.Context, request *gatewayv1.HeartbeatRequest) (*gatewayv1.HeartbeatResponse, error) {
	err := s.heartbeatsEmitter.Emit(heartbeats.DeviceIdHeartbeat_Emitter_Message{
		Key: strconv.Itoa(int(request.DeviceId)),
		Value: &deviceId.Heartbeat{
			Time:      request.Time,
			IsHealthy: request.IsHealthy,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to emit device heartbeat")
	}

	return &gatewayv1.HeartbeatResponse{}, nil
}
