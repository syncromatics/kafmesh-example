package assignments

import (
	"context"
	"strconv"
	"time"

	"kafmesh-example/internal/definitions/assignments"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/warehouse"

	"github.com/pkg/errors"
)

// Repository is a customer details repository
type Repository interface {
	GetAll(ctx context.Context) ([]warehouse.CustomerDetails, error)
}

// CustomerDetailsSynchronizer synchronizes kafka with the customer details in the database
type CustomerDetailsSynchronizer struct {
	repository Repository
}

// NewCustomerDetailsSynchronizer creates a customer details synchronizer
func NewCustomerDetailsSynchronizer(repository Repository) *CustomerDetailsSynchronizer {
	return &CustomerDetailsSynchronizer{repository}
}

// Sync syncs kafka with the repository
func (s *CustomerDetailsSynchronizer) Sync(sync assignments.CustomerIdDetails_Synchronizer_Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	customerDetails, err := s.repository.GetAll(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get details from repository")
	}

	keys, err := sync.Keys()
	if err != nil {
		return errors.Wrap(err, "failed to get keys")
	}

	cKeys := map[string]struct{}{}
	for _, k := range keys {
		cKeys[k] = struct{}{}
	}

	dKeys := map[string]struct{}{}
	for _, details := range customerDetails {
		key := strconv.Itoa(int(details.ID))
		dKeys[key] = struct{}{}

		current, err := sync.Get(key)
		if err != nil {
			return errors.Wrap(err, "failed to get")
		}

		shouldUpdate := current == nil

		if !shouldUpdate && current.Name == details.Name {
			continue
		}

		cKeys[key] = struct{}{}

		err = sync.Emit(&assignments.CustomerIdDetails_Synchronizer_Message{
			Key: key,
			Value: &customerId.Details{
				Name: details.Name,
			},
		})
		if err != nil {
			return errors.Wrap(err, "failed to emit customer details")
		}
	}

	for k := range cKeys {
		_, ok := dKeys[k]
		if ok {
			continue
		}

		err = sync.Delete(k)
		if err != nil {
			return errors.Wrap(err, "failed to delete")
		}
	}

	return nil
}
