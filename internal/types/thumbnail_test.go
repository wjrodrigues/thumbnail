package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenURLAndReturnsInstanceOfThumbnailImage(t *testing.T) {
	resource, _ := Open("../../test/testdata/google_logo.png")

	assert.IsType(t, (*ThumbnailImage)(nil), resource)
}

func TestReturnErrorWhenOpeningUnsupportedFormat(t *testing.T) {
	_, err := Open("../../test/testdata/google_logo.mp4")

	assert.Error(t, err)
}
