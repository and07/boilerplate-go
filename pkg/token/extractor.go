package token

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	// HeaderName ...
	HeaderName = "Authorization"
)

// Extractor ...
type Extractor interface {
	ExtractGRPC(ctx context.Context) (header string, existStatus bool)
}

type token struct{}

// ExtractGRPC ...
func (t *token) ExtractGRPC(ctx context.Context) (header string, existStatus bool) {

	md, existStatus := metadata.FromIncomingContext(ctx)
	if !existStatus {
		return
	}

	if authHeader, existStatus := md[strings.ToLower(HeaderName)]; existStatus && len(authHeader) > 0 {
		if header = authHeader[0]; header != "" {
			return header, existStatus
		}
	}

	return header, false

}

// ExtractHTTP ...
func (t *token) ExtractHTTP(r *http.Request) (header string, existStatus bool) {

	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		return
	}
	return authHeaderContent[1], true
}

// NewExtractor ...
func NewExtractor() Extractor {
	return &token{}
}
