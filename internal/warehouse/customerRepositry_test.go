package warehouse_test

import (
	"context"
	"kafmesh-example/internal/warehouse"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gotest.tools/assert"
)

func Test_GetCustomerDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"customer_id", "name"}).
		AddRow(int64(1), "some customer").
		AddRow(int64(2), "other customer")

	mock.ExpectQuery("select (.*) from customers$").WillReturnRows(rows)

	subject := warehouse.NewCustomerRepository(db)
	actual, err := subject.GetCustomerDetails(context.Background())
	assert.NilError(t, err)
	assert.DeepEqual(t, actual, []warehouse.CustomerDetail{
		{
			ID:   1,
			Name: "some customer",
		},
		{
			ID:   2,
			Name: "other customer",
		},
	})

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}
