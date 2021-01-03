package warehouse_test

import (
	"context"
	"kafmesh-example/internal/warehouse"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gotest.tools/assert"
)

func Test_EgressRepository_SaveEgressEndpoint(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("insert into egress_endpoints")
	mock.ExpectExec("insert into egress_endpoints").
		WithArgs(int64(42), "https://example.com/").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := warehouse.NewEgressRepository(db)
	endpoint := warehouse.EgressEndpoint{
		CustomerID: 42,
		URL:        "https://example.com/",
	}
	err = repo.SaveEgressEndpoint(context.Background(), endpoint)
	assert.NilError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_EgressRepository_GetAllEgressEndpoints(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"customer_id", "endpoint_url"}).
		AddRow(int64(42), "https://example.com/1").
		AddRow(int64(43), "https://example.com/2")
	mock.ExpectQuery("select .* from egress_endpoints").
		WillReturnRows(rows)

	repo := warehouse.NewEgressRepository(db)
	actual, err := repo.GetAllEgressEndpoints(context.Background())
	assert.NilError(t, err)

	assert.DeepEqual(t, actual, []warehouse.EgressEndpoint{
		{
			CustomerID: int64(42),
			URL:        "https://example.com/1",
		},
		{
			CustomerID: 43,
			URL:        "https://example.com/2",
		},
	})
}

func Test_EgressRepository_GetEgressEndpoint(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"customer_id", "endpoint_url"}).
		AddRow(int64(42), "https://example.com/1")
	mock.ExpectQuery("select .* from egress_endpoints").
		WithArgs(int64(42)).
		WillReturnRows(rows)

	repo := warehouse.NewEgressRepository(db)
	actual, err := repo.GetEgressEndpoint(context.Background(), 42)
	assert.NilError(t, err)

	assert.DeepEqual(t, actual, &warehouse.EgressEndpoint{
		CustomerID: int64(42),
		URL:        "https://example.com/1",
	})
}
