package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	assert.Equal(t, Add(1, 2), 3)
	assert.Equal(t, Add(1, 3), 4)
}
