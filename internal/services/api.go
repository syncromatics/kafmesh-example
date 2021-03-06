package services

import (
	"context"
	"strconv"

	"kafmesh-example/internal/definitions/assignments"
	apiv1 "kafmesh-example/internal/definitions/models/kafmesh/api/v1"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
)

// APIService is the service that provides device configuration management
type APIService struct {
	source        assignments.DeviceIDCustomer_Source
	view          assignments.DeviceIDCustomer_View
	detailsSource assignments.CustomerIDDetails_Source
	detailsView   assignments.CustomerIDDetails_View
}

// NewAPIService creates a new api service
func NewAPIService(
	source assignments.DeviceIDCustomer_Source,
	view assignments.DeviceIDCustomer_View,
	detailsEmitter assignments.CustomerIDDetails_Source,
	detailsView assignments.CustomerIDDetails_View) *APIService {
	return &APIService{
		source:        source,
		view:          view,
		detailsSource: detailsEmitter,
		detailsView:   detailsView,
	}
}

// GetAssignment retrives the most recent details sent by the device
func (s *APIService) GetAssignment(ctx context.Context, request *apiv1.GetAssignmentRequest) (*apiv1.GetAssignmentResponse, error) {
	customer, err := s.view.Get(strconv.Itoa(int(request.DeviceId)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get assignment from view")
	}

	if customer != nil {
		return &apiv1.GetAssignmentResponse{
			CustomerId: &wrappers.Int64Value{Value: customer.Id},
		}, nil
	}

	return &apiv1.GetAssignmentResponse{}, nil
}

// AssignDevice assigns the device to a customer
func (s *APIService) AssignDevice(ctx context.Context, request *apiv1.AssignDeviceRequest) (*apiv1.AssignDeviceResponse, error) {
	err := s.source.Emit(assignments.DeviceIDCustomer_Source_Message{
		Key: strconv.Itoa(int(request.DeviceId)),
		Value: &deviceId.Customer{
			Id: request.CustomerId,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to emit device assignment")
	}

	return &apiv1.AssignDeviceResponse{}, nil
}

// GetCustomerDetails returns details about a customer
func (s *APIService) GetCustomerDetails(ctx context.Context, request *apiv1.GetCustomerDetailsRequest) (*apiv1.GetCustomerDetailsResponse, error) {
	customer, err := s.detailsView.Get(strconv.Itoa(int(request.CustomerId)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get customer details from view")
	}

	if customer != nil {
		return &apiv1.GetCustomerDetailsResponse{
			Name: &wrappers.StringValue{Value: customer.Name},
		}, nil
	}
	return &apiv1.GetCustomerDetailsResponse{}, nil
}

// UpdateCustomerDetails updates the customer's details
func (s *APIService) UpdateCustomerDetails(ctx context.Context, request *apiv1.UpdateCustomerDetailsRequest) (*apiv1.UpdateCustomerDetailsResponse, error) {
	err := s.detailsSource.Emit(assignments.CustomerIDDetails_Source_Message{
		Key: strconv.Itoa(int(request.CustomerId)),
		Value: &customerId.Details{
			Name: request.Name,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to emit customer details")
	}

	return &apiv1.UpdateCustomerDetailsResponse{}, nil
}
