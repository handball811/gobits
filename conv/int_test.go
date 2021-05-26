package conv_test

import (
	"testing"

	"github.com/handball811/gobits"
	. "github.com/handball811/gobits/conv"
	"github.com/stretchr/testify/assert"
)

func TestBitInt32WriterPositive(t *testing.T) {
	// setup
	var err error
	var num int32 = (1 << 23) - 1
	target := NewBitInt32Writer(25)
	buffer := make([]byte, 4)
	slice := gobits.NewSlice(buffer, 0, 0, 32)
	segment := gobits.NewSegment(slice)

	// when
	err = target.Write(segment, num)

	// then
	num += (1 << 24)
	assert.Nil(t, err)
	assert.Equal(t, 25, segment.Len())
	assert.Equal(t, byte(num&255), buffer[0])
	assert.Equal(t, byte((num>>8)&255), buffer[1])
	assert.Equal(t, byte((num>>16)&255), buffer[2])
	assert.Equal(t, byte((num>>24)&255), buffer[3])
}

func TestBitInt32WriterNegative(t *testing.T) {
	// setup
	var err error
	var num int32 = -(1 << 23) + 1
	target := NewBitInt32Writer(25)
	buffer := make([]byte, 4)
	slice := gobits.NewSlice(buffer, 0, 0, 32)
	segment := gobits.NewSegment(slice)

	// when
	err = target.Write(segment, num)

	// then
	num += (1 << 24)
	assert.Nil(t, err)
	assert.Equal(t, 25, segment.Len())
	assert.Equal(t, byte(num&255), buffer[0])
	assert.Equal(t, byte((num>>8)&255), buffer[1])
	assert.Equal(t, byte((num>>16)&255), buffer[2])
	assert.Equal(t, byte((num>>24)&255), buffer[3])
}

func TestBitInt32ReaderPositive(t *testing.T) {
	// setup
	var err error
	var num int32 = (1 << 23) - 1
	var result int32
	target := NewBitInt32Reader(25)
	writer := NewBitInt32Writer(25)
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

func TestBitInt32ReaderNegative(t *testing.T) {
	// setup
	var err error
	var num int32 = -(1 << 23) + 1
	var result int32
	target := NewBitInt32Reader(25)
	writer := NewBitInt32Writer(25)
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
