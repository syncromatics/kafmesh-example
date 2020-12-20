package warehouse_test

import (
	"context"
	"testing"
	"time"

	"kafmesh-example/internal/warehouse"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_DetailsRepository_SaveDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	now := time.Now()

	mock.ExpectBegin()

	mock.ExpectPrepare("")

	mock.ExpectExec(`^COPY "device_details" (.+) FROM STDIN`).
		WithArgs(int64(42), now, "test device", int64(12), "testing customer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		{
			DeviceID:     42,
			Time:         now,
			Name:         "test device",
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.NilError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_DetailsRepository_SaveDetails_ShouldReturnERrorIfCommitFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	now := time.Now()

	mock.ExpectBegin()

	mock.ExpectPrepare("")

	mock.ExpectExec(`^COPY "device_details" (.+) FROM STDIN`).
		WithArgs(int64(42), now, "test device", int64(12), "testing customer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit().WillReturnError(errors.Errorf("boom"))

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		{
			DeviceID:     42,
			Time:         now,
			Name:         "test device",
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "failed to commit: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_DetailsRepository_SaveDetails_ShouldReturnErrorIfBeginFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.Errorf("boom"))

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		warehouse.Details{
			DeviceID:     42,
			Name:         "test device",
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "begin transaction failed: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_DetailsRepository_SaveDetails_ShouldReturnErrorIfPrepareContextFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	now := time.Now()

	mock.ExpectBegin()
	mock.ExpectPrepare("").WillReturnError(errors.Errorf("boom"))
	mock.ExpectRollback()

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		warehouse.Details{
			DeviceID:     42,
			Time:         now,
			Name:         "test device",
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "failed to prepare context: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_DetailsRepository_SaveDetails_ShouldReturnErrorIfExecContextFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	now := time.Now()

	mock.ExpectBegin()
	mock.ExpectPrepare("")

	mock.ExpectExec(`^COPY "device_details" (.+) FROM STDIN`).
		WithArgs(int64(42), now, "test device", int64(12), "testing customer").
		WillReturnError(errors.Errorf("boom"))
	mock.ExpectRollback()

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		warehouse.Details{
			DeviceID:     42,
			Time:         now,
			Name:         "test device",
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "failed exec: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_DetailsRepository_SaveDetails_ShouldReturnErrorIfCloseFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	now := time.Now()

	mock.ExpectBegin()

	mock.ExpectPrepare("").WillReturnCloseError(errors.Errorf("boom"))

	mock.ExpectExec(`^COPY "device_details" (.+) FROM STDIN`).
		WithArgs(int64(42), now, "test device", int64(12), "testing customer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		warehouse.Details{
			DeviceID:     42,
			Time:         now,
			Name:         "test device",
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "failed to close: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_DetailsRepository_LastDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	now := time.Now()

	rows := sqlmock.NewRows([]string{"time", "name", "customer_id", "customer_name"}).
		AddRow(now, "test", 65, "test customer")

	mock.ExpectQuery("^select time, name, customer_id, customer_name from device_details where device_id=\\$1 order by time desc").
		WithArgs(4).
		WillReturnRows(rows)

	repository := warehouse.NewDetailsRepository(db)
	last, err := repository.LastDetails(context.Background(), 4)
	assert.NilError(t, err)

	assert.DeepEqual(t, last, &warehouse.Details{
		DeviceID:     4,
		Time:         now,
		Name:         "test",
		CustomerID:   65,
		CustomerName: "test customer",
	})

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_DetailsRepository_LastDetails_ShouldReturnNilWithNoResult(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectQuery("^select time, name, customer_id, customer_name from device_details where device_id=\\$1 order by time desc").
		WithArgs(5).
		WillReturnRows(sqlmock.NewRows([]string{"time", "name", "customer_id", "customer_name"}))

	repository := warehouse.NewDetailsRepository(db)
	last, err := repository.LastDetails(context.Background(), 5)
	assert.NilError(t, err)

	assert.Assert(t, last == nil)

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_DetailsRepository_LastDetails_ShouldReturnErrorIfQueryFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectQuery("^select time, name, customer_id, customer_name from device_details where device_id=\\$1 order by time desc").
		WithArgs(7).
		WillReturnError(errors.Errorf("boom"))

	repository := warehouse.NewDetailsRepository(db)
	_, err = repository.LastDetails(context.Background(), 7)
	assert.ErrorContains(t, err, "failed to query details: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}
