// Package greeter providers a greeting service
package greeter

import (
	"context"
	"fmt"
)

const (
	// defaultGreeting if nothing is provided
	defaultGreeting = `World`
	// helloGreetingMessage as a formatting string
	helloGreetingMessage = `Hello, %s!`
)

// New instance of a string greeter
func New() BasicGreeter {
	return BasicGreeter{}
}

// BasicGreeter with strings
type BasicGreeter struct{}

// Greet the name in the given string or return a default value if it is empty
func (BasicGreeter) Greet(_ context.Context, in string) string {
	greeting := defaultGreeting
	if in != "" {
		greeting = in
	}
	return fmt.Sprintf(helloGreetingMessage, greeting)
}
