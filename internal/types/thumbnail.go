package types

import (
	"thumbnail/internal/storage"
)

type Thumbnail interface {
	Open(storage storage.Storage) (Thumbnail, error)
	Generate(width, height int, storageFile storage.Storage) (string, error)
}

func Open(path string) (Thumbnail, error) {
	storage := storage.StorageFile{Path: path}
	thumbnail := ThumbnailImage{}
	resource, err := thumbnail.Open(&storage)

	if err != nil {
		return nil, err
	}

	return resource, nil
}
