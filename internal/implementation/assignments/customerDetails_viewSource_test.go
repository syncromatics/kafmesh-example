package assignments_test

import (
	"kafmesh-example/internal/data"
	kmMocks "kafmesh-example/internal/definitions/assignments"
	"kafmesh-example/internal/definitions/models/kafmesh/customerId"
	"kafmesh-example/internal/implementation/assignments"
	"kafmesh-example/internal/implementation/assignments/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_CustomerDetailsViewSource_Sync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockcustomerRepo(ctrl)
	repo.EXPECT().GetCustomerDetails(gomock.Any()).
		Return([]data.CustomerDetail{
			{
				ID:   1,
				Name: "some customer",
			},
			{
				ID:   2,
				Name: "other customer",
			},
		}, nil)

	ctx := kmMocks.NewMockCustomerDetails_ViewSource_Context(ctrl)
	ctx.EXPECT().Update("1", &customerId.Details{
		Name: "some customer",
	})
	ctx.EXPECT().Update("2", &customerId.Details{
		Name: "other customer",
	})

	subject := assignments.NewCustomerDetailsViewSource(repo)
	subject.Sync(ctx)
}
