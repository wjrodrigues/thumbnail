package internal

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
)

type ThumbnailImage struct {
	resource image.Image
	format   string
}

func (img ThumbnailImage) Open(storage *StorageFile) (Thumbnail, error) {
	formats := Formats["image"]
	if !storage.Supported(formats) {
		return nil, errors.New("unsupported format")
	}

	_, err := storage.GetFile()

	if err != nil {
		return nil, err
	}

	src, format, err := image.Decode(storage.file)

	if err != nil {
		return nil, err
	}

	img.format = format
	img.resource = src

	return img, nil
}
