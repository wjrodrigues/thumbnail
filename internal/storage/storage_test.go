package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsIfFormatIsSupported(t *testing.T) {
	extensions := []string{"png", "jpg", "jpeg", "gif"}
	storage := StorageFile{Path: "image.png"}
	response := storage.Supported(extensions)

	assert.True(t, response)
}

func TestReturnsIfFormatIsNotSupported(t *testing.T) {
	extensions := []string{"png", "jpg", "jpeg", "gif"}
	storage := StorageFile{Path: "image.mp4"}
	response := storage.Supported(extensions)

	assert.False(t, response)
}

func TestReturnsFileInstanceWhenGettingFileFromURL(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
	storage := StorageFile{Path: url}
	file, _ := storage.GetFile()

	assert.IsType(t, file, (*StorageFile)(nil))
}

func TestReturnsErrorWhenURLIsInvalid(t *testing.T) {
	storage := StorageFile{Path: "https://localhost/invalid.png"}
	_, err := storage.GetFile()

	assert.Error(t, err)
}

func TestReturnsFileInstanceWhenGettingFileFromPath(t *testing.T) {
	storage := StorageFile{Path: "../../test/testdata/google_logo.png"}
	file, _ := storage.GetFile()

	assert.IsType(t, file, (*StorageFile)(nil))
}

func TestReturnsErrorWhenURLOrPath(t *testing.T) {
	storage := StorageFile{Path: "../test/testdata/invalid.png"}
	_, err := storage.GetFile()

	assert.Error(t, err)
}

func TestReturnsErrorWhenUnableToCreateTemporaryFileByURL(t *testing.T) {
	bkpTargetPath := TargetPath
	defer func() { TargetPath = bkpTargetPath }()

	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"

	TargetPath = "/nil/"
	storage := StorageFile{Path: url}
	_, err := storage.GetFile()

	assert.Error(t, err)
}

func TestReturnsErrorWhenUnableToCreateTemporaryFileByPath(t *testing.T) {
	bkpTargetPath := TargetPath
	defer func() { TargetPath = bkpTargetPath }()

	url := "../../test/testdata/google_logo.png"

	TargetPath = "/nil/"
	storage := StorageFile{Path: url}
	_, err := storage.GetFile()

	assert.Error(t, err)
}
