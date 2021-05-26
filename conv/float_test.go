package conv_test

import (
	"testing"

	"github.com/handball811/gobits"
	. "github.com/handball811/gobits/conv"
	"github.com/stretchr/testify/assert"
)

func TestClampFloat32InRange(t *testing.T) {
	// setup
	var num float32 = 3.65487
	var result float32 = 0
	writer := NewClampFloat32Writer(12, -5, 5)
	reader := NewClampFloat32Reader(12, -5, 5)
	buffer := make([]byte, 4)
	slice := gobits.NewSlice(buffer, 0, 0, 32)
	segment := gobits.NewSegment(slice)

	// when
	writer.Write(segment, num)
	reader.Read(segment, &result)

	// then
	assert.InDelta(t, num, result, 5/float64(1<<12))
}

func TestClampFloat32InRange2(t *testing.T) {
	// setup
	var num float32 = -3.65487
	var result float32 = 0
	writer := NewClampFloat32Writer(12, -5, 5)
	reader := NewClampFloat32Reader(12, -5, 5)
	buffer := make([]byte, 4)
	slice := gobits.NewSlice(buffer, 0, 0, 32)
	segment := gobits.NewSegment(slice)

	// when
	writer.Write(segment, num)
	reader.Read(segment, &result)

	// then
	assert.InDelta(t, num, result, 5/float64(1<<12))
}
func TestClampFloat32Lower(t *testing.T) {
	// setup
	var num float32 = -8.65487
	var result float32 = 0
	writer := NewClampFloat32Writer(12, -5, 5)
	reader := NewClampFloat32Reader(12, -5, 5)
	buffer := make([]byte, 4)
	slice := gobits.NewSlice(buffer, 0, 0, 32)
	segment := gobits.NewSegment(slice)

	// when
	writer.Write(segment, num)
	reader.Read(segment, &result)

	// then
	assert.InDelta(t, -5, result, 0.00000001)
}

func TestClampFloat32Upper(t *testing.T) {
	// setup
	var num float32 = 8.65487
	var result float32 = 0
	writer := NewClampFloat32Writer(12, -5, 5)
	reader := NewClampFloat32Reader(12, -5, 5)
	buffer := make([]byte, 4)
	slice := gobits.NewSlice(buffer, 0, 0, 32)
	segment := gobits.NewSegment(slice)

	// when
	writer.Write(segment, num)
	reader.Read(segment, &result)

	// then
	assert.InDelta(t, 5, result, 0.00000001)
}
