package client

import (
	"bytes"
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	_grpc "github.com/Tommy647/grpc"
)

func TestClient_Run(t *testing.T) {

	tests := []struct {
		name   string
		expect string
	}{
		{
			name:   "should...",
			expect: "Message:  Hello, World!\n",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			// write to a buffer instead of stdout
			log.SetOutput(&buf)
			// no date stamp - because those are a pain to test
			log.SetFlags(0)
			defer func() {
				log.SetOutput(os.Stderr)
			}()

			server := &mockServer{}

			server.On("HelloWorld", mock.Anything, &_grpc.HelloRequest{}).Return(&_grpc.HelloResponse{Response: "Hello, World!"}, nil)

			c := Client{
				client: server,
				// names:  tc.names,
			}

			c.Run(context.Background())

			assert.Equal(t, tc.expect, buf.String())
		})
	}
}

type mockServer struct {
	mock.Mock
}

func (m *mockServer) HelloWorld(ctx context.Context, request *_grpc.HelloRequest, _ ...grpc.CallOption) (*_grpc.HelloResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*_grpc.HelloResponse), args.Error(1)
}
