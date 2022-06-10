package service

import (
	"fmt"
	"os"
	"thumbnail/internal/types"
)

type Creator[T string, R *os.File] struct {
	response Response[R]
	Width    int
	Height   int
	Time     int
}

func (service *Creator[T, R]) Call(path T) Response[R] {
	thumbnail, err := types.Open(fmt.Sprint(path))

	if err != nil {
		service.response.errors = err
		return service.response
	}

	pathResult, err := thumbnail.Generate(service.Width, service.Height, service.Time)
	if err != nil {
		service.response.errors = err
		return service.response
	}

	result, _ := os.Open(pathResult)
	service.response.result = result

	return service.response
}
