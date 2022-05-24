package types

import (
	"testing"
	"thumbnail/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestReturnsErrorWhenOpenImage(t *testing.T) {
	thumbnailImage := ThumbnailVideo{}
	storage := storage.StorageFile{Path: "../../test/testdata/google_logo.png"}
	_, err := thumbnailImage.Open(&storage)

	assert.EqualError(t, err, "unsupported format")
}

func TestReturnsThumbnailInstanceWhenOpeningVideo(t *testing.T) {
	thumbnailVideo := ThumbnailVideo{}
	storage := storage.StorageFile{Path: "../../test/testdata/go_land.mp4"}
	response, _ := thumbnailVideo.Open(&storage)

	assert.Implements(t, (*Thumbnail)(nil), response)
}

func TestReturnsErrorOpeningInvalidVideoPath(t *testing.T) {
	thumbnailVideo := ThumbnailVideo{}
	storage := storage.StorageFile{Path: "../../test/testdata/invalid.mp4"}
	_, err := thumbnailVideo.Open(&storage)

	assert.Error(t, err)
}

func TestReturnsGIFThumbnailOfVideo(t *testing.T) {
	thumbnailVideo := ThumbnailVideo{}
	storageFile := storage.StorageFile{Path: "../../test/testdata/go_land.mp4"}
	thumbnail, _ := thumbnailVideo.Open(&storageFile)
	response, _ := thumbnail.Generate(10, 10, 5, &storageFile)

	assert.Contains(t, response, ".gif")
}
