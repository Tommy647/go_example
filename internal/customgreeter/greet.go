// Package customgreeter providers a custom greeting service
package customgreeter

import (
	"context"
	"fmt"

	"github.com/Tommy647/go_example/internal/greeter"
)

// New instance of a custom greeter
func New() CustomGreeter {
	return CustomGreeter{}
}

// CustomGreeter with strings
type CustomGreeter struct{}

// Greet the name in the given string with the given greeting or return a default value if it is empty
func (CustomGreeter) Greet(ctx context.Context, greeting string, name string) string {
	basicGreeter := greeter.New()
	if greeting != "" {
		return fmt.Sprintf("%s, %s!", greeting, name)
	}
	return basicGreeter.Greet(ctx, name)
}
