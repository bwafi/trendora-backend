package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/spf13/viper"
)

func NewCloudinary(config *viper.Viper) *cloudinary.Cloudinary {
	url := config.GetString("cloudinary.url")

	cld, _ := cloudinary.NewFromURL(url)

	cld.Config.URL.Secure = true

	return cld
}
