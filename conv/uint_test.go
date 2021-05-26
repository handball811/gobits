package conv_test

import (
	"testing"

	"github.com/handball811/gobits"
	. "github.com/handball811/gobits/conv"
	"github.com/stretchr/testify/assert"
)

func TestBitUint32Writer(t *testing.T) {
	// setup
	var err error
	var num uint32 = (1 << 25) - 1
	target := NewBitUint32Writer(25)
	buffer := make([]byte, 4)
	slice := gobits.NewSlice(buffer, 0, 0, 32)
	segment := gobits.NewSegment(slice)

	// when
	err = target.Write(segment, num)

	// then
	assert.Nil(t, err)
	assert.Equal(t, 25, segment.Len())
	assert.Equal(t, byte(num&255), buffer[0])
	assert.Equal(t, byte((num>>8)&255), buffer[1])
	assert.Equal(t, byte((num>>16)&255), buffer[2])
	assert.Equal(t, byte((num>>24)&255), buffer[3])
}

func TestBitUint32Reader(t *testing.T) {
	// setup
	var err error
	var num uint32 = (1 << 25) - 1
	var result uint32
	target := NewBitUint32Reader(25)
	writer := NewBitUint32Writer(25)
	buffer := make([]byte, 4)
	slice := gobits.NewSlice(buffer, 0, 0, 32)
	segment := gobits.NewSegment(slice)

	// when
	writer.Write(segment, num)
	err = target.Read(segment, &result)

	// then
	assert.Nil(t, err)
	assert.Equal(t, num, result)
}
