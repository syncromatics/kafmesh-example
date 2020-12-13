package warehouse

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

// CustomerRepository provides information on known customers
type CustomerRepository struct {
	db *sql.DB
}

// NewCustomerRepository initializes a new CustomerRepository
func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

// GetCustomerDetails retrieves all known Customer details
func (r *CustomerRepository) GetCustomerDetails(ctx context.Context) ([]CustomerDetail, error) {
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
