package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetHello(t *testing.T) {
	expected := "Hello World from TEST"
	actual := GetHello("TEST")

	assert.Equal(t, actual, expected)
}
