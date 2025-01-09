package pkg

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(cld *cloudinary.Cloudinary, ctx context.Context, file *multipart.FileHeader, productId string, variantpPefix string) (string, error) {
	if file == nil {
		return "", fmt.Errorf("file is nil")
	}

	uploadFile, _ := file.Open()
	defer uploadFile.Close()

	var folder, publicID string
	if variantpPefix != "" {

		folder = fmt.Sprintf("%s/%s", productId, variantpPefix)
		publicID = fmt.Sprintf("%s-%d", strings.Split(variantpPefix, "/")[1], time.Now().Unix())
	} else {

		folder = fmt.Sprint(productId)
		publicID = fmt.Sprintf("%s-%d", productId, time.Now().Unix())

	}

	resp, err := cld.Upload.Upload(ctx, uploadFile, uploader.UploadParams{
		UniqueFilename: api.Bool(false),
		PublicID:       publicID,
		Folder:         folder,
	})

	return resp.SecureURL, err
}
