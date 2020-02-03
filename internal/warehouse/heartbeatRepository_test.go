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

	mock.ExpectCommit()

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

func Test_HeartbeatRepository_SaveHeartbeats_ShouldReturnErrorIfCommitFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectBegin()

	mock.ExpectPrepare("")

	now := time.Now()

	mock.ExpectExec(`^COPY "device_heartbeats" (.+) FROM STDIN*`).
		WithArgs(int64(42), now, true, int64(12), "testing customer").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit().WillReturnError(errors.Errorf("boom"))

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
	assert.ErrorContains(t, err, "failed to commit: boom")

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

func Test_HeartbeatRepository_LastHeartbeat(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	now := time.Now()
	rows := sqlmock.NewRows([]string{"time", "is_healthy", "customer_id", "customer_name"}).
		AddRow(now, true, 65, "test customer")

	mock.ExpectQuery("^select time, is_healthy, customer_id, customer_name from device_heartbeats where device_id=\\$1 order by	time desc*").
		WithArgs(4).
		WillReturnRows(rows)

	repository := warehouse.NewHeartbeatRepository(db)
	last, err := repository.LastHeartbeat(context.Background(), 4)
	assert.NilError(t, err)

	assert.DeepEqual(t, last, &warehouse.Heartbeat{
		DeviceID:     4,
		Time:         now,
		IsHealthy:    true,
		CustomerID:   65,
		CustomerName: "test customer",
	})

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_HeartbeatRepository_LastHeartbeat_ShouldReturnNilWithNoResult(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectQuery("^select time, is_healthy, customer_id, customer_name from device_heartbeats where device_id=\\$1 order by	time desc*").
		WithArgs(5).
		WillReturnRows(sqlmock.NewRows([]string{"time", "is_healthy", "customer_id", "customer_name"}))

	repository := warehouse.NewHeartbeatRepository(db)
	last, err := repository.LastHeartbeat(context.Background(), 5)
	assert.NilError(t, err)

	assert.Assert(t, last == nil)

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}

func Test_HeartbeatRepository_LastHeartbeat_ShouldReturnErrorIfQueryFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	mock.ExpectQuery("^select time, is_healthy, customer_id, customer_name from device_heartbeats where device_id=\\$1 order by	time desc*").
		WithArgs(7).
		WillReturnError(errors.Errorf("boom"))

	repository := warehouse.NewHeartbeatRepository(db)
	_, err = repository.LastHeartbeat(context.Background(), 7)
	assert.ErrorContains(t, err, "failed to query heartbeats: boom")

	err = mock.ExpectationsWereMet()
	assert.NilError(t, err)
}
