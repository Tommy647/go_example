package router

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRouterService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("tests get router by id", func(t *testing.T) {
		routerStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		routerStoreMock.
			EXPECT().
			GetRouterByID(id).
			Return(Router{
				ID: id,
			}, nil)

		routerService := New(routerStoreMock)
		rtr, err := routerService.
			GetRouterByID(
				context.Background(),
				id,
			)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rtr.ID)
	})

	t.Run("tests insert router", func(t *testing.T) {
		routerStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		routerStoreMock.
			EXPECT().
			InsertRouter(Router{
				ID: id,
			}).
			Return(Router{
				ID: id,
			}, nil)

		routerService := New(routerStoreMock)
		rtr, err := routerService.
			InsertRouter(
				context.Background(),
				Router{
					ID: id,
				},
			)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rtr.ID)
	})

	t.Run("tests delete router", func(t *testing.T) {
		routerStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		routerStoreMock.
			EXPECT().
			DeleteRouter(id).
			Return(nil)

		routerService := New(routerStoreMock)
		err := routerService.
			DeleteRouter(
				context.Background(),
				id,
			)

		assert.NoError(t, err)
	})
}
