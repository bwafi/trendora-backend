package pkg

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(cld *cloudinary.Cloudinary, ctx context.Context, file any) (string, error) {
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})

	return resp.SecureURL, err
}
