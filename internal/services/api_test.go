package services_test

import (
	"context"
	"testing"

	"kafmesh-example/internal/definitions/assignments"
	apiv1 "kafmesh-example/internal/definitions/models/kafmesh/api/v1"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/services"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_API_GetAssignment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerSource := assignments.NewMockDeviceIDCustomer_Source(ctrl)
	customerView := assignments.NewMockDeviceIDCustomer_View(ctrl)

	api := services.NewAPIService(customerSource, customerView, nil, nil)

	customerView.EXPECT().Get("42").Return(&deviceId.Customer{
		Id: 433,
	}, nil)

	r, err := api.GetAssignment(context.Background(), &apiv1.GetAssignmentRequest{
		DeviceId: 42,
	})
	assert.NilError(t, err)

	areEqual := proto.Equal(r, &apiv1.GetAssignmentResponse{
		CustomerId: &wrappers.Int64Value{Value: 433},
	})
	assert.Assert(t, areEqual)
}

func Test_API_GetAssignment_ShouldReturnErrorIfViewFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerSource := assignments.NewMockDeviceIDCustomer_Source(ctrl)
	customerView := assignments.NewMockDeviceIDCustomer_View(ctrl)

	api := services.NewAPIService(customerSource, customerView, nil, nil)

	customerView.EXPECT().Get(gomock.Any()).Return(nil, errors.Errorf("boom"))

	_, err := api.GetAssignment(context.Background(), &apiv1.GetAssignmentRequest{
		DeviceId: 42,
	})
	assert.ErrorContains(t, err, "failed to get assignment from view: boom")
}

func Test_API_GetAssignment_ShouldReturnNullForUnassignedDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerSource := assignments.NewMockDeviceIDCustomer_Source(ctrl)
	customerView := assignments.NewMockDeviceIDCustomer_View(ctrl)
	api := services.NewAPIService(customerSource, customerView, nil, nil)

	customerView.EXPECT().Get(gomock.Any()).Return(nil, nil)

	r, err := api.GetAssignment(context.Background(), &apiv1.GetAssignmentRequest{
		DeviceId: 42,
	})
	assert.NilError(t, err)

	areEqual := proto.Equal(r, &apiv1.GetAssignmentResponse{
		CustomerId: nil,
	})
	assert.Assert(t, areEqual)
}

func Test_API_AssignDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerSource := assignments.NewMockDeviceIDCustomer_Source(ctrl)
	customerView := assignments.NewMockDeviceIDCustomer_View(ctrl)

	api := services.NewAPIService(customerSource, customerView, nil, nil)

	customerSource.EXPECT().Emit(assignments.DeviceIDCustomer_Source_Message{
		Key: "42",
		Value: &deviceId.Customer{
			Id: 12,
		},
	})

	_, err := api.AssignDevice(context.Background(), &apiv1.AssignDeviceRequest{
		DeviceId:   42,
		CustomerId: 12,
	})
	assert.NilError(t, err)
}

func Test_API_AssignDevice_ShouldReturnErrorIfEmitterFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerSource := assignments.NewMockDeviceIDCustomer_Source(ctrl)
	customerView := assignments.NewMockDeviceIDCustomer_View(ctrl)

	api := services.NewAPIService(customerSource, customerView, nil, nil)

	customerSource.EXPECT().Emit(gomock.Any()).Return(errors.Errorf("boom"))

	_, err := api.AssignDevice(context.Background(), &apiv1.AssignDeviceRequest{
		DeviceId:   42,
		CustomerId: 12,
	})
	assert.ErrorContains(t, err, "failed to emit device assignment: boom")
}

func Test_API_UpdateCustomerDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	source := assignments.NewMockCustomerIDDetails_Source(ctrl)
	source.EXPECT().Emit(assignments.CustomerIDDetails_Source_Message{
		Key: "45",
		Value: &customerId.Details{
			Name: "testing customer",
		},
	})

	api := services.NewAPIService(nil, nil, source, nil)

	_, err := api.UpdateCustomerDetails(context.Background(), &apiv1.UpdateCustomerDetailsRequest{
		CustomerId: 45,
		Name:       "testing customer",
	})
	assert.NilError(t, err)
}

func Test_API_UpdateCustomerDetailsShouldReturnErrorIfEmitterFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	source := assignments.NewMockCustomerIDDetails_Source(ctrl)
	source.EXPECT().Emit(gomock.Any()).Return(errors.Errorf("boom"))

	api := services.NewAPIService(nil, nil, source, nil)

	_, err := api.UpdateCustomerDetails(context.Background(), &apiv1.UpdateCustomerDetailsRequest{
		CustomerId: 45,
		Name:       "testing customer",
	})
	assert.ErrorContains(t, err, "failed to emit customer details: boom")
}

func Test_API_GetCustomerDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	view := assignments.NewMockCustomerIDDetails_View(ctrl)
	view.EXPECT().Get("42").Return(&customerId.Details{
		Name: "testing customer",
	}, nil)

	api := services.NewAPIService(nil, nil, nil, view)

	r, err := api.GetCustomerDetails(context.Background(), &apiv1.GetCustomerDetailsRequest{
		CustomerId: 42,
	})
	assert.NilError(t, err)

	areEqual := proto.Equal(r, &apiv1.GetCustomerDetailsResponse{
		Name: &wrappers.StringValue{Value: "testing customer"},
	})
	assert.Assert(t, areEqual)
}

func Test_API_GetCustomerDetails_ShouldReturnEmptyIfDetailsDontExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	view := assignments.NewMockCustomerIDDetails_View(ctrl)
	view.EXPECT().Get("42").Return(nil, nil)

	api := services.NewAPIService(nil, nil, nil, view)

	r, err := api.GetCustomerDetails(context.Background(), &apiv1.GetCustomerDetailsRequest{
		CustomerId: 42,
	})
	assert.NilError(t, err)

	areEqual := proto.Equal(r, &apiv1.GetCustomerDetailsResponse{
		Name: nil,
	})
	assert.Assert(t, areEqual)
}

func Test_API_GetCustomerDetails_ShouldErrorIfViewFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	view := assignments.NewMockCustomerIDDetails_View(ctrl)
	view.EXPECT().Get("42").Return(nil, errors.Errorf("boom"))

	api := services.NewAPIService(nil, nil, nil, view)

	_, err := api.GetCustomerDetails(context.Background(), &apiv1.GetCustomerDetailsRequest{
		CustomerId: 42,
	})
	assert.ErrorContains(t, err, "failed to get customer details from view: boom")
}
