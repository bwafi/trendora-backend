package pkg

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(cld *cloudinary.Cloudinary, ctx context.Context, file *multipart.FileHeader) (string, error) {
	if file == nil {
		return "", fmt.Errorf("file is nil")
	}

	uploadFile, _ := file.Open()

	defer uploadFile.Close()

	resp, err := cld.Upload.Upload(ctx, uploadFile, uploader.UploadParams{})

	return resp.SecureURL, err
}
