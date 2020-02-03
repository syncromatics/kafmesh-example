package warehouse

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

// Heartbeat is the a historical heartbeats update from the device
type Heartbeat struct {
	DeviceID     int64
	Time         time.Time
	IsHealthy    bool
	CustomerID   int64
	CustomerName string
}

// HeartbeatRepository saves and retrieves hearbeats from the database
type HeartbeatRepository struct {
	db *sql.DB
}

// NewHeartbeatRepository creates a heartbeat repository
func NewHeartbeatRepository(db *sql.DB) *HeartbeatRepository {
	return &HeartbeatRepository{db}
}

// SaveHeartbeats saves heartbeats to the warehouse
func (r *HeartbeatRepository) SaveHeartbeats(ctx context.Context, heartbeats []Heartbeat) error {
	txn, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, "begin transaction failed")
	}
	defer txn.Rollback()

	stmt, err := txn.PrepareContext(ctx, pq.CopyIn("device_heartbeats", "device_id", "time", "is_healthy", "customer_id", "customer_name"))
	if err != nil {
		return errors.Wrap(err, "failed to prepare context")
	}

	for _, d := range heartbeats {
		_, err = stmt.ExecContext(ctx, d.DeviceID, d.Time, d.IsHealthy, d.CustomerID, d.CustomerName)
		if err != nil {
			return errors.Wrap(err, "failed exec")
		}
	}

	err = stmt.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close")
	}

	err = txn.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit")
	}

	return nil
}

// LastHeartbeat returns the last heartbeat for a device
func (r *HeartbeatRepository) LastHeartbeat(ctx context.Context, deviceID int64) (*Heartbeat, error) {
	row := r.db.QueryRowContext(ctx, `
	select
		time, is_healthy, customer_id, customer_name
	from
		device_heartbeats
	where
		device_id=$1
	order by
		time desc`, deviceID)

	heartbeat := &Heartbeat{
		DeviceID: deviceID,
	}
	err := row.Scan(&heartbeat.Time, &heartbeat.IsHealthy, &heartbeat.CustomerID, &heartbeat.CustomerName)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to query heartbeats")
	}

	return heartbeat, nil
}
