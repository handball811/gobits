package conv

import (
	"github.com/handball811/gobits"
)

type Uint32Reader interface {
	Read(gobits.Reader, *uint32) error
}

type Uint32Writer interface {
	Write(gobits.Writer, uint32) error
}

type BitUint32Reader struct {
	Uint32Reader
	b  []byte
	s  *gobits.Slice
	bs int
}

func NewBitUint32Reader(
	bitsize int,
) *BitUint32Reader {
	b := make([]byte, 4)
	s := gobits.NewSlice(b, 0, 0, bitsize)
	return &BitUint32Reader{
		b:  b,
		s:  s,
		bs: bitsize,
	}
}

func (c *BitUint32Reader) Read(r gobits.Reader, output *uint32) error {
	c.s.Move(0, 0)
	_, err := r.Read(c.s)
	if err != nil {
		return err
	}
	var ret uint32 = 0
	for i, b := range c.b {
		ret |= uint32(b) << (i << 3)
	}
	*output = ret
	return nil
}

type BitUint32Writer struct {
	Uint32Writer
	b  []byte
	s  *gobits.Slice
	bs int
}

func NewBitUint32Writer(
	bitsize int,
) *BitUint32Writer {
	b := make([]byte, 4)
	s := gobits.NewSlice(b, 0, 0, bitsize)
	return &BitUint32Writer{
		b:  b,
		s:  s,
		bs: bitsize,
	}
}

func (c *BitUint32Writer) Write(w gobits.Writer, input uint32) error {
	for i := range c.b {
		c.b[i] = byte((input >> (i << 3)) & 255)
	}
	c.s.Move(0, c.bs)
	_, err := w.Write(c.s)
	return err
}
