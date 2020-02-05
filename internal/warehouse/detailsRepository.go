package warehouse

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// Details is the historical details update from the device
type Details struct {
	DeviceID     int64
	Time         time.Time
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

	stmt, err := txn.PrepareContext(ctx, pq.CopyIn("device_details", "device_id", "time", "name", "customer_id", "customer_name"))
	if err != nil {
		return errors.Wrap(err, "failed to prepare context")
	}

	for _, d := range details {
		_, err = stmt.ExecContext(ctx, d.DeviceID, d.Time, d.Name, d.CustomerID, d.CustomerName)
		if err != nil {
			return errors.Wrap(err, "failed exec")
		}
	}

	err = stmt.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close")
	}

	err = txn.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit")
	}

	return nil
}

// LastDetails returnswarehouse the last details for a device
func (r *DetailsRepository) LastDetails(ctx context.Context, deviceID int64) (*Details, error) {
	row := r.db.QueryRowContext(ctx, `
select
	time, name, customer_id, customer_name
from
	device_details
where
	device_id=$1
order by
	time desc`, deviceID)

	details := &Details{
		DeviceID: deviceID,
	}
	err := row.Scan(&details.Time, &details.Name, &details.CustomerID, &details.CustomerName)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to query details")
	}

	return details, nil
}
