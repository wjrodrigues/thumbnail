package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsIfFormatIsSupported(t *testing.T) {
	extensions := []string{"png", "jpg", "jpeg", "gif"}
	storage := StorageFile{path: "image.png"}
	response := storage.Supported(extensions)

	assert.True(t, response)
}

func TestReturnsIfFormatIsNotSupported(t *testing.T) {
	extensions := []string{"png", "jpg", "jpeg", "gif"}
	storage := StorageFile{path: "image.mp4"}
	response := storage.Supported(extensions)

	assert.False(t, response)
}

func TestReturnsFileInstanceWhenGettingFileFromURL(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
	storage := StorageFile{path: url}
	file, _ := storage.GetFile()

	assert.IsType(t, file, (*StorageFile)(nil))
}

func TestReturnsErrorWhenURLIsInvalid(t *testing.T) {
	storage := StorageFile{path: "https://localhost/invalid.png"}
	_, err := storage.GetFile()

	assert.Error(t, err)
}

func TestReturnsFileInstanceWhenGettingFileFromPath(t *testing.T) {
	storage := StorageFile{path: "../test/testdata/google_logo.png"}
	file, _ := storage.GetFile()

	assert.IsType(t, file, (*StorageFile)(nil))
}

func TestReturnsErrorWhenURLOrPath(t *testing.T) {
	storage := StorageFile{path: "../test/testdata/invalid.png"}
	_, err := storage.GetFile()

	assert.Error(t, err)
}

func TestOpenURLAndReturnsInstanceOfThumbnailImage(t *testing.T) {
	storage := StorageFile{path: "../test/testdata/google_logo.png"}
	response, _ := storage.Open()

	assert.IsType(t, ThumbnailImage{}, response)
}

func TestReturnsErrorWhenOpeningUnsupportedFormat(t *testing.T) {
	storage := StorageFile{path: "../test/testdata/google_logo.mp4"}
	_, err := storage.Open()

	assert.Error(t, err)
}

func TestReturnsErrorWhenUnableToCreateTemporaryFile(t *testing.T) {
	bkpTargetPath := TargetPath
	defer func() { TargetPath = bkpTargetPath }()

	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"

	TargetPath = "/nil/"
	storage := StorageFile{path: url}
	_, err := storage.Open()

	assert.Error(t, err)
}
