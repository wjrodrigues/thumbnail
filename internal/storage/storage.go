package storage

import (
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

var TargetPath = fmt.Sprintf("%s/thumbnail_", os.TempDir())

type Storage interface {
	GetFile() (Storage, error)
	Supported(extensions []string) bool
}

type StorageFile struct {
	File *os.File
	Path string
}

func (storage *StorageFile) GetFile() (Storage, error) {
	file, err := os.Open(storage.Path)
	if err == nil {
		storage.File = file
		return storage, nil
	}

	defer file.Close()

	url, err := url.ParseRequestURI(storage.Path)
	if err != nil {
		return nil, err
	}

	fileName, err := download(url)

	if err != nil {
		return nil, err
	}

	file, _ = os.Open(fileName)
	storage.File = file

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
	extension := getExtension(storage.Path)

	for _, target := range extensions {
		if target == extension {
			return true
		}
	}

	return false
}
