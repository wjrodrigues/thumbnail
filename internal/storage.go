package internal

import (
	"errors"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var Formats = map[string][]string{"image": {"jpeg", "jpg", "gif", "png"}}
var TargetPath = fmt.Sprintf("%s/thumbnail_", os.TempDir())

type Storage interface {
	Open(path string) (Thumbnail, error)
}

type StorageFile struct {
	file *os.File
	path string
}

func (storage *StorageFile) Open() (Thumbnail, error) {
	if storage.Supported(Formats["image"]) {
		file, err := ThumbnailImage{}.Open(storage)
		if err != nil {
			return nil, err
		}

		return file, nil
	}
	return nil, errors.New("unsupported format")
}

func (storage *StorageFile) GetFile() (*StorageFile, error) {
	file, err := os.Open(storage.path)
	if err == nil {
		storage.file = file
		return storage, nil
	}
	defer file.Close()

	url, err := url.ParseRequestURI(storage.path)
	if err != nil {
		return nil, err
	}

	fileName, err := download(url)

	if err != nil {
		return nil, err
	}

	file, _ = os.Open(fileName)
	storage.file = file

	return storage, nil
}

func getExtension(fileName string) string {
	splitPath := strings.Split(fileName, ".")

	return splitPath[len(splitPath)-1]
}

func download(path *url.URL) (string, error) {
	extension := getExtension(path.String())

	resp, err := http.Get(path.String())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	timestamp := time.Now().UnixNano()
	fileName := fmt.Sprintf("%s%d.%s", TargetPath, timestamp, extension)

	out, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	io.Copy(out, resp.Body)

	return fileName, nil
}

func (storage *StorageFile) Supported(extensions []string) bool {
	extension := getExtension(storage.path)

	for _, target := range extensions {
		if target == extension {
			return true
		}
	}

	return false
}
