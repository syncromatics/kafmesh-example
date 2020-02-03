package testing

import (
	"context"
	"fmt"
	apiv1 "kafmesh-example/internal/definitions/models/kafmesh/api/v1"
	historyv1 "kafmesh-example/internal/definitions/models/kafmesh/history/v1"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/pkg/errors"
)

func UpdateDeviceAssignment(api apiv1.ApiClient, device *Device, customer *Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := api.AssignDevice(ctx, &apiv1.AssignDeviceRequest{
		DeviceId:   device.ID,
		CustomerId: customer.ID,
	})
	if err != nil {
		return errors.Wrap(err, "failed to assign device")
	}

	for {
		r, err := api.GetAssignment(ctx, &apiv1.GetAssignmentRequest{
			DeviceId: device.ID,
		})
		if err != nil {
			return errors.Wrap(err, "failed to get assignment")
		}

		if r.CustomerId != nil && r.CustomerId.Value == customer.ID {
			return nil
		}

		select {
		case <-ctx.Done():
			return errors.Errorf("timeout waiting for device assignment")
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func UpdateCustomer(api apiv1.ApiClient, customer *Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if customer.Name == nil {
		return nil
	}

	_, err := api.UpdateCustomerDetails(ctx, &apiv1.UpdateCustomerDetailsRequest{
		CustomerId: customer.ID,
		Name:       *customer.Name,
	})
	if err != nil {
		return errors.Wrap(err, "failed to update customer details")
	}

	for {
		r, err := api.GetCustomerDetails(ctx, &apiv1.GetCustomerDetailsRequest{
			CustomerId: customer.ID,
		})
		if err != nil {
			return errors.Wrap(err, "failed to get customer details")
		}

		if r.Name != nil && r.Name.Value == *customer.Name {
			return nil
		}

		select {
		case <-ctx.Done():
			return errors.Errorf("timeout waiting for customer details update")
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func WaitForExpectedDetails(ctx context.Context, history historyv1.HistoryAPIClient, details *Details, device *Device, customer *Customer) error {
	expected := &historyv1.Details{
		Time:         details.Time,
		Name:         details.Name,
		CustomerId:   customer.ID,
		CustomerName: *customer.Name,
	}
	expected.Time.Nanos = 0

	for {
		r, err := history.LastDetails(ctx, &historyv1.LastDetailsRequest{
			DeviceId: device.ID,
		})
		if err != nil {
			return errors.Wrap(err, "failed to get last details")
		}

		switch v := r.Response.(type) {
		case *historyv1.LastDetailsResponse_ResponseNone:
			goto checkForCancel
		case *historyv1.LastDetailsResponse_ResponseDetails:
			v.ResponseDetails.Time.Nanos = 0

			if proto.Equal(v.ResponseDetails, expected) {
				return nil
			}

			fmt.Printf("'%v' does not equal '%v'\n", v.ResponseDetails, expected)
		}

	checkForCancel:
		select {
		case <-ctx.Done():
			return errors.Errorf("timed out waiting for details")
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func WaitForNoDetails(ctx context.Context, history historyv1.HistoryAPIClient, deviceID int64) error {
	for {
		r, err := history.LastDetails(ctx, &historyv1.LastDetailsRequest{
			DeviceId: deviceID,
		})

		if err != nil && strings.Contains(err.Error(), "DeadlineExceeded") {
			return nil
		}

		if err != nil {
			return errors.Wrap(err, "failed to get last details")
		}

		switch v := r.Response.(type) {
		case *historyv1.LastDetailsResponse_ResponseNone:
			goto checkForCancel
		case *historyv1.LastDetailsResponse_ResponseDetails:
			return errors.Errorf("expected no details and got '%v", v.ResponseDetails)
		}

	checkForCancel:
		select {
		case <-ctx.Done():
			return nil
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func WaitForExpectedHeartbeat(ctx context.Context, history historyv1.HistoryAPIClient, heartbeat *Heartbeat, device *Device, customer *Customer) error {
	expected := &historyv1.Heartbeat{
		Time:         heartbeat.Time,
		IsHealthy:    heartbeat.IsHealthy,
		CustomerId:   customer.ID,
		CustomerName: *customer.Name,
	}
	expected.Time.Nanos = 0

	for {
		r, err := history.LastHeartbeat(ctx, &historyv1.LastHeartbeatRequest{
			DeviceId: device.ID,
		})
		if err != nil {
			return errors.Wrap(err, "failed to get last heartbeat")
		}

		switch v := r.Response.(type) {
		case *historyv1.LastHeartbeatResponse_ResponseNone:
			goto checkForCancel
		case *historyv1.LastHeartbeatResponse_ResponseHeartbeat:
			v.ResponseHeartbeat.Time.Nanos = 0

			if proto.Equal(v.ResponseHeartbeat, expected) {
				return nil
			}

			fmt.Printf("'%v' does not equal '%v'\n", v.ResponseHeartbeat, expected)
		}

	checkForCancel:
		select {
		case <-ctx.Done():
			return errors.Errorf("timed out waiting for heartbeats")
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func WaitForNoHeartbeat(ctx context.Context, history historyv1.HistoryAPIClient, deviceID int64) error {
	for {
		r, err := history.LastHeartbeat(ctx, &historyv1.LastHeartbeatRequest{
			DeviceId: deviceID,
		})

		if err != nil && strings.Contains(err.Error(), "DeadlineExceeded") {
			return nil
		}

		if err != nil {
			return errors.Wrap(err, "failed to get last details")
		}

		switch v := r.Response.(type) {
		case *historyv1.LastHeartbeatResponse_ResponseNone:
			goto checkForCancel
		case *historyv1.LastHeartbeatResponse_ResponseHeartbeat:
			return errors.Errorf("expected no heartbeat and got '%v", v.ResponseHeartbeat)
		}

	checkForCancel:
		select {
		case <-ctx.Done():
			return nil
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
