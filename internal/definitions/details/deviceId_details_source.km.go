// Code generated by kafmesh-gen. DO NOT EDIT.

package details

import (
	"context"

	"github.com/burdiyan/kafkautil"
	"github.com/lovoo/goka"
	"github.com/pkg/errors"

	"github.com/syncromatics/kafmesh/pkg/runner"

	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
)

type DeviceIDDetails_Source interface {
	Emit(message DeviceIDDetails_Source_Message) error
	EmitBulk(ctx context.Context, messages []DeviceIDDetails_Source_Message) error
	Delete(key string) error
}

type DeviceIDDetails_Source_impl struct {
	emitter *runner.Emitter
	metrics *runner.Metrics
}

type DeviceIDDetails_Source_Message struct {
	Key string
	Value *deviceId.Details
}

type impl_DeviceIDDetails_Source_Message struct {
	msg DeviceIDDetails_Source_Message
}

func (m *impl_DeviceIDDetails_Source_Message) Key() string {
	return m.msg.Key
}

func (m *impl_DeviceIDDetails_Source_Message) Value() interface{} {
	return m.msg.Value
}

func New_DeviceIDDetails_Source(service *runner.Service) (*DeviceIDDetails_Source_impl, error) {
	options := service.Options()
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	codec, err := protoWrapper.Codec("kafmesh.deviceId.details", &deviceId.Details{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	emitter, err := goka.NewEmitter(brokers,
		goka.Stream("kafmesh.deviceId.details"),
		codec,
		goka.WithEmitterHasher(kafkautil.MurmurHasher))

	if err != nil {
		return nil, errors.Wrap(err, "failed creating source")
	}

	return &DeviceIDDetails_Source_impl{
		emitter: runner.NewEmitter(emitter),
		metrics: service.Metrics,
	}, nil
}

func (e *DeviceIDDetails_Source_impl) Watch(ctx context.Context) func() error {
	return e.emitter.Watch(ctx)
}

func (e *DeviceIDDetails_Source_impl) Emit(message DeviceIDDetails_Source_Message) error {
	err := e.emitter.Emit(message.Key, message.Value)
	if err != nil {
		e.metrics.SourceError("kafmesh", "details", "kafmesh.deviceId.details")
		return err
	}

	e.metrics.SourceHit("kafmesh", "details", "kafmesh.deviceId.details", 1)
	return nil
}

func (e *DeviceIDDetails_Source_impl) EmitBulk(ctx context.Context, messages []DeviceIDDetails_Source_Message) error {
	b := []runner.EmitMessage{}
	for _, m := range messages {
		b = append(b, &impl_DeviceIDDetails_Source_Message{msg: m})
	}
	err := e.emitter.EmitBulk(ctx, b)
	if err != nil {
		e.metrics.SourceError("kafmesh", "details", "kafmesh.deviceId.details")
		return err
	}

	e.metrics.SourceHit("kafmesh", "details", "kafmesh.deviceId.details", len(b))
	return nil
}

func (e *DeviceIDDetails_Source_impl) Delete(key string) error {
	return e.emitter.Emit(key, nil)
}