package assignments

import (
	"context"
	"kafmesh-example/internal/definitions/assignments"
	models "kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	"kafmesh-example/internal/warehouse"
	"strconv"

	"github.com/pkg/errors"
)

var _ assignments.DeviceCustomer_ViewSource = &DeviceCustomerViewSource{}

//go:generate mockgen -source=./deviceCustomerSink.go -destination=./mocks/mock_assignments_repo.go -package=mocks
type assignmentRepo interface {
	AllDetails(context.Context) ([]warehouse.Details, error)
}

// DeviceCustomerViewSource puts the latest device assignments into kafka
type DeviceCustomerViewSource struct {
	repo assignmentRepo
}

// NewDeviceCustomerViewSource initializes a new DeviceCustomerViewSource
func NewDeviceCustomerViewSource(repo assignmentRepo) *DeviceCustomerViewSource {
	return &DeviceCustomerViewSource{
		repo: repo,
	}
}

// Sync outputs the latest device assignments to kafka
func (vs *DeviceCustomerViewSource) Sync(ctx assignments.DeviceCustomer_ViewSource_Context) error {
	assignments, err := vs.repo.AllDetails(ctx)
	if err != nil {
		return errors.Wrap(err, "failed getting device assignments from the database")
	}

	for _, a := range assignments {
		deviceID := strconv.FormatInt(a.DeviceID, 10)
		msg := &models.Customer{
			Id: a.CustomerID,
		}
		err = ctx.Update(deviceID, msg)
		if err != nil {
			return errors.Wrap(err, "failed updating device assignment")
		}
	}

	return nil
}
