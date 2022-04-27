package s3

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type client struct {
	session *session.Session
}

func (c *client) FileUpload(ctx context.Context, request FileUploadRequest) (response FileUploadResponse, err error) {

	uploader := s3manager.NewUploader(c.session)
	input := &s3manager.UploadInput{
		Bucket:      aws.String(request.Bucket),      // bucket's name
		Key:         aws.String(request.Key),         // files destination location
		Body:        bytes.NewReader(request.Body),   // content of the file
		ContentType: aws.String(request.ContentType), // content type
	}
	output, err := uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		return
	}
	response.URL = output.Location
	return

}

func NewFileUploader(cfg UploaderCfg) *client {

	s3Config := &aws.Config{
		Region:      aws.String(cfg.Region),
		Endpoint:    aws.String(cfg.Endpoint),
		Credentials: credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
	}
	s3Session := session.New(s3Config)

	return &client{
		session: s3Session,
	}
}
