//go:generate mockgen -destination=router_mocks_test.go -package=router github.com/ferrier1/network-device-grpc/internal/router Store

package router

import "context"

// Router - Contains the definition for the router
type Router struct {
	ID       string
	Hostname string
	MgmtIP   string
	Vendor   string
}

// Store - defines the interface we expect
// our database implementation to follow
type Store interface {
	GetRouterByID(id string) (Router, error)
	InsertRouter(router Router) (Router, error)
	DeleteRouter(id string) error
}

// Service - defines the router service
// responsible for updating the router inventory
type Service struct {
	Store Store
}

// New - returns a new instance of the router service
func New(store Store) Service {
	return Service{Store: store}
}

// GetRouterByID - gets the router based on id from the underlying storage
func (s Service) GetRouterByID(ctx context.Context, id string) (Router, error) {
	r, err := s.Store.GetRouterByID(id)
	if err != nil {
		return Router{}, err
	}
	return r, nil
}

// InsertRouter - inserts a router into the underlying storage
func (s Service) InsertRouter(ctx context.Context, router Router) (Router, error) {
	r, err := s.Store.InsertRouter(router)
	if err != nil {
		return Router{}, err
	}
	return r, nil
}

// DeleteRouter - inserts a router into the underlying storage
func (s Service) DeleteRouter(ctx context.Context, id string) error {
	err := s.Store.DeleteRouter(id)
	if err != nil {
		return err
	}
	return nil
}
