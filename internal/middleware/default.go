package middleware

import "net/http"

// Default middleware to apply to every call
var Default = []func(handler http.Handler) http.Handler{
	WithBasicTelemetry,
}

// Secure middle ware for authorised endpoints
var Secure = []func(handler http.Handler) http.Handler{
	WithAuth,
	WithRole,
}

// WithDefault wrap requests in middle ware, secure first, then default
func WithDefault(next http.Handler, secure ...bool) http.Handler {
	if len(secure) != 0 && secure[0] {
		for i := len(Secure) - 1; i >= 0; i-- {
			next = Secure[i](next)
		}
	}

	for i := len(Default) - 1; i >= 0; i-- {
		next = Default[i](next)
	}
	return next
}
