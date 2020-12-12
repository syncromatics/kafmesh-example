package data

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// CustomerDetail represents a customer in the database
type CustomerDetail struct {
	ID   int64
	Name string
}

// DeviceAssignment represents a device assigned to a customer
type DeviceAssignment struct {
	DeviceID   int64
	CustomerID int64
}

// Repo is a (sparse) repository for retrieving customer information
type Repo struct {
	db *sql.DB
}

// NewRepo initializes a new CustomersRepo
func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// GetCustomerDetails retrieves all known Customer details
func (r *Repo) GetCustomerDetails(ctx context.Context) ([]CustomerDetail, error) {
	query := `select customer_id, name from customers`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query customers rows")
	}
	defer rows.Close()

	customers := []CustomerDetail{}
	for rows.Next() {
		c := CustomerDetail{}
		err = rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed scanning customer")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// GetDeviceAssignments retrieves all known device assignments
func (r *Repo) GetDeviceAssignments(ctx context.Context) ([]DeviceAssignment, error) {
	query := `select device_id, customer_id from device_details`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query device_details rows")
	}
	defer rows.Close()

	assignments := []DeviceAssignment{}
	for rows.Next() {
		a := DeviceAssignment{}
		err = rows.Scan(&a.DeviceID, &a.CustomerID)
		if err != nil {
			return nil, errors.Wrap(err, "failed scanning device_details")
		}
		assignments = append(assignments, a)
	}

	return assignments, nil
}
