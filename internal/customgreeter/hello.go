// Package greeter providers a greeting service
package customgreeter

import (
	"context"
	"fmt"
)

const (
	// helloGreetingMessage as a formatting string
	helloGreetingMessage = `Hello, %s!`
)

// New instance of a string greeter
func New() CustomGreeter {
	return CustomGreeter{}
}

// CustomGreeter with strings
type CustomGreeter struct{}

// Greet the name in the given string or return a default value if it is empty
func (CustomGreeter) Greet(_ context.Context, greeting string, name string) string {

	return fmt.Sprintf(greeting, name)
}
