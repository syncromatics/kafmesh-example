package services_test

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	historyv1 "kafmesh-example/internal/definitions/models/kafmesh/history/v1"
	"kafmesh-example/internal/services"
	"kafmesh-example/internal/warehouse"

	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_HistoryService_LastDetails_ShouldReturnDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockDetailsRepository(ctrl)
	now := time.Now()

	repo.EXPECT().LastDetails(gomock.Any(), int64(12)).Return(&warehouse.Details{
		CustomerID:   42,
		CustomerName: "testing customer",
		DeviceID:     12,
		Name:         "stuff",
		Time:         now,
	}, nil)

	service := services.NewHistoryAPI(repo, nil)

	r, err := service.LastDetails(context.Background(), &historyv1.LastDetailsRequest{
		DeviceId: 12,
	})
	assert.NilError(t, err)

	tt, _ := ptypes.TimestampProto(now)
	areEqual := proto.Equal(r, &historyv1.LastDetailsResponse{
		Response: &historyv1.LastDetailsResponse_ResponseDetails{
			ResponseDetails: &historyv1.Details{
				CustomerId:   42,
				CustomerName: "testing customer",
				Name:         "stuff",
				Time:         tt,
			},
		},
	})
	assert.Assert(t, areEqual)
}

func Test_HistoryService_LastDetails_ShouldReturnNoneIfNoDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockDetailsRepository(ctrl)
	repo.EXPECT().LastDetails(gomock.Any(), int64(12)).Return(nil, nil)

	service := services.NewHistoryAPI(repo, nil)

	r, err := service.LastDetails(context.Background(), &historyv1.LastDetailsRequest{
		DeviceId: 12,
	})
	assert.NilError(t, err)

	areEqual := proto.Equal(r, &historyv1.LastDetailsResponse{
		Response: &historyv1.LastDetailsResponse_ResponseNone{
			ResponseNone: &historyv1.None{},
		},
	})
	assert.Assert(t, areEqual)
}

func Test_HistoryService_LastDetails_ShouldReturnErrorIfRepoFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockDetailsRepository(ctrl)
	repo.EXPECT().LastDetails(gomock.Any(), int64(12)).Return(nil, errors.Errorf("boom"))

	service := services.NewHistoryAPI(repo, nil)

	_, err := service.LastDetails(context.Background(), &historyv1.LastDetailsRequest{
		DeviceId: 12,
	})
	assert.ErrorContains(t, err, "failed to get last details from repository: boom")
}

func Test_HistoryService_LastHeartbeat_ShouldReturnLastHeartbeat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	now := time.Now()
	tNow, _ := ptypes.TimestampProto(now)

	repo := NewMockHeartbeatsRepository(ctrl)
	repo.EXPECT().LastHeartbeat(gomock.Any(), int64(45)).Return(&warehouse.Heartbeat{
		Time:         now,
		IsHealthy:    true,
		DeviceID:     45,
		CustomerID:   67,
		CustomerName: "testing customer",
	}, nil)

	service := services.NewHistoryAPI(nil, repo)

	r, err := service.LastHeartbeat(context.Background(), &historyv1.LastHeartbeatRequest{
		DeviceId: 45,
	})
	assert.NilError(t, err)

	areEqual := proto.Equal(r, &historyv1.LastHeartbeatResponse{
		Response: &historyv1.LastHeartbeatResponse_ResponseHeartbeat{
			ResponseHeartbeat: &historyv1.Heartbeat{
				Time:         tNow,
				IsHealthy:    true,
				CustomerId:   67,
				CustomerName: "testing customer",
			},
		},
	})
	assert.Assert(t, areEqual)
}

func Test_HistoryService_LastHeartbeat_ShouldReturnNoneIfNoHeartbeat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockHeartbeatsRepository(ctrl)
	repo.EXPECT().LastHeartbeat(gomock.Any(), int64(45)).Return(nil, nil)

	service := services.NewHistoryAPI(nil, repo)

	r, err := service.LastHeartbeat(context.Background(), &historyv1.LastHeartbeatRequest{
		DeviceId: 45,
	})
	assert.NilError(t, err)

	areEqual := proto.Equal(r, &historyv1.LastHeartbeatResponse{
		Response: &historyv1.LastHeartbeatResponse_ResponseNone{
			ResponseNone: &historyv1.None{},
		},
	})
	assert.Assert(t, areEqual)
}

func Test_HistoryService_LastHeartbeat_ShouldReturnErrorIfRepoFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockHeartbeatsRepository(ctrl)
	repo.EXPECT().LastHeartbeat(gomock.Any(), int64(45)).Return(nil, errors.Errorf("boom"))

	service := services.NewHistoryAPI(nil, repo)

	_, err := service.LastHeartbeat(context.Background(), &historyv1.LastHeartbeatRequest{
		DeviceId: 45,
	})
	assert.ErrorContains(t, err, "failed to get last heartbeat from repository: boom")
}

func Test_HistoryService_LastHeartbeat_ShouldReturnErrorIfTimestampFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockHeartbeatsRepository(ctrl)
	repo.EXPECT().LastHeartbeat(gomock.Any(), int64(45)).Return(&warehouse.Heartbeat{
		Time: time.Unix(math.MaxInt64, math.MaxInt64),
	}, nil)

	service := services.NewHistoryAPI(nil, repo)

	_, err := service.LastHeartbeat(context.Background(), &historyv1.LastHeartbeatRequest{
		DeviceId: 45,
	})
	assert.ErrorContains(t, err, "failed to convert to timestamp")
}
