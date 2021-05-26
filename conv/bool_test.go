package conv_test

import (
	"testing"

	"github.com/handball811/gobits"
	. "github.com/handball811/gobits/conv"
	"github.com/stretchr/testify/assert"
)

func TestOneBoolWriter(t *testing.T) {
	// setup
	var err error
	target := NewOneBoolWriter()
	buffer := make([]byte, 1)
	slice := gobits.NewSlice(buffer, 0, 0, 8)
	segment := gobits.NewSegment(slice)

	// when
	err = target.Write(segment, true)

	// then
	assert.Nil(t, err)
	assert.Equal(t, 1, segment.Len())
	assert.Equal(t, byte(1), buffer[0])
}

func TestOneBoolReader(t *testing.T) {
	// setup
	var err error
	var result bool = false
	target := NewOneBoolReader()
	writer := NewOneBoolWriter()
	buffer := make([]byte, 1)
	slice := gobits.NewSlice(buffer, 0, 0, 8)
	segment := gobits.NewSegment(slice)

	// when
	writer.Write(segment, true)
	err = target.Read(segment, &result)

	// then
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}
