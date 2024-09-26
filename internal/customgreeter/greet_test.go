package customgreeter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicGreeter_Greet(t *testing.T) {
	cases := []struct {
		name string
		in   struct {
			Name     string
			Greeting string
		}
		expect string
	}{
		{
			name: "should return the default value if given an empty string",
			in: struct {
				Name     string
				Greeting string
			}{Name: "", Greeting: ""},
			expect: "Hello, World!",
		},
		{
			name: "should return the correct value if given a string #1",
			in: struct {
				Name     string
				Greeting string
			}{Name: "Jimmy", Greeting: ""},
			expect: "Hello, Jimmy!",
		},
		{
			name: "should return the correct value if given a string #2",
			in: struct {
				Name     string
				Greeting string
			}{Name: "Jimmy", Greeting: "Welcome"},
			expect: "Welcome, Jimmy!",
		},
		{
			name: "should return the correct value if given a string #3",
			in: struct {
				Name     string
				Greeting string
			}{Name: "Jimmy", Greeting: "Hi"},
			expect: "Hi, Jimmy!",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CustomGreeter{}.Greet(context.Background(), tc.in.Greeting, tc.in.Name)
			assert.Equal(t, tc.expect, got)
		})
	}
}
