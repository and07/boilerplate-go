package token

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	HeaderName = "Authorization"
)

// Token ...
type Extractor interface {
	ExtractGRPC(ctx context.Context) (header string, existStatus bool)
}

type token struct{}

func (t *token) ExtractGRPC(ctx context.Context) (header string, existStatus bool) {

	md, existStatus := metadata.FromIncomingContext(ctx)
	if !existStatus {
		return
	}

	if authHeaderContent, existStatus := md[strings.ToLower(HeaderName)]; existStatus && len(authHeaderContent) > 0 {
		if header = authHeaderContent[0]; header != "" {
			authHeaderContent := strings.Split(header, " ")
			return authHeaderContent[1], existStatus
		}
	}

	return header, false

}

func (t *token) ExtractHTTP(r *http.Request) (header string, existStatus bool) {

	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		return
	}
	return authHeaderContent[1], true
}

// NewToken ...
func NewExtractor() Extractor {
	return &token{}
}