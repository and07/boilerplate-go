package grpcserver

import "context"

type extractor interface {
	ExtractGRPC(ctx context.Context) (header string, existStatus bool)
}
