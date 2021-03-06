// Code generated by kafmesh-gen. DO NOT EDIT.

package assignments

import (
	"context"

	"github.com/burdiyan/kafkautil"
	"github.com/lovoo/goka"
	"github.com/pkg/errors"

	"github.com/syncromatics/kafmesh/pkg/runner"

	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
)

type CustomerIDDetails_Source interface {
	Emit(message CustomerIDDetails_Source_Message) error
	EmitBulk(ctx context.Context, messages []CustomerIDDetails_Source_Message) error
	Delete(key string) error
}

type CustomerIDDetails_Source_impl struct {
	emitter *runner.Emitter
	metrics *runner.Metrics
}

type CustomerIDDetails_Source_Message struct {
	Key string
	Value *customerId.Details
}

type impl_CustomerIDDetails_Source_Message struct {
	msg CustomerIDDetails_Source_Message
}

func (m *impl_CustomerIDDetails_Source_Message) Key() string {
	return m.msg.Key
}

func (m *impl_CustomerIDDetails_Source_Message) Value() interface{} {
	return m.msg.Value
}

func New_CustomerIDDetails_Source(service *runner.Service) (*CustomerIDDetails_Source_impl, error) {
	options := service.Options()
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	codec, err := protoWrapper.Codec("kafmesh.customerId.details", &customerId.Details{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	emitter, err := goka.NewEmitter(brokers,
		goka.Stream("kafmesh.customerId.details"),
		codec,
		goka.WithEmitterHasher(kafkautil.MurmurHasher))

	if err != nil {
		return nil, errors.Wrap(err, "failed creating source")
	}

	return &CustomerIDDetails_Source_impl{
		emitter: runner.NewEmitter(emitter),
		metrics: service.Metrics,
	}, nil
}

func (e *CustomerIDDetails_Source_impl) Watch(ctx context.Context) func() error {
	return e.emitter.Watch(ctx)
}

func (e *CustomerIDDetails_Source_impl) Emit(message CustomerIDDetails_Source_Message) error {
	err := e.emitter.Emit(message.Key, message.Value)
	if err != nil {
		e.metrics.SourceError("kafmesh", "assignments", "kafmesh.customerId.details")
		return err
	}

	e.metrics.SourceHit("kafmesh", "assignments", "kafmesh.customerId.details", 1)
	return nil
}

func (e *CustomerIDDetails_Source_impl) EmitBulk(ctx context.Context, messages []CustomerIDDetails_Source_Message) error {
	b := []runner.EmitMessage{}
	for _, m := range messages {
		b = append(b, &impl_CustomerIDDetails_Source_Message{msg: m})
	}
	err := e.emitter.EmitBulk(ctx, b)
	if err != nil {
		e.metrics.SourceError("kafmesh", "assignments", "kafmesh.customerId.details")
		return err
	}

	e.metrics.SourceHit("kafmesh", "assignments", "kafmesh.customerId.details", len(b))
	return nil
}

func (e *CustomerIDDetails_Source_impl) Delete(key string) error {
	return e.emitter.Emit(key, nil)
}
