package handlers

import (
	"context"

	"github.com/and07/boilerplate-go/internal/pkg/storage/s3"
)

type logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type uploader interface {
	FileUpload(ctx context.Context, request s3.FileUploadRequest) (response s3.FileUploadResponse, err error)
}
