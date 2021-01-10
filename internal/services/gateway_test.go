package services_test

import (
	"context"
	"testing"

	"kafmesh-example/internal/definitions/details"
	"kafmesh-example/internal/definitions/heartbeats"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	gatewayv1 "kafmesh-example/internal/definitions/models/kafmesh/gateway/v1"
	"kafmesh-example/internal/services"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_Gateway_Details(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	detailsSource := details.NewMockDeviceIDDetails_Source(ctrl)
	heartbeatsSource := heartbeats.NewMockDeviceIDHeartbeat_Source(ctrl)

	api := services.NewGatewayService(detailsSource, heartbeatsSource)

	detailsSource.EXPECT().Emit(details.DeviceIDDetails_Source_Message{
		Key: "42",
		Value: &deviceId.Details{
			Name: "testing",
		},
	})

	_, err := api.Details(context.Background(), &gatewayv1.DetailsRequest{
		DeviceId: 42,
		Name:     "testing",
	})
	assert.NilError(t, err)
}

func Test_Gateway_Details_ShouldReturnErrorIfEmitterFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	detailsSource := details.NewMockDeviceIDDetails_Source(ctrl)
	heartbeatsSource := heartbeats.NewMockDeviceIDHeartbeat_Source(ctrl)
	api := services.NewGatewayService(detailsSource, heartbeatsSource)

	detailsSource.EXPECT().Emit(gomock.Any()).Return(errors.Errorf("boom"))

	_, err := api.Details(context.Background(), &gatewayv1.DetailsRequest{
		DeviceId: 42,
		Name:     "testing",
	})
	assert.ErrorContains(t, err, "failed to emit device details: boom")
}

func Test_Gateway_Heartbeat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	detailsSource := details.NewMockDeviceIDDetails_Source(ctrl)
	heartbeatsSource := heartbeats.NewMockDeviceIDHeartbeat_Source(ctrl)

	api := services.NewGatewayService(detailsSource, heartbeatsSource)

	now := ptypes.TimestampNow()
	heartbeatsSource.EXPECT().Emit(heartbeats.DeviceIDHeartbeat_Source_Message{
		Key: "42",
		Value: &deviceId.Heartbeat{
			Time:      now,
			IsHealthy: true,
		},
	})

	_, err := api.Heartbeat(context.Background(), &gatewayv1.HeartbeatRequest{
		DeviceId:  42,
		Time:      now,
		IsHealthy: true,
	})
	assert.NilError(t, err)
}

func Test_Gateway_Heartbeat_ShouldReturnErrorIfEmitterFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	detailsSource := details.NewMockDeviceIDDetails_Source(ctrl)
	heartbeatsSource := heartbeats.NewMockDeviceIDHeartbeat_Source(ctrl)

	api := services.NewGatewayService(detailsSource, heartbeatsSource)

	heartbeatsSource.EXPECT().Emit(gomock.Any()).Return(errors.Errorf("boom"))

	now := ptypes.TimestampNow()

	_, err := api.Heartbeat(context.Background(), &gatewayv1.HeartbeatRequest{
		DeviceId:  42,
		Time:      now,
		IsHealthy: true,
	})
	assert.ErrorContains(t, err, "failed to emit device heartbeat: boom")
}
