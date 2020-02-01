// Code generated by kafmesh-gen. DO NOT EDIT.

package details

import (
	"context"

	"github.com/burdiyan/kafkautil"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/kafka"
	"github.com/lovoo/goka/storage"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/syncromatics/kafmesh/pkg/runner"

	m0 "kafmesh-example/internal/kafmesh/definitions/models/kafmesh/deviceId"
	m1 "kafmesh-example/internal/kafmesh/definitions/models/kafmesh/customerId"
)

type KafmeshDeviceIdEnrichedDetails_ProcessorContext interface {
	Lookup_CustomerIdDetails(key string) *m1.Details
	Output_DeviceIdEnrichedDetails(key string, message *m0.EnrichedDetails)
	SaveState(state *m0.EnrichedDetailsState)
	State() *m0.EnrichedDetailsState
}

type KafmeshDeviceIdEnrichedDetails_Processor interface {
	HandleInput_DeviceIdDetails(ctx KafmeshDeviceIdEnrichedDetails_ProcessorContext, message *m0.Details) error
	HandleInput_DeviceIdCustomer(ctx KafmeshDeviceIdEnrichedDetails_ProcessorContext, message *m0.Customer) error
}

type KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl struct {
	ctx goka.Context
}

func new_KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl(ctx goka.Context) *KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl {
	return &KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl{ctx}
}

func (c *KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl) Lookup_CustomerIdDetails(key string) *m1.Details {
	v := c.ctx.Lookup("kafmesh.customerId.details", key)
	return v.(*m1.Details)
}

func (c *KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl) Output_DeviceIdEnrichedDetails(key string, message *m0.EnrichedDetails) {
	c.ctx.Emit("kafmesh.deviceId.enrichedDetails", key, message)
}

func (c *KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl) SaveState(state *m0.EnrichedDetailsState) {
	c.ctx.SetValue(state)
}

func (c *KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl) State() *m0.EnrichedDetailsState {
	v := c.ctx.Value()
	t := v.(*m0.EnrichedDetailsState)
	return t
}

func Register_KafmeshDeviceIdEnrichedDetails_Processor(options runner.ServiceOptions, service KafmeshDeviceIdEnrichedDetails_Processor) (func(context.Context) func() error, error) {
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	config := kafka.NewConfig()
	config.Consumer.Offsets.Initial = kafka.OffsetOldest

	opts := &opt.Options{
		BlockCacheCapacity: opt.MiB * 1,
		WriteBuffer:        opt.MiB * 1,
	}

	builder := storage.BuilderWithOptions("/tmp/storage", opts)

	c0, err := protoWrapper.Codec("kafmesh.deviceId.details", &m0.Details{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c1, err := protoWrapper.Codec("kafmesh.deviceId.customer", &m0.Customer{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c2, err := protoWrapper.Codec("kafmesh.deviceId.enrichedDetails", &m0.EnrichedDetails{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c3, err := protoWrapper.Codec("kafmesh.deviceId.enrichedDetails-table", &m0.EnrichedDetailsState{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	edges := []goka.Edge{
		goka.Input(goka.Stream("kafmesh.deviceId.details"), c0, func(ctx goka.Context, m interface{}) {
			msg := m.(*m0.Details)
			w := new_KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl(ctx)
			err := service.HandleInput_DeviceIdDetails(w, msg)
			if err != nil {
				ctx.Fail(err)
			}
		}),
		goka.Input(goka.Stream("kafmesh.deviceId.customer"), c1, func(ctx goka.Context, m interface{}) {
			msg := m.(*m0.Customer)
			w := new_KafmeshDeviceIdEnrichedDetails_ProcessorContext_Impl(ctx)
			err := service.HandleInput_DeviceIdCustomer(w, msg)
			if err != nil {
				ctx.Fail(err)
			}
		}),
		goka.Lookup(goka.Table("kafmesh.customerId.details"), c0),
		goka.Output(goka.Stream("kafmesh.deviceId.enrichedDetails"), c2),
		goka.Persist(c3),
	}
	group := goka.DefineGroup(goka.Group("kafmesh.deviceId.enrichedDetails"), edges...)

	processor, err := goka.NewProcessor(brokers,
		group,
		goka.WithConsumerBuilder(kafka.ConsumerBuilderWithConfig(config)),
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
