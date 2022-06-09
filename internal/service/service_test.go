package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateServiceResponseStruct(t *testing.T) {
	response := Response[string]{result: "Golang", errors: nil}

	assert.Equal(t, response.result, "Golang")
	assert.Equal(t, response.errors, nil)
}
