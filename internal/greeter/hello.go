// Package greeter providers a greeting service
package greeter

import (
	"context"
	"fmt"
)

const (
	// defaultGreeting if nothing is provided
	defaultGreeting = `World`
	// greetingMessage as a formatting string
	helloGreetingMessage = `Hello, %s!`
)

// New instance of a string greeter
func New() Greeter {
	return Greeter{}
}

// Greeter with strings
type Greeter struct{}

// HelloGreet the name in the given string or return a default value if it is empty
func (Greeter) HelloGreet(_ context.Context, in string) string {
	greeting := defaultGreeting
	if in != "" {
		greeting = in
	}
	return fmt.Sprintf(helloGreetingMessage, greeting)
}
