package conv

import (
	"math"

	"github.com/handball811/gobits"
)

type Float32Reader interface {
	Read(gobits.Reader, *float32) error
}

type Float32Writer interface {
	Write(gobits.Writer, float32) error
}

type ClampFloat32Reader struct {
	r    *BitUint32Reader
	min  float32
	max  float32
	unit float64
}

func NewClampFloat32Reader(
	bits int,
	min float32,
	max float32,
) *ClampFloat32Reader {
	var maxsize uint32 = math.MaxUint32
	if bits < 32 {
		maxsize = (uint32(1) << bits) - 1
	}
	return &ClampFloat32Reader{
		min:  min,
		max:  max,
		unit: float64(max-min) / float64(maxsize),
		r:    NewBitUint32Reader(bits),
	}
}

func (c *ClampFloat32Reader) Read(r gobits.Reader, output *float32) error {
	var u uint32
	if err := c.r.Read(r, &u); err != nil {
		return err
	}
	*output = float32(c.unit*float64(u)) + c.min
	return nil
}

type ClampFloat32Writer struct {
	w    *BitUint32Writer
	min  float32
	max  float32
	unit float64
}

func NewClampFloat32Writer(
	bits int,
	min float32,
	max float32,
) *ClampFloat32Writer {
	var maxsize uint32 = math.MaxUint32
	if bits < 32 {
		maxsize = (uint32(1) << bits) - 1
	}
	return &ClampFloat32Writer{
		min:  min,
		max:  max,
		unit: float64(max-min) / float64(maxsize),
		w:    NewBitUint32Writer(bits),
	}
}

func (c *ClampFloat32Writer) Write(w gobits.Writer, input float32) error {
	if input < c.min {
		input = c.min
	} else if input > c.max {
		input = c.max
	}
	num := uint32(math.Round(float64(input-c.min) / c.unit))
	return c.w.Write(w, num)
}
