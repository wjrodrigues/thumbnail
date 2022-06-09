package types

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"thumbnail/internal/storage"
)

var VideoFormats = []string{"avi", "mp4", "mov"}

type ThumbnailVideo struct {
	Storage  storage.Storage
	Resource *os.File
}

func (video *ThumbnailVideo) Open() (Thumbnail, error) {
	if !video.Storage.Supported(VideoFormats) {
		return nil, errors.New("unsupported format")
	}

	_, err := video.Storage.GetFile()

	if err != nil {
		return nil, err
	}

	file, _ := os.Open(video.Storage.Resource().Name())

	video.Resource = file

	return video, err
}

func (video *ThumbnailVideo) Generate(width, height, duration int) (string, error) {
	pathWithoutExtension, _ := storage.Extension(video.Storage.Resource().Name())

	gifFullPath := fmt.Sprintf("%s.gif", pathWithoutExtension)
	resourceFullPath := video.Storage.Resource().Name()

	args := prepareArgs(width, height, duration, resourceFullPath, gifFullPath)

	_, err := exec.Command("ffmpeg", args...).CombinedOutput()

	return gifFullPath, err
}

func prepareArgs(width, height, duration int, resourceFullPath, gifFullPath string) []string {
	extraArg := fmt.Sprintf(
		"fps=10,scale=%d:%d:flags=fast_bilinear,split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse", width, height)

	durationString := fmt.Sprintf("%d", duration)
	return []string{
		"-ss", "0", "-t", durationString, "-i", resourceFullPath, "-vf", extraArg, "-loop", "0", gifFullPath,
	}
}
