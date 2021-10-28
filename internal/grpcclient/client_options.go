package grpcclient

import "google.golang.org/grpc"

// Option for variadic configuration of our client
type Option func(*Client)

// WithHost to use for connections
func WithHost(host string) Option {
	return func(c *Client) {
		c.host = host
	}
}

// WithDialOptions for the grpc connection
func WithDialOptions(opts ...grpc.DialOption) Option {
	return func(c *Client) {
		if c.dialOptions == nil {
			c.dialOptions = opts
			return
		}

		c.dialOptions = append(c.dialOptions, opts...)
	}
}
