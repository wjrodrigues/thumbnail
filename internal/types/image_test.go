package types

import (
	"testing"
	"thumbnail/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestReturnsErrorWithUnsupportedFormatMessageWhenOpen(t *testing.T) {
	thumbnailImage := ThumbnailImage{}
	storage := storage.StorageFile{Path: "../../test/testdata/google_logo.mp4"}
	_, err := thumbnailImage.Open(&storage)

	assert.EqualError(t, err, "unsupported format")
}

func TestOpenAndReturnThumbnailImageInstance(t *testing.T) {
	thumbnailImage := ThumbnailImage{}
	storage := storage.StorageFile{Path: "../../test/testdata/google_logo.png"}
	response, _ := thumbnailImage.Open(&storage)

	assert.Implements(t, (*Thumbnail)(nil), response)
}

func TestOpenImageFromURLAndReturnThumbnail(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
	storage := storage.StorageFile{Path: url}
	thumbnailImage := ThumbnailImage{}
	response, _ := thumbnailImage.Open(&storage)

	assert.Implements(t, (*Thumbnail)(nil), response)
}

func TestOpenAndReturnInvalidPathError(t *testing.T) {
	thumbnailImage := ThumbnailImage{}
	storage := storage.StorageFile{Path: "../../test/testdata/invalid.png"}
	_, err := thumbnailImage.Open(&storage)

	assert.Error(t, err)
}

func TestReturnsErrorWhenDecoderFails(t *testing.T) {
	thumbnailImage := ThumbnailImage{}
	storage := storage.StorageFile{Path: "../../test/testdata/invalid_image.png"}
	_, err := thumbnailImage.Open(&storage)

	assert.Error(t, err)
}
