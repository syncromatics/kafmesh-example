package services_test

import (
	"context"
	"testing"

	"kafmesh-example/internal/definitions/assignments"
	apiv1 "kafmesh-example/internal/definitions/models/kafmesh/api/v1"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/services"
	"kafmesh-example/internal/warehouse"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_API_GetAssignment(t *testing.T) {
	customerEmitter := &customerEmitter{}
	customerView := &customerView{}

	api := services.NewAPIService(customerEmitter, customerView, nil, nil, nil)

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

	api := services.NewAPIService(customerEmitter, customerView, nil, nil, nil)

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

	api := services.NewAPIService(customerEmitter, customerView, nil, nil, nil)

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

	api := services.NewAPIService(customerEmitter, customerView, nil, nil, nil)

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

	api := services.NewAPIService(customerEmitter, customerView, nil, nil, nil)

	customerEmitter.emit = func(msg assignments.DeviceIdCustomer_Emitter_Message) error {
		return errors.Errorf("boom")
	}

	_, err := api.AssignDevice(context.Background(), &apiv1.AssignDeviceRequest{
		DeviceId:   42,
		CustomerId: 12,
	})
	assert.ErrorContains(t, err, "failed to emit device assignment: boom")
}

func Test_API_UpdateCustomerDetails(t *testing.T) {
	var emitted assignments.CustomerIdDetails_Emitter_Message
	emitter := &customerDetailsEmitter{
		emit: func(message assignments.CustomerIdDetails_Emitter_Message) error {
			emitted = message
			return nil
		},
	}

	var saved warehouse.CustomerDetails
	repository := &repository{
		save: func(ctx context.Context, details warehouse.CustomerDetails) error {
			saved = details
			return nil
		},
	}

	api := services.NewAPIService(nil, nil, emitter, nil, repository)

	_, err := api.UpdateCustomerDetails(context.Background(), &apiv1.UpdateCustomerDetailsRequest{
		CustomerId: 45,
		Name:       "testing customer",
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, emitted, assignments.CustomerIdDetails_Emitter_Message{
		Key: "45",
		Value: &customerId.Details{
			Name: "testing customer",
		},
	})

	assert.DeepEqual(t, saved, warehouse.CustomerDetails{
		ID:   45,
		Name: "testing customer",
	})
}

func Test_API_UpdateCustomerDetails_ShouldReturnerrorIfRepositoryFails(t *testing.T) {
	emitter := &customerDetailsEmitter{}

	repository := &repository{
		save: func(ctx context.Context, details warehouse.CustomerDetails) error {
			return errors.Errorf("boom")
		},
	}

	api := services.NewAPIService(nil, nil, emitter, nil, repository)

	_, err := api.UpdateCustomerDetails(context.Background(), &apiv1.UpdateCustomerDetailsRequest{
		CustomerId: 45,
		Name:       "testing customer",
	})
	assert.ErrorContains(t, err, "failed to save details to repository: boom")
}

func Test_API_UpdateCustomerDetailsShouldReturnErrorIfEmitterFails(t *testing.T) {
	emitter := &customerDetailsEmitter{
		emit: func(message assignments.CustomerIdDetails_Emitter_Message) error {
			return errors.Errorf("boom")
		},
	}

	repository := &repository{
		save: func(ctx context.Context, details warehouse.CustomerDetails) error {
			return nil
		},
	}

	api := services.NewAPIService(nil, nil, emitter, nil, repository)

	_, err := api.UpdateCustomerDetails(context.Background(), &apiv1.UpdateCustomerDetailsRequest{
		CustomerId: 45,
		Name:       "testing customer",
	})
	assert.ErrorContains(t, err, "failed to emit customer details: boom")
}

func Test_API_GetCustomerDetails(t *testing.T) {
	view := &customerDetailsView{
		get: func(key string) (*customerId.Details, error) {
			assert.Equal(t, key, "42")

			return &customerId.Details{
				Name: "testing customer",
			}, nil
		},
	}

	api := services.NewAPIService(nil, nil, nil, view, nil)

	r, err := api.GetCustomerDetails(context.Background(), &apiv1.GetCustomerDetailsRequest{
		CustomerId: 42,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, r, &apiv1.GetCustomerDetailsResponse{
		Name: &wrappers.StringValue{Value: "testing customer"},
	})
}

func Test_API_GetCustomerDetails_ShouldReturnEmptyIfDetailsDontExist(t *testing.T) {
	view := &customerDetailsView{
		get: func(key string) (*customerId.Details, error) {
			assert.Equal(t, key, "42")

			return nil, nil
		},
	}

	api := services.NewAPIService(nil, nil, nil, view, nil)

	r, err := api.GetCustomerDetails(context.Background(), &apiv1.GetCustomerDetailsRequest{
		CustomerId: 42,
	})
	assert.NilError(t, err)

	assert.DeepEqual(t, r, &apiv1.GetCustomerDetailsResponse{
		Name: nil,
	})
}

func Test_API_GetCustomerDetails_ShouldErrorIfViewFails(t *testing.T) {
	view := &customerDetailsView{
		get: func(key string) (*customerId.Details, error) {
			assert.Equal(t, key, "42")

			return nil, errors.Errorf("boom")
		},
	}

	api := services.NewAPIService(nil, nil, nil, view, nil)

	_, err := api.GetCustomerDetails(context.Background(), &apiv1.GetCustomerDetailsRequest{
		CustomerId: 42,
	})
	assert.ErrorContains(t, err, "failed to get customer details from view: boom")
}
