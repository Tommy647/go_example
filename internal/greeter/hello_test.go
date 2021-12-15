package greeter

import (
	"context"
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
			got := BasicGreeter{}.Greet(context.Background(), tc.in)
			assert.Equal(t, tc.expect, got)
		})
	}

}

func TestBasicGreeter_CoffeeGreet(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		expect string
	}{
		{
			name:   "should return a free coffee if coffee requested not in db",
			in:     "latte",
			expect: "Free latte served from strings"},
		{
			name:   "should propose help when no drink has been requested",
			in:     "",
			expect: "How can we help?",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := BasicGreeter{}.CoffeeGreet(context.Background(), tc.in)
			assert.Equal(t, tc.expect, got)
		})

	}
}
