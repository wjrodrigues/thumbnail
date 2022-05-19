package types

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"thumbnail/internal/storage"
)

var formats = []string{"jpeg", "jpg", "gif", "png"}

type ThumbnailImage struct {
	Resource image.Image
	Format   string
}

func (img *ThumbnailImage) Open(storage *storage.StorageFile) (Thumbnail, error) {
	if !storage.Supported(formats) {
		return nil, errors.New("unsupported format")
	}

	_, err := storage.GetFile()

	if err != nil {
		return nil, err
	}

	src, format, err := image.Decode(storage.File)

	if err != nil {
		return nil, err
	}

	img.Format = format
	img.Resource = src

	return img, nil
}
