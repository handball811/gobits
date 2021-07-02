package gobits_test

import (
	"testing"

	. "github.com/handball811/gobits"
	"github.com/stretchr/testify/assert"
)

func TestFixedSegmentRead(t *testing.T) {
	// setup
	var err error
	var size int
	target := NewFixedSegment(
		NewSegmentWithSize(8),
	)
	readBuffer := make([]byte, 64)
	readSlice := NewSlice(readBuffer, 8, 8, 40)

	// when
	target.Write(NewSliceWithBuffer([]byte{1, 2, 3}))
	target.Read(readSlice)
	size, err = target.Read(readSlice)

	// then
	assert.Nil(t, err)
	assert.Equal(t, 24, size)
	assert.Equal(t, 24, target.Len())
	assert.Equal(t, byte(0), readBuffer[0])
	assert.Equal(t, byte(1), readBuffer[1])
	assert.Equal(t, byte(2), readBuffer[2])
	assert.Equal(t, byte(3), readBuffer[3])
	assert.Equal(t, byte(0), readBuffer[4])
	assert.Equal(t, byte(0), readBuffer[5])
}
