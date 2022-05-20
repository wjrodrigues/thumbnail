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
	Resource() *os.File
}

type StorageFile struct {
	file *os.File
	Path string
}

func init() {
	createTempFolder(TargetPath)
}

func (storage *StorageFile) GetFile() (Storage, error) {
	var file *os.File
	file, _ = localFile(storage)

	if file == nil {
		url, err := url.ParseRequestURI(storage.Path)
		if err != nil {
			return nil, err
		}

		file, err = download(url)

		if err != nil {
			return nil, err
		}
	}

	storage.file = file

	defer file.Close()

	return storage, nil
}

func Extension(fileName string) (string, string) {
	splitPath := strings.Split(fileName, ".")

	return splitPath[0], splitPath[len(splitPath)-1]
}

func (storage *StorageFile) Supported(extensions []string) bool {
	_, extension := Extension(storage.Path)

	for _, target := range extensions {
		if target == extension {
			return true
		}
	}

	return false
}

func (storage *StorageFile) Resource() *os.File {
	return storage.file
}

func localFile(storage *StorageFile) (*os.File, error) {
	if _, err := os.Stat(storage.Path); os.IsNotExist(err) {
		return nil, err
	}

	tmpFile, err := createTempFile(storage.Path)
	if err != nil {
		return nil, err
	}
	defer tmpFile.Close()

	source, _ := os.Open(storage.Path)
	io.Copy(tmpFile, source)

	defer source.Close()

	return tmpFile, nil
}

func download(path *url.URL) (*os.File, error) {
	resp, err := http.Get(path.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	tmpFile, err := createTempFile(path.String())
	if err != nil {
		return nil, err
	}
	defer tmpFile.Close()

	io.Copy(tmpFile, resp.Body)

	return tmpFile, err
}

func createTempFile(path string) (*os.File, error) {
	_, extension := Extension(path)

	uuid := uuid.New().String()
	fileName := fmt.Sprintf("%s/%s.%s", TargetPath, uuid, extension)

	tmpFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return tmpFile, nil
}

func createTempFolder(path string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return
	}

	os.Mkdir(path, 0755)
}
