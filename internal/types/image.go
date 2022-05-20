package types

import (
	"errors"
	"fmt"
	"image"
	"os"
	"thumbnail/internal/storage"

	"github.com/disintegration/imaging"
)

var formats = []string{"jpeg", "jpg", "png"}

type ThumbnailImage struct {
	Resource image.Image
	Format   string
}

func (img *ThumbnailImage) Open(storage storage.Storage) (Thumbnail, error) {
	if !storage.Supported(formats) {
		return nil, errors.New("unsupported format")
	}

	_, err := storage.GetFile()

	if err != nil {
		return nil, err
	}

	file, _ := os.Open(storage.Resource().Name())
	src, format, err := image.Decode(file)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	img.Format = format
	img.Resource = src

	return img, nil
}

func (img *ThumbnailImage) Generate(width, height int, storageFile storage.Storage) (string, error) {
	dstImage := imaging.Resize(img.Resource, width, height, imaging.CatmullRom)

	path, extension := storage.Extension(storageFile.Resource().Name())
	pathNewFile := fmt.Sprintf("%s_%d_%d.%s", path, width, height, extension)

	err := imaging.Save(dstImage, pathNewFile)

	return pathNewFile, err
}
