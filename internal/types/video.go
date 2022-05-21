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
	Resource *os.File
}

func (video *ThumbnailVideo) Open(storage storage.Storage) (Thumbnail, error) {
	if !storage.Supported(VideoFormats) {
		return nil, errors.New("unsupported format")
	}

	_, err := storage.GetFile()

	if err != nil {
		return nil, err
	}

	file, _ := os.Open(storage.Resource().Name())

	video.Resource = file

	return video, err
}

func (video *ThumbnailVideo) Generate(width, height, time int, storageFile storage.Storage) (string, error) {
	pathWithoutExtension, _ := storage.Extension(storageFile.Resource().Name())

	gifPath := fmt.Sprintf("%s.gif", pathWithoutExtension)

	//	filterCmd := fmt.Sprintf("[0:v] fps=15,scale=w=%d:h=%d,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1", width, height)

	//	rawCmd := fmt.Sprintf("-ss 0 -t 10.0 -i %s -filter_complex '%s' %s", storageFile.Resource().Name(), filterCmd, gifPath)
	args := fmt.Sprintf("-ss 0 -t 10.0 -i %s -filter_complex '[0:v] fps=15,scale=w=250:h=100,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1' %s", storageFile.Resource().Name(), gifPath)
	cmd := exec.Command("ffmpeg", args)
	stdout, err := cmd.Output()
	fmt.Println(string(stdout), err)

	return "", nil

	//ffmpeg -ss 0 -t 10.0 -i /tmp/thumbnail/0f4186e0-1c4c-4472-8990-027a1e545bc8.mp4 -filter_complex '[0:v] fps=15,scale=w=250:h=100,split [a][b];[a] palettegen=stats_mode=single [p];[b][p] paletteuse=new=1' /tmp/thumbnail/0f4186e0-1c4c-4472-8990-027a1e545bc8.gif
}
