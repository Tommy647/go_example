package grpcclient

import (
	"google.golang.org/grpc"

	_grpc "github.com/Tommy647/go_example"
)

// Client to handle making the gRPC request to the grpcServer
type Client struct {
	client      _grpc.HelloWorldServiceClient // client for gRPC requests
	conn        *grpc.ClientConn
	workers     int    // workers to create for processing requests
	host        string // host we attempt to connect to
	dialOptions []grpc.DialOption
}

// New instance of our Client, accepts variadic options
func New(opts ...Option) (*Client, error) {
	c := &Client{
		workers: 4, // default value
	}
	for _, opt := range opts {
		opt(c)
	}
	if err := c.connect(); err != nil {
		return nil, err
	}
	return c, nil
}

// Close the service and terminate any open connections
func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// connect attempt to open the tcp connection for grpc
func (c *Client) connect() error {
	var err error
	c.conn, err = grpc.Dial(c.host, c.dialOptions...)
	if err != nil {
		return err
	}
	// create a new client on our connection
	c.client = _grpc.NewHelloWorldServiceClient(c.conn)
	return nil
}
