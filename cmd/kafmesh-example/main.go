package main

import (
	"context"
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
	"kafmesh-example/internal/implementation/details"
	"kafmesh-example/internal/implementation/heartbeats"
	"kafmesh-example/internal/services"

	"github.com/syncromatics/kafmesh/pkg/runner"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	brokers := []string{"localhost"}
	registry, err := runner.NewRegistry("localhost:443")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service := runner.NewService(brokers, registry)

	configureGatewayService(service, server)
	configureAPIService(service, server)

	setupDetailsComponent(service)

	ctx, cancel := context.WithCancel(context.Background())
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

func configureAPIService(service *runner.Service, server *grpc.Server) {
	e, err := kafmesh.New_DeviceIdCustomer_Emitter(service)
	if err != nil {
		log.Fatal(err)
	}

	v, err := kafmesh.New_DeviceIdCustomer_View(service)
	if err != nil {
		log.Fatal(err)
	}

	api := services.NewAPIService(e, v)

	apiv1.RegisterApiServer(server, api)
}

func setupDetailsComponent(service *runner.Service) {
	warehouseSink := details.NewWarehouseSink()

	err := kafmesh.Register_EnrichedDetailWarehouseSink_Sink(service, warehouseSink, 1*time.Minute, 100)
	if err != nil {
		log.Fatal(err)
	}

	processor := details.NewProcessor()

	err = kafmesh.Register_KafmeshDeviceIdEnrichedDetails_Processor(service, processor)
	if err != nil {
		log.Fatal(err)
	}
}

func setupHeartbeatsComponent(service *runner.Service) {
	warehouseSink := heartbeats.NewWarehouseSink()

	err := kafmesh.Register_EnrichedHeartbeatWarehouseSink_Sink(service, warehouseSink, 1*time.Minute, 100)
	if err != nil {
		log.Fatal(err)
	}

	processor := heartbeats.NewProcessor()

	err = kafmesh.Register_KafmeshDeviceIdEnrichedHeartbeat_Processor(service, processor)
	if err != nil {
		log.Fatal(err)
	}
}
