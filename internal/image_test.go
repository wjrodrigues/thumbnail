package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsErrorWithUnsupportedFormatMessageWhenOpen(t *testing.T) {
	thumbnailImage := ThumbnailImage{}
	storage := StorageFile{path: "../test/testdata/google_logo.mp4"}
	_, err := thumbnailImage.Open(&storage)

	assert.EqualError(t, err, "unsupported format")
}

func TestOpenAndReturnThumbnailImageInstance(t *testing.T) {
	thumbnailImage := ThumbnailImage{}
	storage := StorageFile{path: "../test/testdata/google_logo.png"}
	response, _ := thumbnailImage.Open(&storage)

	assert.Implements(t, (*Thumbnail)(nil), response)
}

func TestOpenImageFromURLAndReturnThumbnail(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
	storage := StorageFile{path: url}
	thumbnailImage := ThumbnailImage{}
	response, _ := thumbnailImage.Open(&storage)

	assert.Implements(t, (*Thumbnail)(nil), response)
}

func TestOpenAndReturnInvalidPathError(t *testing.T) {
	thumbnailImage := ThumbnailImage{}
	storage := StorageFile{path: "../test/testdata/invalid.png"}
	_, err := thumbnailImage.Open(&storage)

	assert.Error(t, err)
}

func TestReturnsErrorWhenDecoderFails(t *testing.T) {
	thumbnailImage := ThumbnailImage{}
	storage := StorageFile{path: "../test/testdata/invalid_image.png"}
	_, err := thumbnailImage.Open(&storage)

	assert.Error(t, err)
}
