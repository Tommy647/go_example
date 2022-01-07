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
	// defaultCoffee is used when no kind of coffee is specified in the request
	defaultCoffee = `How can we help?`
	// freeCoffee is used when the specified coffee has no entry in the DB
	freeCoffee = `Free %s served from strings`
	// defaultFruit is used when no fruit was specified in the request
	defaultFruit = `How can we help?`
	// freeFruit is used when the specified fruit has no entry in the DB
	freeFruit = `Free %s served from strings`
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

// CoffeeGreet implements the basic response from strings when no DB is requested
func (BasicGreeter) CoffeeGreet(_ context.Context, in string) string {
	drink := defaultCoffee
	if in != "" {
		return fmt.Sprintf(freeCoffee, in)
	}
	return fmt.Sprintf(drink)
}

func (BasicGreeter) FruitGreet(_ context.Context, in string) string {
	item := defaultFruit
	if in != "" {
		return fmt.Sprintf(freeFruit, in)
	}
	return fmt.Sprintf(item)
}
