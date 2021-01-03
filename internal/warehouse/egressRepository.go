package warehouse

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// EgressEndpoint represents an endpoint for exporting known customer data.
type EgressEndpoint struct {
	CustomerID int64
	URL        string
}

// EgressRepository saves and retrieves EgressEndpoints from the database
type EgressRepository struct {
	db *sql.DB
}

// NewEgressRepository creates an egress repository
func NewEgressRepository(db *sql.DB) *EgressRepository {
	return &EgressRepository{db}
}

// SaveEgressEndpoint saves the egress endpoint for a customer.
func (r *EgressRepository) SaveEgressEndpoint(ctx context.Context, endpoint EgressEndpoint) error {
	query := `
		insert into egress_endpoints (customer_id, endpoint_url)
		values ($1, $2)
		on conflict (customer_id)
		do update
			set endpoint_url = excluded.endpoint_url
	`
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin transaction failed")
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "failed preparing context")
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, endpoint.CustomerID, endpoint.URL)
	if err != nil {
		return errors.Wrap(err, "failed inserting endpoint")
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed committing transaction")
	}

	return nil
}

// GetAllEgressEndpoints returns all egress endpoints from the database
func (r *EgressRepository) GetAllEgressEndpoints(ctx context.Context) ([]EgressEndpoint, error) {
	query := `select customer_id, endpoint_url from egress_endpoints`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "failed getting endpoint rows")
	}
	defer rows.Close()

	endpoints := []EgressEndpoint{}
	for rows.Next() {
		e := EgressEndpoint{}
		err = rows.Scan(&e.CustomerID, &e.URL)
		if err != nil {
			return nil, errors.Wrap(err, "failed scanning endpoint")
		}
		endpoints = append(endpoints, e)
	}

	return endpoints, nil
}

// GetEgressEndpoint returns a single customer's egress endpoint, if any, from the database
func (r *EgressRepository) GetEgressEndpoint(ctx context.Context, customerID int64) (*EgressEndpoint, error) {
	query := `select customer_id, endpoint_url from egress_endpoints where customer_id = $1`
	rows, err := r.db.QueryContext(ctx, query, customerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed getting endpoint rows")
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	e := EgressEndpoint{}
	err = rows.Scan(&e.CustomerID, &e.URL)
	if err != nil {
		return nil, errors.Wrap(err, "failed scanning endpoint")
	}
	return &e, nil
}
