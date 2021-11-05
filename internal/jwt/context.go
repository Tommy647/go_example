package jwt

import "context"

// WithUser adds our user to the new context
func WithUser(ctx context.Context, user *CustomClaims) context.Context {
	return context.WithValue(ctx, CTXUser, user)
}

// GetUser from context
func GetUser(ctx context.Context) *CustomClaims {
	if user, ok := ctx.Value(CTXUser).(*CustomClaims); ok {
		return user
	}
	return nil
}
