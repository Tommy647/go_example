package grpcserver

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	grpc "github.com/Tommy647/go_example"
	"github.com/Tommy647/go_example/internal/greeter"
)

func TestHelloWorldServer_HelloWorld(t *testing.T) {
	var expectDefault = &grpc.HelloResponse{Response: "Hello, World!"}

	tests := []struct {
		name       string
		ctxTimeout time.Duration
		in         *grpc.HelloRequest
		want       *grpc.HelloResponse
		wantErr    error
	}{
		{
			name:       "should return default if no request is provided",
			ctxTimeout: 10 * time.Second,
			in:         nil,
			want:       expectDefault,
			wantErr:    nil,
		},
		{
			name:       "should return default if request name is blank",
			ctxTimeout: 10 * time.Second,
			in:         &grpc.HelloRequest{Name: ""},
			want:       expectDefault,
			wantErr:    nil,
		},
		{
			name:       "should return the correct name",
			ctxTimeout: 10 * time.Second,
			in:         &grpc.HelloRequest{Name: "Tom"},
			want:       &grpc.HelloResponse{Response: "Hello, Tom!"},
			wantErr:    nil,
		},
		{
			name:       "should error if context has expired",
			ctxTimeout: 0,
			in:         &grpc.HelloRequest{Name: "Tom"},
			want:       nil,
			wantErr:    errors.New("context deadline exceeded"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), tc.ctxTimeout)
			defer cancel()

			got, err := (HelloServer{greeter: greeter.Greet{}}).Hello(ctx, tc.in)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.wantErr.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
