package s3

import (
	"context"
	"os"
	"testing"
)

func TestUploader(t *testing.T) {

	dat, err := os.ReadFile("/Users/myakotkin.andrey/1.png")

	uploader := NewFileUploader(UploaderCfg{
		Region:    "ru-1",
		Endpoint:  "https://s3.storage.selcloud.ru",
		AccessKey: "211825",
		SecretKey: "m+8M2RZdLo",
	})

	res, err := uploader.FileUpload(context.Background(), FileUploadRequest{
		Bucket:      "gogreat1",
		Key:         "images/profile/ddd2.jpg",
		Body:        dat,
		ContentType: "image/png",
	})
	if err != nil {
		t.Error(err)
	}

	t.Errorf("%#v", res)

}
