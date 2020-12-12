package data_test

import (
	"context"
	"kafmesh-example/internal/data"
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

	subject := data.NewRepo(db)
	actual, err := subject.GetCustomerDetails(context.Background())
	assert.NilError(t, err)
	assert.DeepEqual(t, actual, []data.CustomerDetail{
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

func Test_GetDeviceAssignments(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"device_id", "customer_id"}).
		AddRow(int64(1), int64(2)).
		AddRow(int64(3), int64(4))

	mock.ExpectQuery("select (.*) from device_details$").WillReturnRows(rows)

	subject := data.NewRepo(db)
	actual, err := subject.GetDeviceAssignments(context.Background())
	assert.NilError(t, err)
	assert.DeepEqual(t, actual, []data.DeviceAssignment{
		{
			DeviceID:   1,
			CustomerID: 2,
		},
		{
			DeviceID:   3,
			CustomerID: 4,
		},
	})

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}
