// Package greeter providers a greeting service
package greeter

import "fmt"

const (
	// defaultGreeting if nothing is provided
	defaultGreeting = `World`
	// greetingMessage as a formatting string
	helloGreetingMessage = `Hello, %s!`
)

// HelloGreet the name in the given string or return a default value if it is empty
func HelloGreet(in string) string {
	greeting := defaultGreeting
	if in != "" {
		greeting = in
	}
	return fmt.Sprintf(helloGreetingMessage, greeting)
}
