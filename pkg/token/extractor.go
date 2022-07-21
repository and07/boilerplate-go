package token

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	// HeaderName ...
	HeaderName = "Authorization"
)

// Extractor ...
type Extractor interface {
	ExtractGRPC(ctx context.Context) (header string, err error)
	ExtractHTTP(r *http.Request) (header string, err error)
}

type token struct{}

// ExtractGRPC ...
func (t *token) ExtractGRPC(ctx context.Context) (header string, err error) {
	md, existStatus := metadata.FromIncomingContext(ctx)
	if !existStatus {
		err = status.Errorf(codes.Unauthenticated, "metadata is not provided")
		return
	}

	if authHeaderContent, existStatus := md[strings.ToLower(HeaderName)]; existStatus && len(authHeaderContent) > 0 {
		if header = authHeaderContent[0]; header != "" {
			authHeaderContent := strings.Split(header, " ")
			return authHeaderContent[1], nil
		}
	}

	return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")

}

// ExtractHTTP ...
func (t *token) ExtractHTTP(r *http.Request) (header string, err error) {

	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		err = status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		return
	}
	return authHeaderContent[1], nil
}

// NewToken ...
func NewExtractor() Extractor {
	return &token{}
}
