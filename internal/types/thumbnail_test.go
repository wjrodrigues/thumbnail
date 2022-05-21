package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenURLAndReturnsInstanceOfThumbnailImage(t *testing.T) {
	resource, _ := Open("../../test/testdata/google_logo.png")

	assert.IsType(t, (*ThumbnailImage)(nil), resource)
}

func TestOpenURLAndReturnsInstanceOfThumbnailVideo(t *testing.T) {
	resource, _ := Open("../../test/testdata/go_land.mp4")

	assert.IsType(t, (*ThumbnailVideo)(nil), resource)
}

func TestReturnsErrorWhenFormatIsNotSupported(t *testing.T) {
	_, err := Open("../../test/testdata/invalid_format.html")

	assert.EqualError(t, err, "unsupported format")
}
