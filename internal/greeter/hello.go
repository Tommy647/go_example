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
	// freeCoffee message as a formatting string
	defaultCoffee = `How can we help?`
	// freeCoffee message as a formatting string
	freeCoffee = `Free %s served from strings`
)

// Greet with strings
type Greet struct{}

// New instance of a string greeter
func New() Greet {
	return Greet{}
}

// Greet the name in the given string or return a default value if it is empty
func (Greet) Greet(_ context.Context, in string) string {
	greeting := defaultGreeting
	if in != "" {
		greeting = in
	}
	return fmt.Sprintf(helloGreetingMessage, greeting)
}

// CoffeeGreet implements the string response when no DB is used
func (Greet) CoffeeGreet(_ context.Context, in string) string {
	drink := defaultCoffee
	if in != "" {
		return fmt.Sprintf(freeCoffee, in)
	}
	return fmt.Sprint(drink)
}
