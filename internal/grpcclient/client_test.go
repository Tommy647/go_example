package grpcclient

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	_grpc "github.com/Tommy647/go_example"
)

func TestClient_Run(t *testing.T) {
	tests := []struct {
		name   string
		names  []string
		expect string
	}{
		{
			name:   "should correctly handle an empty list of names",
			names:  nil,
			expect: "Message: Hello, World!\n",
		},
		{
			name:  "should correctly handle a list of names",
			names: []string{"Tom", "Orson", "Kurt"},
			expect: `Message: Hello, Tom!
Message: Hello, Orson!
Message: Hello, Kurt!
`,
		},
		{
			name:   "should correctly handle errors when using a list of names",
			names:  []string{"error"}, // we are going to watch for this particular string when creating mocks
			expect: "error messaging grpcServer oops! something went wrong\n",
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

			// use a mock grpc service - we are testing we send the correct requests and handle the responses correctly
			mServer := &mockServer{}

			// mock each of the expected HelloWorld request
			if len(tc.names) == 0 {
				mServer.On("Hello", mock.Anything, &_grpc.HelloRequest{}).
					Return(&_grpc.HelloResponse{Response: "Hello, World!"}, nil)
			}

			for _, name := range tc.names {
				if name == "error" {
					mServer.On(
						"Hello", mock.Anything, &_grpc.HelloRequest{Name: name}).
						Return((*_grpc.HelloResponse)(nil), errors.New("oops! something went wrong"))
				}

				mServer.On("Hello",
					mock.Anything, &_grpc.HelloRequest{Name: name}).
					Return(&_grpc.HelloResponse{Response: fmt.Sprintf("Hello, %s!", name)}, nil)
			}

			c := Client{
				helloClient: mServer,
				workers:     1,
			}

			ops := RequestOpts{
				Context: context.Background(),
				Names:   tc.names,
			}

			c.Run("BasicGreeter", ops)

			assert.Equal(t, tc.expect, buf.String())
			// assert all expected calls to the mServer were made
			mServer.AssertExpectations(t)
		})
	}
}

type mockServer struct {
	mock.Mock
}

func (m *mockServer) Hello(ctx context.Context, request *_grpc.HelloRequest, _ ...grpc.CallOption) (*_grpc.HelloResponse, error) {
	args := m.Called(ctx, request)
	return args.Get(0).(*_grpc.HelloResponse), args.Error(1)
}
