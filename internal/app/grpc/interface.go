package grpc

import "net/http"

type handler interface {
	BinaryFileUpload(w http.ResponseWriter, r *http.Request, params map[string]string)
}
