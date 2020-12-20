// Code generated by kafmesh-gen. DO NOT EDIT.

package heartbeats

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/Shopify/sarama"
	"github.com/burdiyan/kafkautil"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/storage"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/syncromatics/kafmesh/pkg/runner"

	m0 "kafmesh-example/internal/definitions/models/kafmesh/deviceId"
	m1 "kafmesh-example/internal/definitions/models/kafmesh/customerId"
)

type HeartbeatEnricher_ProcessorContext interface {
	Key() string
	Timestamp() time.Time
	Lookup_CustomerIDDetails(key string) *m1.Details
	Join_DeviceIDCustomer() *m0.Customer
	Output_DeviceIDEnrichedHeartbeat(key string, message *m0.EnrichedHeartbeat)
}

type HeartbeatEnricher_Processor interface {
	HandleDeviceIDHeartbeat(ctx HeartbeatEnricher_ProcessorContext, message *m0.Heartbeat) error
}

type HeartbeatEnricher_ProcessorContext_Impl struct {
	ctx              goka.Context
	processorContext *runner.ProcessorContext
}

func new_HeartbeatEnricher_ProcessorContext_Impl(ctx goka.Context, pc *runner.ProcessorContext) *HeartbeatEnricher_ProcessorContext_Impl {
	return &HeartbeatEnricher_ProcessorContext_Impl{ctx, pc}
}

func (c *HeartbeatEnricher_ProcessorContext_Impl) Key() string {
	return c.ctx.Key()
}

func (c *HeartbeatEnricher_ProcessorContext_Impl) Timestamp() time.Time {
	return c.ctx.Timestamp()
}

func (c *HeartbeatEnricher_ProcessorContext_Impl) Lookup_CustomerIDDetails(key string) *m1.Details {
	v := c.ctx.Lookup("kafmesh.customerId.details", key)
	if v == nil {
		c.processorContext.Lookup("kafmesh.customerId.details", "customerId.details", key, "")
		return nil
	}

	m := v.(*m1.Details)
	value, _ := json.Marshal(m)
	c.processorContext.Lookup("kafmesh.customerId.details", "customerId.details", key, string(value))

	return m
}

func (c *HeartbeatEnricher_ProcessorContext_Impl) Join_DeviceIDCustomer() *m0.Customer {
	v := c.ctx.Join("kafmesh.deviceId.customer")
	if v == nil {
		c.processorContext.Join("kafmesh.deviceId.customer", "deviceId.customer", "")
		return nil
	}

	m := v.(*m0.Customer)
	value, _ := json.Marshal(m)
	c.processorContext.Join("kafmesh.deviceId.customer", "deviceId.customer", string(value))

	return m
}

func (c *HeartbeatEnricher_ProcessorContext_Impl) Output_DeviceIDEnrichedHeartbeat(key string, message *m0.EnrichedHeartbeat) {
	value, _ := json.Marshal(message)
	c.processorContext.Output("kafmesh.deviceId.enrichedHeartbeat", "deviceId.enrichedHeartbeat", key, string(value))
	c.ctx.Emit("kafmesh.deviceId.enrichedHeartbeat", key, message)
}

func Register_HeartbeatEnricher_Processor(service *runner.Service, impl HeartbeatEnricher_Processor) (func(context.Context) func() error, error) {
	options := service.Options()
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second

	opts := &opt.Options{
		BlockCacheCapacity: opt.MiB * 1,
		WriteBuffer:        opt.MiB * 1,
	}

	path := filepath.Join("/tmp/storage", "processor", "deviceId.enrichedHeartbeat")

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create processor db directory")
	}

	builder := storage.BuilderWithOptions(path, opts)


	c0, err := protoWrapper.Codec("kafmesh.deviceId.heartbeat", &m0.Heartbeat{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c1, err := protoWrapper.Codec("kafmesh.customerId.details", &m1.Details{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c2, err := protoWrapper.Codec("kafmesh.deviceId.customer", &m0.Customer{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c3, err := protoWrapper.Codec("kafmesh.deviceId.enrichedHeartbeat", &m0.EnrichedHeartbeat{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	edges := []goka.Edge{
		goka.Input(goka.Stream("kafmesh.deviceId.heartbeat"), c0, func(ctx goka.Context, m interface{}) {
			msg := m.(*m0.Heartbeat)

			pc := service.ProcessorContext(ctx.Context(), "heartbeats", "heartbeat enricher", ctx.Key())
			defer pc.Finish()

			v, err := json.Marshal(msg)
			if err != nil {
				ctx.Fail(err)
			}
			pc.Input("kafmesh.deviceId.heartbeat", "deviceId.heartbeat", string(v))

			w := new_HeartbeatEnricher_ProcessorContext_Impl(ctx, pc)
			err = impl.HandleDeviceIDHeartbeat(w, msg)
			if err != nil {
				ctx.Fail(err)
			}
		}),
		goka.Lookup(goka.Table("kafmesh.customerId.details"), c1),
		goka.Join(goka.Table("kafmesh.deviceId.customer"), c2),
		goka.Output(goka.Stream("kafmesh.deviceId.enrichedHeartbeat"), c3),
	}
	group := goka.DefineGroup(goka.Group("deviceId.enrichedHeartbeat"), edges...)

	processor, err := goka.NewProcessor(brokers,
		group,
		goka.WithConsumerGroupBuilder(goka.ConsumerGroupBuilderWithConfig(config)),
		goka.WithStorageBuilder(builder),
		goka.WithHasher(kafkautil.MurmurHasher))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create goka processor")
	}

	return func(ctx context.Context) func() error {
		return func() error {
			err := processor.Run(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to run goka processor")
			}

			return nil
		}
	}, nil
}
