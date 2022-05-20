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

	"github.com/google/uuid"
)

var TargetPath = fmt.Sprintf("%s/thumbnail", os.TempDir())

type Storage interface {
	GetFile() (Storage, error)
	Supported(extensions []string) bool
}

type StorageFile struct {
	File *os.File
	Path string
}

func init() {
	createTempFolder(TargetPath)
}

func (storage *StorageFile) GetFile() (Storage, error) {
	var file *os.File

	if _, err := os.Stat(storage.Path); !os.IsNotExist(err) {
		file, err = saveTempFile(storage.Path)

		if err != nil {
			return nil, err
		}
	} else {
		url, err := url.ParseRequestURI(storage.Path)
		if err != nil {
			return nil, err
		}

		file, err = download(url)

		if err != nil {
			return nil, err
		}
	}

	storage.File = file

	defer file.Close()

	return storage, nil
}

func getExtension(fileName string) string {
	splitPath := strings.Split(fileName, ".")

	return splitPath[len(splitPath)-1]
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

func download(path *url.URL) (*os.File, error) {
	resp, err := http.Get(path.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	out, err := saveTempFile(path.String())
	if err != nil {
		return nil, err
	}

	io.Copy(out, resp.Body)

	return out, nil
}

func saveTempFile(path string) (*os.File, error) {
	extension := getExtension(path)

	uuid := uuid.New().String()
	fileName := fmt.Sprintf("%s/%s.%s", TargetPath, uuid, extension)

	out, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	return out, nil
}

func createTempFolder(path string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return
	}

	os.Mkdir(path, 0755)
}
