package boilerplate

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	log "gitlab.com/and07/boilerplate-go/internal/pkg/logger"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
)

// Boilerplate ...
type Boilerplate struct {
}

// Option ...
type Option func(*Boilerplate)

// New ...
func New(ctx context.Context, opts ...Option) *Boilerplate {
	log.Info("Start Boilerplate")

	b := &Boilerplate{}

	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(b)
	}

	return b
}

// HelloWorld ...
func (*Boilerplate) HelloWorld(ctx context.Context, in *empty.Empty) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("Hello World"),
	}, nil
}
