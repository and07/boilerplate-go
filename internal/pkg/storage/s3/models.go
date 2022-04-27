package s3

type UploaderCfg struct {
	Region    string
	Endpoint  string
	AccessKey string
	SecretKey string
}

type FileUploadRequest struct {
	Bucket      string
	Key         string
	Body        []byte
	ContentType string
}

type FileUploadResponse struct {
	URL string
}
