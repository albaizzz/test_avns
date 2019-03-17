package test_avns

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestBinary_gap(t *testing.T) {
	assert.Equal(t, 2, Binary_gap(9))
	assert.Equal(t, 4, Binary_gap(529))
	assert.Equal(t, 0, Binary_gap(32))
}
