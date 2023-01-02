package utils

import (
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/notblessy/memoriku/config"
)

func UploadImage(file interface{}, fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.CloudinaryCloudName(), config.CloudinaryAPIKey(), config.CloudinaryAPISecret())
	if err != nil {
		return "", err
	}

	// upload file
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: config.CloudinaryUploadFolder()})
	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}

func RemoveImage(file interface{}, fileName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.CloudinaryCloudName(), config.CloudinaryAPIKey(), config.CloudinaryAPISecret())
	if err != nil {
		return err
	}

	// remove file
	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: fileName,
	})
	if err != nil {
		return err
	}
	return nil
}
