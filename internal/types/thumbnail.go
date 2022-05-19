package types

import (
	"thumbnail/internal/storage"
)

type Thumbnail interface {
	Open(storage *storage.StorageFile) (Thumbnail, error)
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
