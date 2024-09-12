package grpcserver

import (
	"context"
	"log"
	"net"

	rtr "github.com/Tommy647/go_example/internal/router"

	"google.golang.org/grpc/reflection"

	"github.com/Tommy647/go_example"
	"google.golang.org/grpc"
)

// ensure our client implements the interface - this breaks compilation if it fails
var _ go_example.HelloServiceServer = &HelloServer{}

// GreetProvider something that greets
type GreetProvider interface {
	Greet(context.Context, string) string
}

// HelloServer provides the implementation of our gRPC service
// has to meet the go_example.HelloServiceServer interface
type HelloServer struct {
	greeter GreetProvider
}

// New instance of our gRPC service
func New(g GreetProvider) *HelloServer {
	return &HelloServer{
		greeter: g,
	}
}

// Hello responds to the Hello gRPC call
func (h HelloServer) Hello(ctx context.Context, request *go_example.HelloRequest) (*go_example.HelloResponse, error) {
	// ensure our context is still valid
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default: // intentionally blank
	}

	return &go_example.HelloResponse{Response: h.greeter.Greet(ctx, request.GetName())}, nil
}

// RouterService - defines the router service interface
// which any service passed in to our Handler will need to
// conform to

type RouterService interface {
	GetRouterByID(ctx context.Context, id string) (rtr.Router, error)
	InsertRouter(ctx context.Context, rtr rtr.Router) (rtr.Router, error)
	DeleteRouter(ctx context.Context, id string) error
}

// Handler -
type Handler struct {
	RouterService RouterService
}

// NewRtrService - returns a new gRPC handler for the router service
func NewRtrService(rtrService RouterService) Handler {
	return Handler{
		RouterService: rtrService,
	}
}

// Serve - starts the gRPC listener
func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	go_example.RegisterRouterServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}

	return nil
}

// GetRouter - retrieves a router by id and returns the response.
func (h Handler) GetRouter(ctx context.Context, r *go_example.GetRouterRequest) (*go_example.GetRouterResponse, error) {
	log.Print("Get Router gRPC Endpoint Hit")
	router, err := h.RouterService.GetRouterByID(ctx, r.Id)
	if err != nil {
		log.Print("Failed to retrieve router by ID")
		return &go_example.GetRouterResponse{}, err
	}

	return &go_example.GetRouterResponse{
		Router: &go_example.Router{
			ID:       router.ID,
			Hostname: router.Hostname,
			Vendor:   router.Vendor,
			MgmtIP:   router.MgmtIP,
		},
	}, nil
}

func (h Handler) AddRouter(ctx context.Context, r *go_example.AddRouterRequest) (*go_example.AddRouterResponse, error) {
	log.Println("Add Router gRPC Endpoint Hit")

	routerToAdd := rtr.Router{
		ID:       r.Router.ID,
		Hostname: r.Router.Hostname,
		Vendor:   r.Router.Vendor,
		MgmtIP:   r.Router.MgmtIP,
	}
	routerAdded, err := h.RouterService.InsertRouter(ctx, routerToAdd)
	if err != nil {
		log.Println("failed to add router")
		return &go_example.AddRouterResponse{}, err
	}

	return &go_example.AddRouterResponse{Router: &go_example.Router{
		ID:       routerAdded.ID,
		Hostname: routerAdded.Hostname,
		Vendor:   routerAdded.Vendor,
		MgmtIP:   routerAdded.MgmtIP,
	}}, nil
}

func (h Handler) DeleteRouter(ctx context.Context, r *go_example.DeleteRouterRequest) (*go_example.DeleteRouterResponse, error) {
	log.Println("Delete Router gRPC Endpoint Hit")

	err := h.RouterService.DeleteRouter(ctx, r.Router.ID)
	if err != nil {
		log.Println("failed to delete router by id")
		return &go_example.DeleteRouterResponse{Status: "Unsuccessful"}, err
	}

	return &go_example.DeleteRouterResponse{Status: "Success"}, nil
}
