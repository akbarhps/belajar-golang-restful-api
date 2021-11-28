package test

import (
	"programmerzamannow/belajar-golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.Error(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	assert.NoError(t, err)
	assert.NotNil(t, simpleService)
}
