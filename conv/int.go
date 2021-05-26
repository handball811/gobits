package conv

import "github.com/handball811/gobits"

type Int32Reader interface {
	Read(gobits.Reader, *int32) error
}

type Int32Writer interface {
	Write(gobits.Writer, int32) error
}

type BitInt32Reader struct {
	Int32Reader
	r   *BitUint32Reader
	sub uint32
}

func NewBitInt32Reader(
	bitsize int,
) *BitInt32Reader {
	return &BitInt32Reader{
		r:   NewBitUint32Reader(bitsize),
		sub: uint32(1) << (bitsize - 1),
	}
}

func (c *BitInt32Reader) Read(r gobits.Reader, output *int32) error {
	var u uint32
	if err := c.r.Read(r, &u); err != nil {
		return err
	}
	if u >= c.sub {
		*output = int32(u - c.sub)
	} else {
		*output = -int32(c.sub - u)
	}
	return nil
}

type BitInt32Writer struct {
	Int32Writer
	w   *BitUint32Writer
	sub uint32
}

func NewBitInt32Writer(
	bitsize int,
) *BitInt32Writer {
	return &BitInt32Writer{
		w:   NewBitUint32Writer(bitsize),
		sub: uint32(1) << (bitsize - 1),
	}
}

func (c *BitInt32Writer) Write(w gobits.Writer, input int32) error {
	var u uint32
	if input >= 0 {
		u = uint32(input) + c.sub
	} else {
		u = c.sub - uint32(-input)
	}
	return c.w.Write(w, u)
}
