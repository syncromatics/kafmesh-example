package assignments

import (
	"context"
	"kafmesh-example/internal/definitions/assignments"
	models "kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/warehouse"
	"strconv"

	"github.com/pkg/errors"
)

var _ assignments.CustomerDetails_ViewSource = &CustomerDetailsViewSource{}

//go:generate mockgen -source=./customerDetails_viewSource.go -destination=./mocks/mock_customer_repo.go -package=mocks
type customerRepo interface {
	GetCustomerDetails(context.Context) ([]warehouse.CustomerDetail, error)
}

// CustomerDetailsViewSource puts the latest customer details into kafka
type CustomerDetailsViewSource struct {
	repo customerRepo
}

// NewCustomerDetailsViewSource initializes a new CustomerDetailsViewSource
func NewCustomerDetailsViewSource(repo customerRepo) *CustomerDetailsViewSource {
	return &CustomerDetailsViewSource{
		repo: repo,
	}
}

// Sync outputs the latest customer details to kafka
func (vs *CustomerDetailsViewSource) Sync(ctx assignments.CustomerDetails_ViewSource_Context) error {
	customerDetails, err := vs.repo.GetCustomerDetails(ctx)
	if err != nil {
		return errors.Wrap(err, "failed getting customer details from the database")
	}

	for _, cd := range customerDetails {
		id := strconv.FormatInt(cd.ID, 10)
		msg := &models.Details{
			Name: cd.Name,
		}
		err = ctx.Update(id, msg)
		if err != nil {
			return errors.Wrap(err, "failed updating customer details")
		}
	}

	return nil
}
