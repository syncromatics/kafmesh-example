package warehouse_test

import (
	"context"
	"testing"

	"kafmesh-example/internal/warehouse"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_DetailsRepository_SaveDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare("")

	mock.ExpectExec(`^COPY "device_details" (.+) FROM STDIN*`).
		WithArgs(int64(42), "test device", int64(12), "testing customer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		warehouse.Details{
			DeviceID:     42,
			Name:         "test device",
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.NilError(t, err)

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

	mock.ExpectBegin()
	mock.ExpectPrepare("").WillReturnError(errors.Errorf("boom"))
	mock.ExpectRollback()

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		warehouse.Details{
			DeviceID:     42,
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

	mock.ExpectBegin()
	mock.ExpectPrepare("")

	mock.ExpectExec(`^COPY "device_details" (.+) FROM STDIN*`).
		WithArgs(int64(42), "test device", int64(12), "testing customer").
		WillReturnError(errors.Errorf("boom"))
	mock.ExpectRollback()

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		warehouse.Details{
			DeviceID:     42,
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

	mock.ExpectBegin()

	mock.ExpectPrepare("").WillReturnCloseError(errors.Errorf("boom"))

	mock.ExpectExec(`^COPY "device_details" (.+) FROM STDIN*`).
		WithArgs(int64(42), "test device", int64(12), "testing customer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := warehouse.NewDetailsRepository(db)

	err = repository.SaveDetails(context.Background(), []warehouse.Details{
		warehouse.Details{
			DeviceID:     42,
			Name:         "test device",
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "failed to close: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}
