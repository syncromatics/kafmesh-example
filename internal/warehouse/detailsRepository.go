package warehouse

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// Details is the historical details update from the device
type Details struct {
	DeviceID     int64
	Name         string
	CustomerID   int64
	CustomerName string
}

// DetailsRepository saves and retrieves details from the database
type DetailsRepository struct {
	db *sql.DB
}

// NewDetailsRepository creates a details repository
func NewDetailsRepository(db *sql.DB) *DetailsRepository {
	return &DetailsRepository{db}
}

// SaveDetails saves details to the warehouse
func (r *DetailsRepository) SaveDetails(ctx context.Context, details []Details) error {
	txn, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, "begin transaction failed")
	}
	defer txn.Rollback()

	stmt, err := txn.PrepareContext(ctx, pq.CopyIn("device_details", "device_id", "name", "customer_id", "customer_name"))
	if err != nil {
		return errors.Wrap(err, "failed to prepare context")
	}

	for _, d := range details {
		_, err = stmt.ExecContext(ctx, d.DeviceID, d.Name, d.CustomerID, d.CustomerName)
		if err != nil {
			return errors.Wrap(err, "failed exec")
		}
	}

	err = stmt.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close")
	}

	return nil
}
