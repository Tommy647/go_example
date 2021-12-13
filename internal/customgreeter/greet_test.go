package greeter

import (
	"context"
	"github.com/Tommy647/go_example/internal/greeter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicGreeter_Greet(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		expect string
	}{
		{
			name:   "should return the default value if given an empty string",
			in:     "",
			expect: "Hello, World!",
		},
		{
			name:   "should return the correct value if given a string #1",
			in:     "Tom",
			expect: "Hello, Tom!",
		},
		{
			name:   "should return the correct value if given a string #2",
			in:     "Orson",
			expect: "Hello, Orson!",
		},
		{
			name:   "should return the correct value if given a string #3",
			in:     "Kurt",
			expect: "Hello, Kurt!",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := greeter.BasicGreeter{}.Greet(context.Background(), tc.in)
			assert.Equal(t, tc.expect, got)
		})
	}
}
