package warehouse

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// CustomerDetails are the details about a customer
type CustomerDetails struct {
	ID   int64
	Name string
}

// CustomerDetailsRepository is the data repository for customer details
type CustomerDetailsRepository struct {
	db *sql.DB
}

// NewCustomerDetailsRepository creates a new customer details repository
func NewCustomerDetailsRepository(db *sql.DB) *CustomerDetailsRepository {
	return &CustomerDetailsRepository{db}
}

// Save saves customer details to the data store
func (r *CustomerDetailsRepository) Save(ctx context.Context, details CustomerDetails) error {
	_, err := r.db.ExecContext(ctx, `
	insert into 
		customer_details (customer_id, name)
		VALUES ( $1, $2 ) 
	on conflict (customer_id) 
	do
		update
		set name = EXCLUDED.name
	`, details.ID, details.Name)
	if err != nil {
		return errors.Wrap(err, "failed to exec update")
	}

	return nil
}

// GetAll gets all of the customer details stored in the database
func (r *CustomerDetailsRepository) GetAll(ctx context.Context) ([]CustomerDetails, error) {
	rows, err := r.db.QueryContext(ctx, `
	select
		customer_id,
		name
	from
		customer_details`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query db")
	}
	defer rows.Close()

	results := []CustomerDetails{}
	for rows.Next() {
		detail := CustomerDetails{}
		err = rows.Scan(&detail.ID, &detail.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan")
		}
		results = append(results, detail)
	}
	return results, nil
}

// Delete removes customer details from the database
func (r *CustomerDetailsRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `
	delete from
		customer_details
	where
		customer_id=$1`, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete customer details from db")
	}

	return nil
}
