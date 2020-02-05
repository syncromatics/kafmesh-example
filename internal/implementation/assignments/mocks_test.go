package assignments_test

import (
	"context"

	"kafmesh-example/internal/definitions/assignments"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/warehouse"

	"github.com/pkg/errors"
)

type contextMock struct {
	keys    func() ([]string, error)
	gets    map[string]*customerId.Details
	emits   []*assignments.CustomerIdDetails_Synchronizer_Message
	deletes []string

	keysShouldFail   bool
	getShouldFail    bool
	emitShouldFail   bool
	deleteShouldFail bool
}

func (c *contextMock) Keys() ([]string, error) {
	if c.keysShouldFail {
		return nil, errors.Errorf("boom")
	}

	return c.keys()
}

func (c *contextMock) Get(key string) (*customerId.Details, error) {
	if c.getShouldFail {
		return nil, errors.Errorf("boom")
	}

	return c.gets[key], nil
}

func (c *contextMock) Emit(msg *assignments.CustomerIdDetails_Synchronizer_Message) error {
	if c.emitShouldFail {
		return errors.Errorf("boom")
	}

	c.emits = append(c.emits, msg)
	return nil
}

func (c *contextMock) EmitBulk(context.Context, []*assignments.CustomerIdDetails_Synchronizer_Message) error {
	return nil
}

func (c *contextMock) Delete(key string) error {
	if c.deleteShouldFail {
		return errors.Errorf("boom")
	}

	c.deletes = append(c.deletes, key)
	return nil
}

type repository struct {
	getAll func(ctx context.Context) ([]warehouse.CustomerDetails, error)
}

func (r *repository) GetAll(ctx context.Context) ([]warehouse.CustomerDetails, error) {
	return r.getAll(ctx)
}
