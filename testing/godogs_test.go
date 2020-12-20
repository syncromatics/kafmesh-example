package testing

import (
	"context"
	"log"
	"time"

	apiv1 "kafmesh-example/internal/definitions/models/kafmesh/api/v1"
	gatewayv1 "kafmesh-example/internal/definitions/models/kafmesh/gateway/v1"
	historyv1 "kafmesh-example/internal/definitions/models/kafmesh/history/v1"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/syncromatics/kafmesh/pkg/runner"
	"google.golang.org/grpc"
)

var (
	gateway gatewayv1.GatewayAPIClient
	api     apiv1.ApiClient
	history historyv1.HistoryAPIClient

	details testDetails
)

type testDetails struct {
	device    *Device
	customer  *Customer
	details   *Details
	heartbeat *Heartbeat
}

func thereIsADevice(arg1 *messages.PickleStepArgument_PickleTable) error {
	d, err := NewDevice(arg1)
	if err != nil {
		return err
	}

	details.device = d

	return nil
}

func itIsAssignedToCustomer(arg1 *messages.PickleStepArgument_PickleTable) error {
	c, err := NewCustomer(arg1)
	if err != nil {
		return err
	}
	details.customer = c

	err = UpdateCustomer(api, c)
	if err != nil {
		return err
	}

	err = UpdateDeviceAssignment(api, details.device, details.customer)
	if err != nil {
		return err
	}
	return nil
}

func itSendsDetailsToTheGateway(arg1 *messages.PickleStepArgument_PickleTable) error {
	d, err := NewDetails(arg1)
	if err != nil {
		return err
	}
	details.details = d

	_, err = gateway.Details(context.Background(), &gatewayv1.DetailsRequest{
		DeviceId: details.device.ID,
		Name:     d.Name,
		Time:     d.Time,
	})
	if err != nil {
		return err
	}

	return nil
}

func detailsShouldBeSavedToTheWarehouse() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := WaitForExpectedDetails(ctx, history, details.details, details.device, details.customer)
	if err != nil {
		return err
	}

	return nil
}

func theDetailsShouldNotBeSavedToTheWarehouse() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return WaitForNoDetails(ctx, history, details.device.ID)
}

func itSendsAHeartbeatToTheGateway(arg1 *messages.PickleStepArgument_PickleTable) error {
	h, err := NewHeartbeat(arg1)
	if err != nil {
		return err
	}
	details.heartbeat = h

	_, err = gateway.Heartbeat(context.Background(), &gatewayv1.HeartbeatRequest{
		DeviceId:  details.device.ID,
		Time:      h.Time,
		IsHealthy: h.IsHealthy,
	})
	if err != nil {
		return err
	}

	return nil
}

func theHeartbeatShouldBeSavedToTheWarehouse() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := WaitForExpectedHeartbeat(ctx, history, details.heartbeat, details.device, details.customer)
	if err != nil {
		return err
	}
	return nil
}

func theHeartbeatShouldNotBeSavedToTheWarehouse() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return WaitForNoHeartbeat(ctx, history, details.device.ID)
}

func FeatureContext(s *godog.Suite) {
	s.BeforeSuite(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := runner.WaitTillServiceIsRunning(ctx, "kafmesh-example:8888")
		if err != nil {
			log.Fatal(err)
		}

		con, err := grpc.Dial("kafmesh-example:8888", grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}

		gateway = gatewayv1.NewGatewayAPIClient(con)
		api = apiv1.NewApiClient(con)
		history = historyv1.NewHistoryAPIClient(con)
	})

	s.BeforeScenario(func(*messages.Pickle) {
		details = testDetails{}
	})

	s.Step(`^there is a device$`, thereIsADevice)
	s.Step(`^it is assigned to customer$`, itIsAssignedToCustomer)

	s.Step(`^it sends details to the gateway$`, itSendsDetailsToTheGateway)
	s.Step(`^details should be saved to the warehouse$`, detailsShouldBeSavedToTheWarehouse)
	s.Step(`^the details should not be saved to the warehouse$`, theDetailsShouldNotBeSavedToTheWarehouse)

	s.Step(`^it sends a heartbeat to the gateway$`, itSendsAHeartbeatToTheGateway)
	s.Step(`^the heartbeat should be saved to the warehouse$`, theHeartbeatShouldBeSavedToTheWarehouse)
	s.Step(`^the heartbeat should not be saved to the warehouse$`, theHeartbeatShouldNotBeSavedToTheWarehouse)
}
