package services

import (
	"context"

	historyv1 "kafmesh-example/internal/definitions/models/kafmesh/history/v1"
	"kafmesh-example/internal/warehouse"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
)

//go:generate mockgen -source=./history.go -destination=./history_mock_test.go -package=services_test

// DetailsRepository gets the last details from the warehouse
type DetailsRepository interface {
	LastDetails(context.Context, int64) (*warehouse.Details, error)
}

// HeartbeatsRepository gets the last heartbeat from the warehouse
type HeartbeatsRepository interface {
	LastHeartbeat(context.Context, int64) (*warehouse.Heartbeat, error)
}

// HistoryAPI provides access to historical device telemetry
type HistoryAPI struct {
	detailsRepository    DetailsRepository
	HeartbeatsRepository HeartbeatsRepository
}

// NewHistoryAPI creates a new history api
func NewHistoryAPI(detailsRepository DetailsRepository, heartbeatsRepository HeartbeatsRepository) *HistoryAPI {
	return &HistoryAPI{
		detailsRepository:    detailsRepository,
		HeartbeatsRepository: heartbeatsRepository,
	}
}

// LastDetails returns the last details for a device. Will return none if no details are in the database for the device.
func (s *HistoryAPI) LastDetails(ctx context.Context, request *historyv1.LastDetailsRequest) (*historyv1.LastDetailsResponse, error) {
	d, err := s.detailsRepository.LastDetails(ctx, request.DeviceId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get last details from repository")
	}

	if d == nil {
		return &historyv1.LastDetailsResponse{
			Response: &historyv1.LastDetailsResponse_ResponseNone{
				ResponseNone: &historyv1.None{},
			},
		}, nil
	}

	t, err := ptypes.TimestampProto(d.Time)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert timestamp")
	}

	return &historyv1.LastDetailsResponse{
		Response: &historyv1.LastDetailsResponse_ResponseDetails{
			ResponseDetails: &historyv1.Details{
				Name:         d.Name,
				Time:         t,
				CustomerId:   d.CustomerID,
				CustomerName: d.CustomerName,
			},
		},
	}, nil
}

// LastHeartbeat returns the last heartbeat for a device. Will return none if no heartbeats are in the database for the device.
func (s *HistoryAPI) LastHeartbeat(ctx context.Context, request *historyv1.LastHeartbeatRequest) (*historyv1.LastHeartbeatResponse, error) {
	h, err := s.HeartbeatsRepository.LastHeartbeat(ctx, request.DeviceId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get last heartbeat from repository")
	}

	if h == nil {
		return &historyv1.LastHeartbeatResponse{
			Response: &historyv1.LastHeartbeatResponse_ResponseNone{
				ResponseNone: &historyv1.None{},
			},
		}, nil
	}

	t, err := ptypes.TimestampProto(h.Time)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert to timestamp")
	}

	return &historyv1.LastHeartbeatResponse{
		Response: &historyv1.LastHeartbeatResponse_ResponseHeartbeat{
			ResponseHeartbeat: &historyv1.Heartbeat{
				Time:         t,
				IsHealthy:    h.IsHealthy,
				CustomerId:   h.CustomerID,
				CustomerName: h.CustomerName,
			},
		},
	}, nil
}
