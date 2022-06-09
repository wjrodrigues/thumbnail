package types

import (
	"testing"
	"thumbnail/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestReturnsErrorWhenOpenImage(t *testing.T) {
	storage := storage.StorageFile{Path: "../../test/testdata/google_logo.png"}
	thumbnailImage := ThumbnailVideo{Storage: &storage}

	_, err := thumbnailImage.Open()

	assert.EqualError(t, err, "unsupported format")
}

func TestReturnsThumbnailInstanceWhenOpeningVideo(t *testing.T) {
	storage := storage.StorageFile{Path: "../../test/testdata/go_land.mp4"}
	thumbnailVideo := ThumbnailVideo{Storage: &storage}

	response, _ := thumbnailVideo.Open()

	assert.Implements(t, (*Thumbnail)(nil), response)
}

func TestReturnsErrorOpeningInvalidVideoPath(t *testing.T) {
	storage := storage.StorageFile{Path: "../../test/testdata/invalid.mp4"}
	thumbnailVideo := ThumbnailVideo{Storage: &storage}

	_, err := thumbnailVideo.Open()

	assert.Error(t, err)
}

func TestReturnsGIFThumbnailOfVideo(t *testing.T) {
	storageFile := storage.StorageFile{Path: "../../test/testdata/go_land.mp4"}
	thumbnailVideo := ThumbnailVideo{Storage: &storageFile}

	thumbnail, _ := thumbnailVideo.Open()
	response, _ := thumbnail.Generate(10, 10, 5)

	assert.Contains(t, response, ".gif")
}
