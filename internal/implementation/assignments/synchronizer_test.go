package assignments_test

import (
	"context"
	"testing"

	assign "kafmesh-example/internal/definitions/assignments"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/implementation/assignments"
	"kafmesh-example/internal/warehouse"

	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func Test_Synchronzier(t *testing.T) {
	cm := &contextMock{
		keys: func() ([]string, error) {
			return []string{
				"1",
				"3",
				"5",
			}, nil
		},
		gets: map[string]*customerId.Details{
			"1": &customerId.Details{
				Name: "1 old name",
			},
			"5": &customerId.Details{
				Name: "5 no change",
			},
		},
	}
	repository := &repository{
		getAll: func(ctx context.Context) ([]warehouse.CustomerDetails, error) {
			return []warehouse.CustomerDetails{
				warehouse.CustomerDetails{
					ID:   1,
					Name: "1 new name",
				},
				warehouse.CustomerDetails{
					ID:   2,
					Name: "new customer",
				},
				warehouse.CustomerDetails{
					ID:   5,
					Name: "5 no change",
				},
			}, nil
		},
	}

	synchronizer := assignments.NewCustomerDetailsSynchronizer(repository)

	err := synchronizer.Sync(cm)
	assert.NilError(t, err)

	assert.DeepEqual(t, cm.emits, []*assign.CustomerIdDetails_Synchronizer_Message{
		&assign.CustomerIdDetails_Synchronizer_Message{
			Key: "1",
			Value: &customerId.Details{
				Name: "1 new name",
			},
		},
		&assign.CustomerIdDetails_Synchronizer_Message{
			Key: "2",
			Value: &customerId.Details{
				Name: "new customer",
			},
		},
	})

	assert.DeepEqual(t, cm.deletes, []string{"3"})
}

func Test_Synchronzier_ShouldReturnErrorIfRepositoryFails(t *testing.T) {
	cm := &contextMock{
		keys: func() ([]string, error) {
			return []string{
				"1",
				"3",
				"5",
			}, nil
		},
		gets: map[string]*customerId.Details{
			"1": &customerId.Details{
				Name: "1 old name",
			},
			"5": &customerId.Details{
				Name: "5 no change",
			},
		},
	}
	repository := &repository{
		getAll: func(ctx context.Context) ([]warehouse.CustomerDetails, error) {
			return nil, errors.Errorf("boom")
		},
	}

	synchronizer := assignments.NewCustomerDetailsSynchronizer(repository)

	err := synchronizer.Sync(cm)
	assert.ErrorContains(t, err, "failed to get details from repository: boom")
}

func Test_Synchronzier_ShouldReturnErrorIfKeysFails(t *testing.T) {
	cm := &contextMock{
		keysShouldFail: true,
	}
	repository := &repository{
		getAll: func(ctx context.Context) ([]warehouse.CustomerDetails, error) {
			return []warehouse.CustomerDetails{
				warehouse.CustomerDetails{
					ID:   1,
					Name: "1 new name",
				},
				warehouse.CustomerDetails{
					ID:   2,
					Name: "new customer",
				},
				warehouse.CustomerDetails{
					ID:   5,
					Name: "5 no change",
				},
			}, nil
		},
	}

	synchronizer := assignments.NewCustomerDetailsSynchronizer(repository)

	err := synchronizer.Sync(cm)
	assert.ErrorContains(t, err, "failed to get keys: boom")
}

func Test_Synchronzier_ShouldReturnErrorIfGetFails(t *testing.T) {
	cm := &contextMock{
		keys: func() ([]string, error) {
			return []string{
				"1",
				"3",
				"5",
			}, nil
		},
		getShouldFail: true,
	}
	repository := &repository{
		getAll: func(ctx context.Context) ([]warehouse.CustomerDetails, error) {
			return []warehouse.CustomerDetails{
				warehouse.CustomerDetails{
					ID:   1,
					Name: "1 new name",
				},
				warehouse.CustomerDetails{
					ID:   2,
					Name: "new customer",
				},
				warehouse.CustomerDetails{
					ID:   5,
					Name: "5 no change",
				},
			}, nil
		},
	}

	synchronizer := assignments.NewCustomerDetailsSynchronizer(repository)

	err := synchronizer.Sync(cm)
	assert.ErrorContains(t, err, "failed to get: boom")
}

func Test_Synchronzier_ShouldReturnErrorIfEmitFails(t *testing.T) {
	cm := &contextMock{
		keys: func() ([]string, error) {
			return []string{
				"1",
				"3",
				"5",
			}, nil
		},
		gets: map[string]*customerId.Details{
			"1": &customerId.Details{
				Name: "1 old name",
			},
			"5": &customerId.Details{
				Name: "5 no change",
			},
		},
		emitShouldFail: true,
	}
	repository := &repository{
		getAll: func(ctx context.Context) ([]warehouse.CustomerDetails, error) {
			return []warehouse.CustomerDetails{
				warehouse.CustomerDetails{
					ID:   1,
					Name: "1 new name",
				},
				warehouse.CustomerDetails{
					ID:   2,
					Name: "new customer",
				},
				warehouse.CustomerDetails{
					ID:   5,
					Name: "5 no change",
				},
			}, nil
		},
	}

	synchronizer := assignments.NewCustomerDetailsSynchronizer(repository)

	err := synchronizer.Sync(cm)
	assert.ErrorContains(t, err, "failed to emit customer details: boom")
}

func Test_Synchronzier_ShouldReturnErrorIfDeleteFails(t *testing.T) {
	cm := &contextMock{
		keys: func() ([]string, error) {
			return []string{
				"1",
				"3",
				"5",
			}, nil
		},
		gets: map[string]*customerId.Details{
			"1": &customerId.Details{
				Name: "1 old name",
			},
			"5": &customerId.Details{
				Name: "5 no change",
			},
		},
		deleteShouldFail: true,
	}
	repository := &repository{
		getAll: func(ctx context.Context) ([]warehouse.CustomerDetails, error) {
			return []warehouse.CustomerDetails{
				warehouse.CustomerDetails{
					ID:   1,
					Name: "1 new name",
				},
				warehouse.CustomerDetails{
					ID:   2,
					Name: "new customer",
				},
				warehouse.CustomerDetails{
					ID:   5,
					Name: "5 no change",
				},
			}, nil
		},
	}

	synchronizer := assignments.NewCustomerDetailsSynchronizer(repository)

	err := synchronizer.Sync(cm)
	assert.ErrorContains(t, err, "failed to delete: boom")
}
