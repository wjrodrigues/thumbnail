package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsResponseWhenImageThumbnail(t *testing.T) {
	service := Creator[string, *os.File]{Width: 10, Height: 10, Time: 0}
	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"

	response := service.Call(url)

	assert.NotNil(t, response.result)
	assert.Equal(t, response.errors, nil)
}

func TestReturnsResponseWhenVideoThumbnail(t *testing.T) {
	service := Creator[string, *os.File]{Width: 10, Height: 10, Time: 5}
	url := "../../test/testdata/go_land.mp4"

	response := service.Call(url)

	assert.NotNil(t, response.result)
	assert.Equal(t, response.errors, nil)
}

func TestReturnsSErrorResponseWhenCreatorFails(t *testing.T) {
	service := Creator[string, *os.File]{Width: -10, Height: 10, Time: 0}
	url := "../../test/testdata/google_logo.png"

	response := service.Call(url)

	assert.Nil(t, response.result)
	assert.EqualError(t, response.errors, "png: invalid format: invalid image size: 0x0")
}

func TestReturnErrorResponseWhenFormatIsInvalid(t *testing.T) {
	service := Creator[string, *os.File]{Width: 10, Height: 10, Time: 0}
	url := "../../test/testdata/invalid_format.html"

	response := service.Call(url)

	assert.Nil(t, response.result)
	assert.EqualError(t, response.errors, "unsupported format or invalid file")
}
