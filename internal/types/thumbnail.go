package types

import (
	"errors"
	"thumbnail/internal/storage"
)

type Thumbnail interface {
	Open(storage storage.Storage) (Thumbnail, error)
	Generate(width, height, time int, storageFile storage.Storage) (string, error)
}

func Open(path string) (Thumbnail, error) {
	var resource Thumbnail
	var err error

	storageFile := storage.StorageFile{Path: path}

	if storageFile.Supported(ImageFormats) {
		thumbnail := ThumbnailImage{}
		resource, err = thumbnail.Open(&storageFile)
	}

	if storageFile.Supported(VideoFormats) {
		thumbnail := ThumbnailVideo{}
		resource, err = thumbnail.Open(&storageFile)
	}

	if resource == nil {
		return nil, errors.New("unsupported format")
	}

	return resource, err
}
