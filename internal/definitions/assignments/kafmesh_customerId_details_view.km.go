// Code generated by kafmesh-gen. DO NOT EDIT.

package assignments

import (
	"context"
	"os"
	"path/filepath"

	"github.com/burdiyan/kafkautil"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/storage"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/syncromatics/kafmesh/pkg/runner"

	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
)

type KafmeshCustomerIDDetails_View interface {
	Keys() []string
	Get(key string) (*customerId.Details, error)
}

type KafmeshCustomerIDDetails_View_impl struct {
	view *goka.View
}

func New_KafmeshCustomerIDDetails_View(options runner.ServiceOptions) (*KafmeshCustomerIDDetails_View_impl, error) {
	brokers := options.Brokers
	protoWrapper := options.ProtoWrapper

	codec, err := protoWrapper.Codec("kafmesh.customerId.details", &customerId.Details{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create codec")
	}

	opts := &opt.Options{
		BlockCacheCapacity: opt.MiB * 1,
		WriteBuffer:        opt.MiB * 1,
	}

	path := filepath.Join("/tmp/storage", "view", "kafmesh.customerId.details")

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create view db directory")
	}

	builder := storage.BuilderWithOptions(path, opts)

	view, err := goka.NewView(brokers,
		goka.Table("kafmesh.customerId.details"),
		codec,
		goka.WithViewStorageBuilder(builder),
		goka.WithViewHasher(kafkautil.MurmurHasher),
	)

	if err != nil {
		return nil, errors.Wrap(err, "failed creating view")
	}

	return &KafmeshCustomerIDDetails_View_impl{
		view: view,
	}, nil
}

func (v *KafmeshCustomerIDDetails_View_impl) Watch(ctx context.Context) func() error {
	return func() error {
		return v.view.Run(ctx)
	}
}

func (v *KafmeshCustomerIDDetails_View_impl) Keys() []string {
	return v.Keys()
}

func (v *KafmeshCustomerIDDetails_View_impl) Get(key string) (*customerId.Details, error) {
	m, err := v.view.Get(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get value from view")
	}

	if m == nil {
		return nil, nil
	}

	msg, ok := m.(*customerId.Details)
	if !ok {
		return nil, errors.Errorf("expecting message of type '*customerId.Details' got type '%t'", m)
	}

	return msg, nil
}
