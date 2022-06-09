package types

import (
	"errors"
	"fmt"
	"image"
	"os"
	"thumbnail/internal/storage"

	"github.com/disintegration/imaging"
)

var ImageFormats = []string{"jpeg", "jpg", "png"}

type ThumbnailImage struct {
	Storage  storage.Storage
	Resource image.Image
	Format   string
}

func (img *ThumbnailImage) Open() (Thumbnail, error) {
	if !img.Storage.Supported(ImageFormats) {
		return nil, errors.New("unsupported format")
	}

	_, err := img.Storage.GetFile()

	if err != nil {
		return nil, err
	}

	file, _ := os.Open(img.Storage.Resource().Name())
	src, format, err := image.Decode(file)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	img.Format = format
	img.Resource = src

	return img, nil
}

func (img *ThumbnailImage) Generate(width, height, time int) (string, error) {
	dstImage := imaging.Resize(img.Resource, width, height, imaging.CatmullRom)

	path, extension := storage.Extension(img.Storage.Resource().Name())
	pathNewFile := fmt.Sprintf("%s_%d_%d.%s", path, width, height, extension)

	err := imaging.Save(dstImage, pathNewFile)

	return pathNewFile, err
}
