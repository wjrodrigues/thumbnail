package types

import (
	"errors"
	"thumbnail/internal/storage"
)

type Thumbnail interface {
	Open() (Thumbnail, error)
	Generate(width, height, time int) (string, error)
}

func Open(path string) (Thumbnail, error) {
	var resource Thumbnail
	var err error

	storageFile := storage.StorageFile{Path: path}

	if storageFile.Supported(ImageFormats) {
		thumbnail := ThumbnailImage{Storage: &storageFile}
		resource, err = thumbnail.Open()
	}

	if storageFile.Supported(VideoFormats) {
		thumbnail := ThumbnailVideo{Storage: &storageFile}
		resource, err = thumbnail.Open()
	}

	if resource == nil {
		return nil, errors.New("unsupported format")
	}

	return resource, err
}
