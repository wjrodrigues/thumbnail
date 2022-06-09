package types

import (
	"fmt"
	"os"
	"testing"
	"thumbnail/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestReturnsErrorWithOpenVideo(t *testing.T) {
	storage := storage.StorageFile{Path: "../../test/testdata/go_land.mp4"}
	thumbnailImage := ThumbnailImage{Storage: &storage}
	_, err := thumbnailImage.Open()

	assert.EqualError(t, err, "unsupported format")
}

func TestOpenAndReturnThumbnailImageInstance(t *testing.T) {
	storage := storage.StorageFile{Path: "../../test/testdata/google_logo.png"}
	thumbnailImage := ThumbnailImage{Storage: &storage}
	response, _ := thumbnailImage.Open()

	assert.Implements(t, (*Thumbnail)(nil), response)
}

func TestOpenImageFromURLAndReturnThumbnail(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
	storage := storage.StorageFile{Path: url}
	thumbnailImage := ThumbnailImage{Storage: &storage}
	response, _ := thumbnailImage.Open()

	assert.Implements(t, (*Thumbnail)(nil), response)
}

func TestOpenAndReturnInvalidPathError(t *testing.T) {
	storage := storage.StorageFile{Path: "../../test/testdata/invalid.png"}
	thumbnailImage := ThumbnailImage{Storage: &storage}
	_, err := thumbnailImage.Open()

	assert.Error(t, err)
}

func TestReturnsErrorWhenDecoderFails(t *testing.T) {
	storage := storage.StorageFile{Path: "../../test/testdata/invalid_image.png"}
	thumbnailImage := ThumbnailImage{Storage: &storage}
	_, err := thumbnailImage.Open()

	assert.Error(t, err)
}

func TestReturnsThumbnailNameWithNewDimensions(t *testing.T) {
	storageFile := storage.StorageFile{Path: "../../test/testdata/google_logo.png"}
	thumbnailImage := ThumbnailImage{Storage: &storageFile}
	response, _ := thumbnailImage.Open()

	thumbnailName, _ := response.Generate(1024, 768, 0)

	path, extension := storage.Extension(storageFile.Resource().Name())
	expectedThumbnailName := fmt.Sprintf("%s_%d_%d.%s", path, 1024, 768, extension)

	assert.Equal(t, expectedThumbnailName, thumbnailName)
}

func TestValidateThumbnailCreation(t *testing.T) {
	storageFile := storage.StorageFile{Path: "../../test/testdata/google_logo.png"}
	thumbnailImage := ThumbnailImage{Storage: &storageFile}

	response, _ := thumbnailImage.Open()

	thumbnailName, _ := response.Generate(250, 100, 0)

	_, err := os.Stat(thumbnailName)

	assert.Equal(t, nil, err)
}
