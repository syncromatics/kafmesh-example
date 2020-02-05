package warehouse_test

import (
	"context"
	"kafmesh-example/internal/warehouse"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_CustomerDetailsRepository_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectExec("^insert into customer_details \\(customer_id, name\\) VALUES \\( \\$1, \\$2 \\) on conflict \\(customer_id\\) do update set name = EXCLUDED.name").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := warehouse.NewCustomerDetailsRepository(db)

	err = repository.Save(context.Background(), warehouse.CustomerDetails{
		ID:   42,
		Name: "testing customer",
	})
	assert.NilError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_CustomerDetailsRepository_Save_ShouldReturnErrorIfExecFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectExec("^insert into customer_details \\(customer_id, name\\) VALUES \\( \\$1, \\$2 \\) on conflict \\(customer_id\\) do update set name = EXCLUDED.name").
		WillReturnError(errors.Errorf("boom"))

	repository := warehouse.NewCustomerDetailsRepository(db)

	err = repository.Save(context.Background(), warehouse.CustomerDetails{
		ID:   42,
		Name: "testing customer",
	})
	assert.ErrorContains(t, err, "failed to exec update: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_CustomerDetailsRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"customer_id", "name"}).
		AddRow(42, "testing customer").
		AddRow(24, "other customer")

	mock.ExpectQuery("^select customer_id, name from customer_details").
		WillReturnRows(rows)

	repository := warehouse.NewCustomerDetailsRepository(db)

	result, err := repository.GetAll(context.Background())
	assert.NilError(t, err)

	assert.DeepEqual(t, result, []warehouse.CustomerDetails{
		warehouse.CustomerDetails{
			ID:   42,
			Name: "testing customer",
		},
		warehouse.CustomerDetails{
			ID:   24,
			Name: "other customer",
		},
	})

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_CustomerDetailsRepository_GetAll_ShouldReturnErrorIfQueryFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectQuery("^select customer_id, name from customer_details").
		WillReturnError(errors.Errorf("boom"))

	repository := warehouse.NewCustomerDetailsRepository(db)

	_, err = repository.GetAll(context.Background())
	assert.ErrorContains(t, err, "failed to query db: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_CustomerDetailsRepository_GetAll_ReturnsErrofIfScanFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"customer_id", "name"}).
		AddRow("stuff", "other customer")

	mock.ExpectQuery("^select customer_id, name from customer_details").
		WillReturnRows(rows)

	repository := warehouse.NewCustomerDetailsRepository(db)

	_, err = repository.GetAll(context.Background())
	assert.ErrorContains(t, err, "failed to scan:")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_CustomerDetailsRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectExec("^delete from customer_details where customer_id=\\$1").
		WithArgs(4).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := warehouse.NewCustomerDetailsRepository(db)

	err = repository.Delete(context.Background(), 4)
	assert.NilError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_CustomerDetailsRepository_Delete_ShouldReturnErrorIfExecFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectExec("^delete from customer_details where customer_id=\\$1").
		WithArgs(4).
		WillReturnError(errors.Errorf("boom"))

	repository := warehouse.NewCustomerDetailsRepository(db)

	err = repository.Delete(context.Background(), 4)
	assert.ErrorContains(t, err, "failed to delete customer details from db: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}
