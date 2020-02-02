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

func Test_HeartbeatRepository_SaveHeartbeats(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare("")

	now := time.Now()

	mock.ExpectExec(`^COPY "device_heartbeats" (.+) FROM STDIN*`).
		WithArgs(int64(42), now, true, int64(12), "testing customer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := warehouse.NewHeartbeatRepository(db)

	err = repository.SaveHeartbeats(context.Background(), []warehouse.Heartbeat{
		warehouse.Heartbeat{
			DeviceID:     42,
			Time:         now,
			IsHealthy:    true,
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.NilError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_HeartbeatRepository_SaveHeartbeats_ShouldReturnErrorIfBeginFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectBegin().WillReturnError(errors.Errorf("boom"))

	now := time.Now()

	repository := warehouse.NewHeartbeatRepository(db)

	err = repository.SaveHeartbeats(context.Background(), []warehouse.Heartbeat{
		warehouse.Heartbeat{
			DeviceID:     42,
			Time:         now,
			IsHealthy:    true,
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "begin transaction failed: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_HeartbeatRepository_SaveHeartbeats_ShouldReturnErrorIfPrepareFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare("").WillReturnError(errors.Errorf("boom"))

	mock.ExpectRollback()

	now := time.Now()

	repository := warehouse.NewHeartbeatRepository(db)

	err = repository.SaveHeartbeats(context.Background(), []warehouse.Heartbeat{
		warehouse.Heartbeat{
			DeviceID:     42,
			Time:         now,
			IsHealthy:    true,
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "failed to prepare context: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_HeartbeatRepository_SaveHeartbeats_ShouldReturnErrorIfExecFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare("")

	now := time.Now()

	mock.ExpectExec(`^COPY "device_heartbeats" (.+) FROM STDIN*`).
		WithArgs(int64(42), now, true, int64(12), "testing customer").
		WillReturnError(errors.Errorf("boom"))

	mock.ExpectRollback()

	repository := warehouse.NewHeartbeatRepository(db)

	err = repository.SaveHeartbeats(context.Background(), []warehouse.Heartbeat{
		warehouse.Heartbeat{
			DeviceID:     42,
			Time:         now,
			IsHealthy:    true,
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "failed exec: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_HeartbeatRepository_SaveHeartbeats_ShouldReturnErrorOncloseFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare("").WillReturnCloseError(errors.Errorf("boom"))

	now := time.Now()

	mock.ExpectExec(`^COPY "device_heartbeats" (.+) FROM STDIN*`).
		WithArgs(int64(42), now, true, int64(12), "testing customer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectRollback()

	repository := warehouse.NewHeartbeatRepository(db)

	err = repository.SaveHeartbeats(context.Background(), []warehouse.Heartbeat{
		warehouse.Heartbeat{
			DeviceID:     42,
			Time:         now,
			IsHealthy:    true,
			CustomerID:   12,
			CustomerName: "testing customer",
		},
	})
	assert.ErrorContains(t, err, "failed to close: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}
