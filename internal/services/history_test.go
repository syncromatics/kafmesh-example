package services_test

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"

	historyv1 "kafmesh-example/internal/definitions/models/kafmesh/history/v1"
	"kafmesh-example/internal/services"
	"kafmesh-example/internal/warehouse"

	"github.com/pkg/errors"
	"gotest.tools/assert"
)

type detailsRepository struct {
	lastDetails func(ctx context.Context, deviceID int64) (*warehouse.Details, error)
}

func (r *detailsRepository) LastDetails(ctx context.Context, deviceID int64) (*warehouse.Details, error) {
	return r.lastDetails(ctx, deviceID)
}

func Test_HistoryService_LastDetails_ShouldReturnDetails(t *testing.T) {
	now := time.Now()
	repo := &detailsRepository{
		lastDetails: func(ctx context.Context, deviceID int64) (*warehouse.Details, error) {
			assert.Equal(t, deviceID, int64(12))
			return &warehouse.Details{
				CustomerID:   42,
				CustomerName: "testing customer",
				DeviceID:     12,
				Name:         "stuff",
				Time:         now,
			}, nil
		},
	}

	service := services.NewHistoryAPI(repo, nil)

	r, err := service.LastDetails(context.Background(), &historyv1.LastDetailsRequest{
		DeviceId: 12,
	})
	assert.NilError(t, err)

	tt, _ := ptypes.TimestampProto(now)
	assert.DeepEqual(t, r, &historyv1.LastDetailsResponse{
		Response: &historyv1.LastDetailsResponse_ResponseDetails{
			ResponseDetails: &historyv1.Details{
				CustomerId:   42,
				CustomerName: "testing customer",
				Name:         "stuff",
				Time:         tt,
			},
		},
	})
}

func Test_HistoryService_LastDetails_ShouldReturnNoneIfNoDetails(t *testing.T) {
	repo := &detailsRepository{
		lastDetails: func(ctx context.Context, deviceID int64) (*warehouse.Details, error) {
			assert.Equal(t, deviceID, int64(12))
			return nil, nil
		},
	}

	service := services.NewHistoryAPI(repo, nil)

	r, err := service.LastDetails(context.Background(), &historyv1.LastDetailsRequest{
		DeviceId: 12,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, r, &historyv1.LastDetailsResponse{
		Response: &historyv1.LastDetailsResponse_ResponseNone{
			ResponseNone: &historyv1.None{},
		},
	})
}

func Test_HistoryService_LastDetails_ShouldReturnErrorIfRepoFails(t *testing.T) {
	repo := &detailsRepository{
		lastDetails: func(ctx context.Context, deviceID int64) (*warehouse.Details, error) {
			assert.Equal(t, deviceID, int64(12))
			return nil, errors.Errorf("boom")
		},
	}

	service := services.NewHistoryAPI(repo, nil)

	_, err := service.LastDetails(context.Background(), &historyv1.LastDetailsRequest{
		DeviceId: 12,
	})
	assert.ErrorContains(t, err, "failed to get last details from repository: boom")
}

func Test_HistoryService_LastDetails_ShouldReturnErrorIfTimestamp(t *testing.T) {
	repo := &detailsRepository{
		lastDetails: func(ctx context.Context, deviceID int64) (*warehouse.Details, error) {
			assert.Equal(t, deviceID, int64(12))
			return &warehouse.Details{
				Time: time.Unix(math.MaxInt64, math.MaxInt64),
			}, nil
		},
	}

	service := services.NewHistoryAPI(repo, nil)

	_, err := service.LastDetails(context.Background(), &historyv1.LastDetailsRequest{
		DeviceId: 12,
	})
	assert.ErrorContains(t, err, "failed to convert timestamp: timestamp")
}

type heartbeatsRepository struct {
	lastHeartbeat func(ctx context.Context, deviceID int64) (*warehouse.Heartbeat, error)
}

func (r *heartbeatsRepository) LastHeartbeat(ctx context.Context, deviceID int64) (*warehouse.Heartbeat, error) {
	return r.lastHeartbeat(ctx, deviceID)
}

func Test_HistoryService_LastHeartbeat_ShouldReturnLastHeartbeat(t *testing.T) {
	now := time.Now()
	tNow, _ := ptypes.TimestampProto(now)

	repo := &heartbeatsRepository{
		lastHeartbeat: func(ctx context.Context, deviceID int64) (*warehouse.Heartbeat, error) {
			assert.Equal(t, deviceID, int64(45))

			return &warehouse.Heartbeat{
				Time:         now,
				IsHealthy:    true,
				DeviceID:     45,
				CustomerID:   67,
				CustomerName: "testing customer",
			}, nil
		},
	}

	service := services.NewHistoryAPI(nil, repo)

	r, err := service.LastHeartbeat(context.Background(), &historyv1.LastHeartbeatRequest{
		DeviceId: 45,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, r, &historyv1.LastHeartbeatResponse{
		Response: &historyv1.LastHeartbeatResponse_ResponseHeartbeat{
			ResponseHeartbeat: &historyv1.Heartbeat{
				Time:         tNow,
				IsHealthy:    true,
				CustomerId:   67,
				CustomerName: "testing customer",
			},
		},
	})
}

func Test_HistoryService_LastHeartbeat_ShouldReturnNoneIfNoHeartbeat(t *testing.T) {
	repo := &heartbeatsRepository{
		lastHeartbeat: func(ctx context.Context, deviceID int64) (*warehouse.Heartbeat, error) {
			assert.Equal(t, deviceID, int64(45))

			return nil, nil
		},
	}

	service := services.NewHistoryAPI(nil, repo)

	r, err := service.LastHeartbeat(context.Background(), &historyv1.LastHeartbeatRequest{
		DeviceId: 45,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, r, &historyv1.LastHeartbeatResponse{
		Response: &historyv1.LastHeartbeatResponse_ResponseNone{
			ResponseNone: &historyv1.None{},
		},
	})
}

func Test_HistoryService_LastHeartbeat_ShouldReturnErrorIfRepoFails(t *testing.T) {
	repo := &heartbeatsRepository{
		lastHeartbeat: func(ctx context.Context, deviceID int64) (*warehouse.Heartbeat, error) {
			assert.Equal(t, deviceID, int64(45))

			return nil, errors.Errorf("boom")
		},
	}

	service := services.NewHistoryAPI(nil, repo)

	_, err := service.LastHeartbeat(context.Background(), &historyv1.LastHeartbeatRequest{
		DeviceId: 45,
	})
	assert.ErrorContains(t, err, "failed to get last heartbeat from repository: boom")
}

func Test_HistoryService_LastHeartbeat_ShouldReturnErrorIfTimestampFails(t *testing.T) {
	repo := &heartbeatsRepository{
		lastHeartbeat: func(ctx context.Context, deviceID int64) (*warehouse.Heartbeat, error) {
			assert.Equal(t, deviceID, int64(45))

			return &warehouse.Heartbeat{
				Time: time.Unix(math.MaxInt64, math.MaxInt64),
			}, nil
		},
	}

	service := services.NewHistoryAPI(nil, repo)

	_, err := service.LastHeartbeat(context.Background(), &historyv1.LastHeartbeatRequest{
		DeviceId: 45,
	})
	assert.ErrorContains(t, err, "failed to convert to timestamp")
}
