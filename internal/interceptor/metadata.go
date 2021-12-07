package interceptor

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

// errMissingMetaData no metadata available for given key
var errMissingMetaData = errors.New("missing metadata")

// getMetadata value for the given key
func getMetadata(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md) == 0 {
		return "", errMissingMetaData
	}

	value := md[key]
	if len(value) == 0 {
		return "", errMissingMetaData
	}
	return value[0], nil
}
