// Code generated by kafmesh-gen. DO NOT EDIT.

package heartbeats

import (
	"context"

	"github.com/burdiyan/kafkautil"
	"github.com/lovoo/goka"
	"github.com/pkg/errors"

	"github.com/syncromatics/kafmesh/pkg/runner"

	deviceId "kafmesh-example/internal/definitions/models/kafmesh/deviceId"
)

type DeviceIdHeartbeat_Emitter interface {
	Emit(message DeviceIdHeartbeat_Emitter_Message) error
	EmitBulk(ctx context.Context, messages []DeviceIdHeartbeat_Emitter_Message) error
}

type DeviceIdHeartbeat_Emitter_impl struct {
	emitter *runner.Emitter
}

type DeviceIdHeartbeat_Emitter_Message struct {
	Key string
	Value *deviceId.Heartbeat
}

type impl_DeviceIdHeartbeat_Emitter_Message struct {
	msg DeviceIdHeartbeat_Emitter_Message
}

func (m *impl_DeviceIdHeartbeat_Emitter_Message) Key() string {
	return m.msg.Key
}

func (m *impl_DeviceIdHeartbeat_Emitter_Message) Value() interface{} {
	return m.msg.Value
}

func New_DeviceIdHeartbeat_Emitter(options runner.ServiceOptions) (*DeviceIdHeartbeat_Emitter_impl, error) {
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
		return nil, errors.Wrap(err, "failed creating emitter")
	}

	return &DeviceIdHeartbeat_Emitter_impl{
		emitter: runner.NewEmitter(emitter),
	}, nil
}

func (e *DeviceIdHeartbeat_Emitter_impl) Watch(ctx context.Context) func() error {
	return e.emitter.Watch(ctx)
}

func (e *DeviceIdHeartbeat_Emitter_impl) Emit(message DeviceIdHeartbeat_Emitter_Message) error {
	return e.emitter.Emit(message.Key, message.Value)
}

func (e *DeviceIdHeartbeat_Emitter_impl) EmitBulk(ctx context.Context, messages []DeviceIdHeartbeat_Emitter_Message) error {
	b := []runner.EmitMessage{}
	for _, m := range messages {
		b = append(b, &impl_DeviceIdHeartbeat_Emitter_Message{msg: m})
	}
	return e.emitter.EmitBulk(ctx, b)
}
