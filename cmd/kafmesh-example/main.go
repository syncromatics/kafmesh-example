package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	kafmesh "kafmesh-example/internal/definitions"
	apiv1 "kafmesh-example/internal/definitions/models/kafmesh/api/v1"
	gatewayv1 "kafmesh-example/internal/definitions/models/kafmesh/gateway/v1"
	historyv1 "kafmesh-example/internal/definitions/models/kafmesh/history/v1"
	"kafmesh-example/internal/implementation/assignments"
	"kafmesh-example/internal/implementation/details"
	"kafmesh-example/internal/implementation/heartbeats"
	"kafmesh-example/internal/services"
	"kafmesh-example/internal/warehouse"

	"github.com/pkg/errors"
	"github.com/syncromatics/kafmesh/pkg/runner"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	_ "kafmesh-example/internal/migrations/statik"

	_ "github.com/lib/pq"
)

func main() {
	settings, err := getSettingsFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	registry, err := runner.NewRegistry(settings.Registry)
	if err != nil {
		log.Fatal(err)
	}

	err = registry.WaitForRegistryToBeReady(30 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = settings.DatabaseSettings.WaitForDatabaseToBeOnline(30)
	if err != nil {
		log.Fatal(err)
	}

	err = settings.DatabaseSettings.MigrateUpWithStatik("/")
	if err != nil {
		log.Fatal(err)
	}

	db, err := settings.DatabaseSettings.EnsureDatabaseExistsAndGetConnection()
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := runner.NewService(settings.Brokers, registry, server)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	err = service.ConfigureKafka(ctx, kafmesh.ConfigureTopics)
	if err != nil {
		log.Fatal(err)
	}
	cancel()

	configureGatewayService(service, server)
	configureAPIService(service, server, db)
	configureHistoryService(server, db)

	configureAssignments(service, db)
	setupDetailsComponent(service, db)
	setupHeartbeatsComponent(service, db)

	ctx, cancel = context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(service.Run(ctx))
	go func() {
		select {
		case <-ctx.Done():
			go func() {
				time.Sleep(10 * time.Second)
				server.Stop()
			}()
			server.GracefulStop()
			return
		}
	}()
	grp.Go(func() error {
		lis, err := net.Listen("tcp", "0.0.0.0:8888")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		err = server.Serve(lis)
		if err != nil {
			return err
		}
		return nil
	})

	eventChan := make(chan os.Signal)
	signal.Notify(eventChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("server started...")

	select {
	case <-eventChan:
	case <-ctx.Done():
	}

	fmt.Println("server stopping...")

	cancel()

	if err := grp.Wait(); err != nil {
		log.Fatal(errors.Wrap(err, "service failed"))
	}
}

func configureAssignments(service *runner.Service, db *sql.DB) {
	repository := warehouse.NewCustomerDetailsRepository(db)
	synchronizer := assignments.NewCustomerDetailsSynchronizer(repository)

	err := kafmesh.Register_CustomerIdDetails_Synchronizer(service, synchronizer, 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

func configureGatewayService(service *runner.Service, server *grpc.Server) {
	e, err := kafmesh.New_DeviceIdDetails_Emitter(service)
	if err != nil {
		log.Fatal(err)
	}

	he, err := kafmesh.New_DeviceIdHeartbeat_Emitter(service)
	if err != nil {
		log.Fatal(err)
	}

	gateway := services.NewGatewayService(e, he)

	gatewayv1.RegisterGatewayAPIServer(server, gateway)
}

func configureAPIService(service *runner.Service, server *grpc.Server, db *sql.DB) {
	e, err := kafmesh.New_DeviceIdCustomer_Emitter(service)
	if err != nil {
		log.Fatal(err)
	}

	v, err := kafmesh.New_DeviceIdCustomer_View(service)
	if err != nil {
		log.Fatal(err)
	}

	de, err := kafmesh.New_CustomerIdDetails_Emitter(service)
	if err != nil {
		log.Fatal(err)
	}

	dv, err := kafmesh.New_CustomerIdDetails_View(service)
	if err != nil {
		log.Fatal(err)
	}

	repository := warehouse.NewCustomerDetailsRepository(db)

	api := services.NewAPIService(e, v, de, dv, repository)

	apiv1.RegisterApiServer(server, api)
}

func configureHistoryService(server *grpc.Server, db *sql.DB) {
	detailsRepository := warehouse.NewDetailsRepository(db)
	heartbeatsRepository := warehouse.NewHeartbeatRepository(db)

	history := services.NewHistoryAPI(detailsRepository, heartbeatsRepository)

	historyv1.RegisterHistoryAPIServer(server, history)
}

func setupDetailsComponent(service *runner.Service, db *sql.DB) {
	repository := warehouse.NewDetailsRepository(db)
	warehouseSink := details.NewWarehouseSink(repository)

	err := kafmesh.Register_EnrichedDetailWarehouseSink_Sink(service, warehouseSink, 5*time.Second, 100)
	if err != nil {
		log.Fatal(err)
	}

	processor := details.NewProcessor()

	err = kafmesh.Register_KafmeshDeviceIdEnrichedDetails_Processor(service, processor)
	if err != nil {
		log.Fatal(err)
	}
}

func setupHeartbeatsComponent(service *runner.Service, db *sql.DB) {
	repository := warehouse.NewHeartbeatRepository(db)
	warehouseSink := heartbeats.NewWarehouseSink(repository)

	err := kafmesh.Register_EnrichedHeartbeatWarehouseSink_Sink(service, warehouseSink, 5*time.Second, 100)
	if err != nil {
		log.Fatal(err)
	}

	processor := heartbeats.NewProcessor()

	err = kafmesh.Register_KafmeshDeviceIdEnrichedHeartbeat_Processor(service, processor)
	if err != nil {
		log.Fatal(err)
	}
}
