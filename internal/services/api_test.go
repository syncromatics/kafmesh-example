package services_test

import (
	"context"
	"testing"

	"kafmesh-example/internal/definitions/assignments"
	apiv1 "kafmesh-example/internal/definitions/models/kafmesh/api/v1"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/services"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_API_GetAssignment(t *testing.T) {
	customerEmitter := &customerEmitter{}
	customerView := &customerView{}

	api := services.NewAPIService(customerEmitter, customerView)

	customerView.get = func(key string) (*deviceId.Customer, error) {
		assert.Equal(t, key, "42")
		return &deviceId.Customer{Id: 433}, nil
	}

	r, err := api.GetAssignment(context.Background(), &apiv1.GetAssignmentRequest{
		DeviceId: 42,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, r, &apiv1.GetAssignmentResponse{
		CustomerId: &wrappers.Int64Value{Value: 433},
	})
}

func Test_API_GetAssignment_ShouldReturnErrorIfViewFails(t *testing.T) {
	customerEmitter := &customerEmitter{}
	customerView := &customerView{}

	api := services.NewAPIService(customerEmitter, customerView)

	customerView.get = func(key string) (*deviceId.Customer, error) {
		return nil, errors.Errorf("boom")
	}

	_, err := api.GetAssignment(context.Background(), &apiv1.GetAssignmentRequest{
		DeviceId: 42,
	})
	assert.ErrorContains(t, err, "failed to get assignment from view: boom")
}

func Test_API_GetAssignment_ShouldReturnNullForUnassignedDevice(t *testing.T) {
	customerEmitter := &customerEmitter{}
	customerView := &customerView{}

	api := services.NewAPIService(customerEmitter, customerView)

	customerView.get = func(key string) (*deviceId.Customer, error) {
		return nil, nil
	}

	r, err := api.GetAssignment(context.Background(), &apiv1.GetAssignmentRequest{
		DeviceId: 42,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, r, &apiv1.GetAssignmentResponse{
		CustomerId: nil,
	})
}

func Test_API_AssignDevice(t *testing.T) {
	customerEmitter := &customerEmitter{}
	customerView := &customerView{}

	api := services.NewAPIService(customerEmitter, customerView)

	var emitted assignments.DeviceIdCustomer_Emitter_Message
	customerEmitter.emit = func(msg assignments.DeviceIdCustomer_Emitter_Message) error {
		emitted = msg
		return nil
	}

	_, err := api.AssignDevice(context.Background(), &apiv1.AssignDeviceRequest{
		DeviceId:   42,
		CustomerId: 12,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, emitted, assignments.DeviceIdCustomer_Emitter_Message{
		Key: "42",
		Value: &deviceId.Customer{
			Id: 12,
		},
	})
}

func Test_API_AssignDevice_ShouldReturnErrorIfEmitterFails(t *testing.T) {
	customerEmitter := &customerEmitter{}
	customerView := &customerView{}

	api := services.NewAPIService(customerEmitter, customerView)

	customerEmitter.emit = func(msg assignments.DeviceIdCustomer_Emitter_Message) error {
		return errors.Errorf("boom")
	}

	_, err := api.AssignDevice(context.Background(), &apiv1.AssignDeviceRequest{
		DeviceId:   42,
		CustomerId: 12,
	})
	assert.ErrorContains(t, err, "failed to emit device assignment: boom")
}
