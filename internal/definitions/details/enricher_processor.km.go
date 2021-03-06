// Code generated by kafmesh-gen. DO NOT EDIT.

package details

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

type Enricher_ProcessorContext interface {
	Key() string
	Timestamp() time.Time
	Lookup_CustomerIDDetails(key string) *m1.Details
	Output_DeviceIDEnrichedDetails(key string, message *m0.EnrichedDetails)
	SaveState(state *m0.EnrichedDetailsState)
	State() *m0.EnrichedDetailsState
}

type Enricher_Processor interface {
	HandleDeviceIDDetails(ctx Enricher_ProcessorContext, message *m0.Details) error
	HandleDeviceIDCustomer(ctx Enricher_ProcessorContext, message *m0.Customer) error
}

type Enricher_ProcessorContext_Impl struct {
	ctx              goka.Context
	processorContext *runner.ProcessorContext
}

func new_Enricher_ProcessorContext_Impl(ctx goka.Context, pc *runner.ProcessorContext) *Enricher_ProcessorContext_Impl {
	return &Enricher_ProcessorContext_Impl{ctx, pc}
}

func (c *Enricher_ProcessorContext_Impl) Key() string {
	return c.ctx.Key()
}

func (c *Enricher_ProcessorContext_Impl) Timestamp() time.Time {
	return c.ctx.Timestamp()
}

func (c *Enricher_ProcessorContext_Impl) Lookup_CustomerIDDetails(key string) *m1.Details {
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

func (c *Enricher_ProcessorContext_Impl) Output_DeviceIDEnrichedDetails(key string, message *m0.EnrichedDetails) {
	value, _ := json.Marshal(message)
	c.processorContext.Output("kafmesh.deviceId.enrichedDetails", "deviceId.enrichedDetails", key, string(value))
	c.ctx.Emit("kafmesh.deviceId.enrichedDetails", key, message)
}

func (c *Enricher_ProcessorContext_Impl) SaveState(state *m0.EnrichedDetailsState) {
	value, _ := json.Marshal(state)
	c.processorContext.SetState("kafmesh.details.enricher-table", "deviceId.enrichedDetailsState", string(value))

	c.ctx.SetValue(state)
}

func (c *Enricher_ProcessorContext_Impl) State() *m0.EnrichedDetailsState {
	v := c.ctx.Value()
	var m *m0.EnrichedDetailsState
	if v == nil {
		m = &m0.EnrichedDetailsState{}
	} else {
		m = v.(*m0.EnrichedDetailsState)
	}

	value, _ := json.Marshal(m)
	c.processorContext.GetState("kafmesh.details.enricher-table", "deviceId.enrichedDetailsState", string(value))

	return m
}

func Register_Enricher_Processor(service *runner.Service, impl Enricher_Processor) (func(context.Context) func() error, error) {
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

	path := filepath.Join("/tmp/storage", "processor", "kafmesh.details.enricher")

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create processor db directory")
	}

	builder := storage.BuilderWithOptions(path, opts)


	c0, err := protoWrapper.Codec("kafmesh.deviceId.details", &m0.Details{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c1, err := protoWrapper.Codec("kafmesh.deviceId.customer", &m0.Customer{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c2, err := protoWrapper.Codec("kafmesh.customerId.details", &m1.Details{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c3, err := protoWrapper.Codec("kafmesh.deviceId.enrichedDetails", &m0.EnrichedDetails{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	c4, err := protoWrapper.Codec("kafmesh.details.enricher-table", &m0.EnrichedDetailsState{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	edges := []goka.Edge{
		goka.Input(goka.Stream("kafmesh.deviceId.details"), c0, func(ctx goka.Context, m interface{}) {
			msg := m.(*m0.Details)

			pc := service.ProcessorContext(ctx.Context(), "details", "enricher", ctx.Key())
			defer pc.Finish()

			v, err := json.Marshal(msg)
			if err != nil {
				ctx.Fail(err)
			}
			pc.Input("kafmesh.deviceId.details", "deviceId.details", string(v))

			w := new_Enricher_ProcessorContext_Impl(ctx, pc)
			err = impl.HandleDeviceIDDetails(w, msg)
			if err != nil {
				ctx.Fail(err)
			}
		}),
		goka.Input(goka.Stream("kafmesh.deviceId.customer"), c1, func(ctx goka.Context, m interface{}) {
			msg := m.(*m0.Customer)

			pc := service.ProcessorContext(ctx.Context(), "details", "enricher", ctx.Key())
			defer pc.Finish()

			v, err := json.Marshal(msg)
			if err != nil {
				ctx.Fail(err)
			}
			pc.Input("kafmesh.deviceId.customer", "deviceId.customer", string(v))

			w := new_Enricher_ProcessorContext_Impl(ctx, pc)
			err = impl.HandleDeviceIDCustomer(w, msg)
			if err != nil {
				ctx.Fail(err)
			}
		}),
		goka.Lookup(goka.Table("kafmesh.customerId.details"), c2),
		goka.Output(goka.Stream("kafmesh.deviceId.enrichedDetails"), c3),
		goka.Persist(c4),
	}
	group := goka.DefineGroup(goka.Group("kafmesh.details.enricher"), edges...)

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
