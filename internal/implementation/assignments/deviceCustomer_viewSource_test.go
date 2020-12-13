package assignments_test

import (
	"kafmesh-example/internal/implementation/assignments"
	"kafmesh-example/internal/implementation/assignments/mocks"
	"kafmesh-example/internal/warehouse"
	"testing"

	kmMocks "kafmesh-example/internal/definitions/assignments"
	"kafmesh-example/internal/definitions/models/kafmesh/deviceId"

	"github.com/golang/mock/gomock"
)

func Test_DeviceCustomerViewSource_Sync(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockassignmentRepo(ctrl)
	repo.EXPECT().AllDetails(gomock.Any()).
		Return([]warehouse.Details{
			{
				DeviceID:   1,
				CustomerID: 2,
			},
			{
				DeviceID:   3,
				CustomerID: 4,
			},
		}, nil)

	ctx := kmMocks.NewMockDeviceCustomer_ViewSource_Context(ctrl)
	ctx.EXPECT().Update("1", &deviceId.Customer{
		Id: 2,
	})
	ctx.EXPECT().Update("3", &deviceId.Customer{
		Id: 4,
	})

	subject := assignments.NewDeviceCustomerViewSource(repo)
	subject.Sync(ctx)
}
