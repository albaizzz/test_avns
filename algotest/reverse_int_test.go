package test_avns

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestReverse_int(t *testing.T) {
	assert.Equal(t, 321, Reverse_int(123))
	assert.Equal(t, 43, Reverse_int(34))
	assert.Equal(t, 3, Reverse_int(3))
}
