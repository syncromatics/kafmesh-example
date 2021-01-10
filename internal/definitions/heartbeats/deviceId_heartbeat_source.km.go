// Code generated by kafmesh-gen. DO NOT EDIT.

package heartbeats

import (
	"context"

	"github.com/burdiyan/kafkautil"
	"github.com/lovoo/goka"
	"github.com/pkg/errors"

	"github.com/syncromatics/kafmesh/pkg/runner"

	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"
)

type DeviceIDHeartbeat_Source interface {
	Emit(message DeviceIDHeartbeat_Source_Message) error
	EmitBulk(ctx context.Context, messages []DeviceIDHeartbeat_Source_Message) error
	Delete(key string) error
}

type DeviceIDHeartbeat_Source_impl struct {
	emitter *runner.Emitter
	metrics *runner.Metrics
}

type DeviceIDHeartbeat_Source_Message struct {
	Key string
	Value *deviceId.Heartbeat
}

type impl_DeviceIDHeartbeat_Source_Message struct {
	msg DeviceIDHeartbeat_Source_Message
}

func (m *impl_DeviceIDHeartbeat_Source_Message) Key() string {
	return m.msg.Key
}

func (m *impl_DeviceIDHeartbeat_Source_Message) Value() interface{} {
	return m.msg.Value
}

func New_DeviceIDHeartbeat_Source(service *runner.Service) (*DeviceIDHeartbeat_Source_impl, error) {
	options := service.Options()
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	codec, err := protoWrapper.Codec("kafmesh.deviceId.heartbeat", &deviceId.Heartbeat{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	emitter, err := goka.NewEmitter(brokers,
		goka.Stream("kafmesh.deviceId.heartbeat"),
		codec,
		goka.WithEmitterHasher(kafkautil.MurmurHasher))

	if err != nil {
		return nil, errors.Wrap(err, "failed creating source")
	}

	return &DeviceIDHeartbeat_Source_impl{
		emitter: runner.NewEmitter(emitter),
		metrics: service.Metrics,
	}, nil
}

func (e *DeviceIDHeartbeat_Source_impl) Watch(ctx context.Context) func() error {
	return e.emitter.Watch(ctx)
}

func (e *DeviceIDHeartbeat_Source_impl) Emit(message DeviceIDHeartbeat_Source_Message) error {
	err := e.emitter.Emit(message.Key, message.Value)
	if err != nil {
		e.metrics.SourceError("kafmesh", "heartbeats", "kafmesh.deviceId.heartbeat")
		return err
	}

	e.metrics.SourceHit("kafmesh", "heartbeats", "kafmesh.deviceId.heartbeat", 1)
	return nil
}

func (e *DeviceIDHeartbeat_Source_impl) EmitBulk(ctx context.Context, messages []DeviceIDHeartbeat_Source_Message) error {
	b := []runner.EmitMessage{}
	for _, m := range messages {
		b = append(b, &impl_DeviceIDHeartbeat_Source_Message{msg: m})
	}
	err := e.emitter.EmitBulk(ctx, b)
	if err != nil {
		e.metrics.SourceError("kafmesh", "heartbeats", "kafmesh.deviceId.heartbeat")
		return err
	}

	e.metrics.SourceHit("kafmesh", "heartbeats", "kafmesh.deviceId.heartbeat", len(b))
	return nil
}

func (e *DeviceIDHeartbeat_Source_impl) Delete(key string) error {
	return e.emitter.Emit(key, nil)
}
