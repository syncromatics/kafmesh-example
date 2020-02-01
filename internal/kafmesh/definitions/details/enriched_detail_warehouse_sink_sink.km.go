// Code generated by kafmesh-gen. DO NOT EDIT.

package details

import (
	"context"
	"time"

	"github.com/lovoo/goka"
	"github.com/pkg/errors"

	"github.com/syncromatics/kafmesh/pkg/runner"

	deviceId "kafmesh-example/internal/kafmesh/definitions/models/kafmesh/deviceId"
)

type EnrichedDetailWarehouseSink_Sink interface {
	Flush() error
	Collect(ctx runner.MessageContext, key string, msg *deviceId.EnrichedDetails) error
}

type impl_EnrichedDetailWarehouseSink_Sink struct {
	sink EnrichedDetailWarehouseSink_Sink
	codec goka.Codec
	group string
	topic string
	maxBufferSize int
	interval time.Duration
}

func (s *impl_EnrichedDetailWarehouseSink_Sink) Codec() goka.Codec {
	return s.codec
}

func (s *impl_EnrichedDetailWarehouseSink_Sink) Group() string {
	return s.group
}

func (s *impl_EnrichedDetailWarehouseSink_Sink) Topic() string {
	return s.topic
}

func (s *impl_EnrichedDetailWarehouseSink_Sink) MaxBufferSize() int {
	return s.maxBufferSize
}

func (s *impl_EnrichedDetailWarehouseSink_Sink) Interval() time.Duration {
	return s.interval
}

func (s *impl_EnrichedDetailWarehouseSink_Sink) Flush() error {
	return s.sink.Flush()
}

func (s *impl_EnrichedDetailWarehouseSink_Sink) Collect(ctx runner.MessageContext, key string, msg interface{}) error {
	m, ok := msg.(*deviceId.EnrichedDetails)
	if !ok {
		return errors.Errorf("expecting message of type '*deviceId.EnrichedDetails' got type '%t'", msg)
	}

	return s.sink.Collect(ctx, key, m)
}

func Register_EnrichedDetailWarehouseSink_Sink(options runner.ServiceOptions, sink EnrichedDetailWarehouseSink_Sink, interval time.Duration, maxBufferSize int) (func(ctx context.Context) func() error, error) {
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	codec, err := protoWrapper.Codec("kafmesh.deviceId.enrichedDetails", &deviceId.EnrichedDetails{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	d := &impl_EnrichedDetailWarehouseSink_Sink{
		sink: sink,
		codec: codec,
		group: "enricheddetailwarehousesink-sink",
		topic: "kafmesh.deviceId.enrichedDetails",
		maxBufferSize: maxBufferSize,
		interval: interval,
	}

	s := runner.NewSinkRunner(d, brokers)

	return func(ctx context.Context) func() error {
		return s.Run(ctx)
	}, nil
}
